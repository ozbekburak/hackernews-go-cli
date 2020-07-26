package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hackernews cli application that written in Go!")
	response, err := http.Get("https://hacker-news.firebaseio.com/v0/maxitem.json")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response)
}
