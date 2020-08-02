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
	fmt.Printf("Type: %s\nURL: %s\n-----------%s", typeOfPost, itemURL, output)
}

// GetMaxItemID returns the ID of last posted item
func GetMaxItemID(response http.Response) int {
	var maxItemID int
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(data, &maxItemID)
	return maxItemID
}

// GetLastItem brings the last posted item which can be in different types like story, job, comment etc.
func GetLastItem() {
	lastItemURL := "https://hacker-news.firebaseio.com/v0/maxitem.json"
	response, err := http.Get(lastItemURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	maxID := GetMaxItemID(*response)
	GetItem(maxID)
}

// GetItem function returns the item that related ID
func GetItem(itemID int) {
	var ItemResponse model.Item
	itemURL := fmt.Sprintf("%s%d%s", BaseURL, itemID, ".json")
	item, err := http.Get(itemURL)

	if err != nil {
		log.Fatalln(err)
	}
	defer item.Body.Close()
	itemBody, err := ioutil.ReadAll(item.Body)
	if err != nil {
		log.Fatalln(err)
	}
	error := json.Unmarshal(itemBody, &ItemResponse)
	if error != nil {
		fmt.Println("Something bad happened: ", error)
	}
	switch ItemResponse.Type {
	case "story":
		story := model.Item.FormattedStory(ItemResponse)
		DisplayItem(story, ItemResponse.Type, itemURL)
	case "poll":
		poll := model.Item.FormattedPoll(ItemResponse)
		DisplayItem(poll, ItemResponse.Type, itemURL)
	case "job":
		job := model.Item.FormattedJob(ItemResponse)
		DisplayItem(job, ItemResponse.Type, itemURL)
	case "comment":
		comment := model.Item.FormattedComment(ItemResponse)
		DisplayItem(comment, ItemResponse.Type, itemURL)
	case "ask":
		ask := model.Item.FormattedAsk(ItemResponse)
		DisplayItem(ask, ItemResponse.Type, itemURL)
	default:
		fmt.Println("Returned no type! Is it possible?")
	}
}
