package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("raw base URL: %v", err)
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("raw current URL: %v", err)
		return
	}
	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}
	norCurURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("normalize current URL: %v", err)
	}
	_, ok := pages[norCurURL]
	if ok {
		pages[norCurURL] += 1
		return
	}

	pages[norCurURL] = 1

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("get html form current URL: %v", err)
		return
	}
	URLs, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("get url list form current URL: %v", err)
	}
	for _, url := range URLs {
		crawlPage(rawBaseURL, url, pages)
	}
}
