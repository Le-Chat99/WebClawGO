package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	reshtml, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("fail to get: %v", err)
	}
	if reshtml.StatusCode > 399 {
		return "", fmt.Errorf("error code: %v", reshtml.StatusCode)
	}
	if !strings.HasPrefix(reshtml.Header.Get("Content-Type"), "text/html") {
		return "", fmt.Errorf("content Type is not text/html")
	}
	readedhtml, err := io.ReadAll(reshtml.Body)
	if err != nil {
		return "", fmt.Errorf("fail to get: %v", err)
	}
	return string(readedhtml), nil
}
