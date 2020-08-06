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
const BaseURL string = "https://hacker-news.firebaseio.com/v0/item/"

// GetItem function returns the item that related ID
func GetItem(itemID int64) {
	itemURL := fmt.Sprintf("%s%d%s", BaseURL, itemID, ".json")
	var itemResponse model.Item

	// To append items to table I had to use array type argument
	itemResponseArray := []model.Item{}

	item, err := http.Get(itemURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer item.Body.Close()

	itemBody, err := ioutil.ReadAll(item.Body)
	if err != nil {
		log.Fatalln(err)
	}

	error := json.Unmarshal(itemBody, &itemResponse)
	if error != nil {
		fmt.Println("Something bad happened: ", error)
	}
	itemResponseArray = append(itemResponseArray, itemResponse)
	ShowTable(itemResponseArray)
}
