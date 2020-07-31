package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/hackernews-go-cli/client"
)

// BaseURL helps us to create customizable request
const BaseURL = "https://hacker-news.firebaseio.com/v0/item/"

func main() {
	var storyCount int
	getLastItem := flag.Bool("last", false, "display last post, it can be story, comment, job, poll or ask")
	flag.IntVar(&storyCount, "top", 5, "display top stories, you can pass arguments until 500")
	flag.Parse()
	if *getLastItem {
		response, err := http.Get("https://hacker-news.firebaseio.com/v0/maxitem.json")
		if err != nil {
			log.Fatalln(err)
		}
		defer response.Body.Close()
		client.GetLastItem(BaseURL, *response)
	}
	if storyCount > 0 && storyCount <= 500 {
		client.GetTopStories(storyCount)
	}
}
