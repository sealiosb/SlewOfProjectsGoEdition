package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var url string
	var shortenedURL string

	fmt.Println("URL Shortener Service")
	fmt.Print("Enter url: ")
	fmt.Scanln(&url)

	if len(url) == 0 {
		fmt.Println("Error: URL cannot be empty.")
		return
	}

	shortenedURL = "short.ly/" + url[len(url)-6:]
	fmt.Println(shortenedURL)

}
