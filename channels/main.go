package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for u := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkUrl(url, c)
		}(u)
	}
}

func checkUrl(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "might be down!")
		c <- url
		return
	}

	fmt.Println(url, "is up! Status:", resp.Status)
	c <- url
}
