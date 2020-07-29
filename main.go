package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/hackernews-go-cli/client"
)

// BaseURL helps us to create customizable request
const BaseURL = "https://hacker-news.firebaseio.com/v0/item/"

func main() {
	getLastItem := flag.Bool("last", false, "display last post, it can be story, comment, job, poll or ask")
	flag.Parse()
	if *getLastItem {
		response, err := http.Get("https://hacker-news.firebaseio.com/v0/maxitem.json")
		if err != nil {
			log.Fatalln(err)
		}
		defer response.Body.Close()
		client.GetLastItem(BaseURL, *response)
	}

	topResponse, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(topResponse.Body)
}
