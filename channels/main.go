package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
	}

	for _, url := range urls {
		go checkUrl(url)
	}
}

func checkUrl(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "might be down!")
		fmt.Println(url, resp.Status)
		return
	}

	fmt.Println(url, "is up! Status:", resp.Status)
}
