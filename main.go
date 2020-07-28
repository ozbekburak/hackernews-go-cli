package main

import (
	"log"
	"net/http"

	"github.com/hackernews-go-cli/client"
)

func main() {
	response, err := http.Get("https://hacker-news.firebaseio.com/v0/maxitem.json")
	if err != nil {
		log.Fatalln(err)
	}

	client.GetLastItem(*response)

}
