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

	// so we can create infinite loop then check link will run until we stopped
	for link := range c {
		go checkLink(link, c)
	}
}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		down := url + " is down"
		fmt.Println(down)
		c <- url
		return
	}
	up := url + " is up"
	fmt.Println(up)
	c <- url
}
