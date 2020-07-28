package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hackernews-go-cli/client"
)

func main() {
	fmt.Println("Hackernews cli application that written in Go!\n---------------")
	response, err := http.Get("https://hacker-news.firebaseio.com/v0/maxitem.json")
	if err != nil {
		log.Fatalln(err)
	}

	client.GetLastItem(*response)

}
