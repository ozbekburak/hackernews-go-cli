package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hackernews-go-cli/client"
)

// BaseURL helps us to create customizable request
const BaseURL = "https://hacker-news.firebaseio.com/v0/item/"

func main() {
	getLastItem := flag.Bool("last", false, "display last post, it can be story, comment, job, poll or ask")
	getTopStories := flag.Bool("top", false, "display top stories, you can pass arguments until 500. default first 100 stories will shown.")
	flag.Parse()
	if *getLastItem {
		response, err := http.Get("https://hacker-news.firebaseio.com/v0/maxitem.json")
		if err != nil {
			log.Fatalln(err)
		}
		defer response.Body.Close()
		client.GetLastItem(BaseURL, *response)
	}

	if *getTopStories {
		response, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
		if err != nil {
			log.Fatalln(err)
		}
		topStory, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var topStories []int64
		json.Unmarshal(topStory, &topStories)
		fmt.Println(topStories)
	}

}
