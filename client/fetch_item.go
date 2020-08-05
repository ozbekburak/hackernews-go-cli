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

// GetItem function returns the item that related ID
func GetItem(itemID int64) {
	var ItemResponse model.Item
	itemresponselist := []model.Item{}
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
	itemresponselist = append(itemresponselist, ItemResponse)
	ShowTable(itemresponselist)
}
