package main

import (
	"net/url"
	"strings"
)

func normalizeURL(urlS string) (string, error) {
	parsedUrl, err := url.Parse(urlS)
	if err != nil {
		return "", err
	}
	normalurl := parsedUrl.Host + parsedUrl.Path
	normalurl = strings.ToLower(normalurl)
	normalurl = strings.TrimSuffix(normalurl, "/")
	return normalurl, nil
}
