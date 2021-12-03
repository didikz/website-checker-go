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

	for _, url := range urls {
		checkLink(url)
	}
}

func checkLink(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down")
		return
	}
	fmt.Println(url, "is up")
}
