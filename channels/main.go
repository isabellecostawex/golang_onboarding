package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.stackoverflow.com/",
		"https://www.golang.org/",
		"https://www.amazon.com/",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}
	fmt.Println(<- channel)
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		channel <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")
	channel <- "Yep its up"
}
