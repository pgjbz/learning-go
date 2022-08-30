package main

import "fmt"

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls[]string) (results map[string]bool) {
	results = make(map[string]bool)
	resultChannel := make(chan result)
	defer close(resultChannel)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++{
		r := <- resultChannel
		results[r.string] = r.bool
	}

	return
}

func main() {
	fmt.Println("Hello World")
}
