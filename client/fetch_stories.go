package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hackernews-go-cli/model"
)

const baseStoryURL = "https://hacker-news.firebaseio.com/v0/"

// GetStories returns top stories
func GetStories(num int, storyType string) {
	storyURL := fmt.Sprintf("%s%s%s", baseStoryURL, storyType, ".json")
	var itemArray []model.Story

	response, err := http.Get(storyURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	storyList, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var topStories []int64
	json.Unmarshal(storyList, &topStories)
	if len(topStories) < num {
		fmt.Println("You have exceeded the max number of item, if you want to list all the item, use this as an argument: ", len(topStories))
	} else {
		for i := 0; i < num; i++ {
			itemArray = append(itemArray, GetItem(topStories[i]))
		}
		ShowTable(itemArray)
	}
}
