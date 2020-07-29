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
	flag.Parse()
	if *getLastItem {
		response, err := http.Get("https://hacker-news.firebaseio.com/v0/maxitem.json")
		if err != nil {
			log.Fatalln(err)
		}
		defer response.Body.Close()
		client.GetLastItem(BaseURL, *response)
	}

	getTopStories := flag.Bool("top", false, "display top stories, you can pass arguments until 500. default first 100 stories will shown.")

	topResponse, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(topResponse.Body)
	topStory, err := ioutil.ReadAll(topResponse.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var topStories []int64
	json.Unmarshal(topStory, &topStories)
	fmt.Println(topStories[0])

}
