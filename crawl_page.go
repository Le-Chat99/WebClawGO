package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer func() {
		cfg.wg.Done()
		<-cfg.concurrencyControl
	}()
	cfg.concurrencyControl <- struct{}{}
	if cfg.pagesLen() >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}
	norCurURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
	}
	isFirst := cfg.addPageVisit(norCurURL)
	if !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("get html form %v: %v\n", rawCurrentURL, err)
		return
	}
	URLs, err := getURLsFromHTML(htmlBody, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("get url list form current URL: %v", err)
	}
	for _, url := range URLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}

}
