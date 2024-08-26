package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer func() {
		cfg.wg.Done()
		<-cfg.concurrencyControl
	}()
	cfg.concurrencyControl <- struct{}{}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}
	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}
	norCurURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}
	isFirst := cfg.addPageVisit(norCurURL)
	if !isFirst {
		return
	}

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("get html form %v: %v\n", rawCurrentURL, err)
		return
	}
	URLs, err := getURLsFromHTML(htmlBody, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("get url list form current URL: %v", err)
		return
	}
	for _, url := range URLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
