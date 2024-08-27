package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 3 {
		fmt.Print("no website provided")
		os.Exit(1)
	}
	if len(argsWithoutProg) > 3 {
		fmt.Print("too many arguments provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %v \n", os.Args[1])
	rawBaseURL := os.Args[1]
	concurrentLimit, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("arguments 2 wrong type: %v\n", err)
	}
	maxpages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("arguments 3 wrong type: %v\n", err)
	}
	cfg, err := conifg(rawBaseURL, concurrentLimit, maxpages)
	if err != nil {
		fmt.Printf("Config fail: %v\n", err)
	}
	cfg.wg.Add(1)
	cfg.crawlPage(os.Args[1])
	cfg.wg.Wait()
	cfg.Report()
}
