package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.google.co.id",
		"https://www.amazon.com",
		"https://www.facebook.com",
		"https://go.dev",
		"https://pkg.go.dev",
		"https://laravel.com/",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkLink(url, c)
	}

	for {
		fmt.Println(<-c)
	}
}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		down := url + " is down"
		fmt.Println(down)
		c <- down + " (channel)"
		return
	}
	up := url + " is up"
	fmt.Println(up)
	c <- up + " (channel)"
}
