package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		fmt.Print("no website provided")
		os.Exit(1)
	}
	if len(argsWithoutProg) > 1 {
		fmt.Print("too many arguments provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %v \n", os.Args[1])
	_, err := getHTML(os.Args[1])
	if err != nil {
		fmt.Printf("ST go wrong: %v", err)
	}
	pages := map[string]int{}
	crawlPage(os.Args[1], os.Args[1], pages)
	fmt.Println(pages)

}
