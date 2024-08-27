package main

import (
	"fmt"
	"strings"
)

func (cfg *config) Report() {
	fmt.Printf("=============================\nREPORT for %v\n============================\n", strings.TrimSuffix(cfg.baseURL.String(), "/"))
	for normalizedURL, count := range cfg.pages {
		fmt.Printf("Found %d internal links to %s\n", count, normalizedURL)
	}
}
