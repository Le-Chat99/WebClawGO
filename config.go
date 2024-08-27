package main

import (
	"fmt"
	"net/url"
	"sync"
)

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	if cfg.pages[normalizedURL] == 0 {
		isFirst = true
	}
	cfg.pages[normalizedURL]++
	cfg.mu.Unlock()
	return isFirst
}

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func conifg(rawBaseURL string, gorolimit, maxpages int) (*config, error) {
	bURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("fail to Parse %v:\n%v", rawBaseURL, err)
	}
	return &config{
		pages:              make(map[string]int),
		baseURL:            bURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, gorolimit),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxpages,
	}, nil

}
