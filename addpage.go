package main

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	if cfg.pages[normalizedURL] == 0 {
		isFirst = true
	}
	cfg.pages[normalizedURL]++
	cfg.mu.Unlock()
	return isFirst
}
