package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hackernews-go-cli/model"
)

// GetStories returns top stories
func GetStories(num int, storyType string) {
	itemArray := []model.Story{}
	response, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	topStory, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var topStories []int64
	json.Unmarshal(topStory, &topStories)
	if len(topStories) < num {
		fmt.Println("You have exceeded the max number of item, if you want to list all the item, use this as an argument: ", len(topStories))
	} else {
		for i := 0; i < num; i++ {
			itemArray = append(itemArray, GetItem(topStories[i]))
		}
		ShowTable(itemArray)
	}
}