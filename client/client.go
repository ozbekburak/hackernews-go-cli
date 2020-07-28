package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hackernews-go-cli/model"
)

// BaseURL helps us to create customizable request
const BaseURL = "https://hacker-news.firebaseio.com/v0/item/"

// DisplayItem displays our output that fetched from last item
func DisplayItem(output string, typeOfPost string, itemURL string) {
	fmt.Printf("Type: %s\nURL: %s\nContent: %s", typeOfPost, itemURL, output)
}

// GetMaxItemID returns the ID of last posted item
func GetMaxItemID(response http.Response) int {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var maxItemID int
	json.Unmarshal(data, &maxItemID)
	return maxItemID

}

// GetLastItem brings the last posted item which can be in different types like story, job, comment etc.
func GetLastItem(response http.Response) {
	var ItemResponse model.Item
	maxID := GetMaxItemID(response)

	lastItemURL := fmt.Sprintf("%s%d%s", BaseURL, maxID, ".json")
	lastItem, err := http.Get(lastItemURL)
	if err != nil {
		log.Fatalln(err)
	}
	bodyLastItem, err := ioutil.ReadAll(lastItem.Body)
	if err != nil {
		log.Fatalln(err)
	}
	error := json.Unmarshal(bodyLastItem, &ItemResponse)
	if error != nil {
		fmt.Println("Something happened unmarshalling: ", error)
	}
	switch ItemResponse.Type {
	case "story":
		story := model.Item.FormattedStory(ItemResponse)
		DisplayItem(story, ItemResponse.Type, lastItemURL)
	case "poll":
		poll := model.Item.FormattedPoll(ItemResponse)
		DisplayItem(poll, ItemResponse.Type, lastItemURL)
	case "job":
		job := model.Item.FormattedJob(ItemResponse)
		DisplayItem(job, ItemResponse.Type, lastItemURL)
	case "comment":
		comment := model.Item.FormattedComment(ItemResponse)
		DisplayItem(comment, ItemResponse.Type, lastItemURL)
	case "ask":
		ask := model.Item.FormattedAsk(ItemResponse)
		DisplayItem(ask, ItemResponse.Type, lastItemURL)
	default:
		fmt.Println("Returned no type! Is it possible?")
	}
}
