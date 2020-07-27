package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL helps us to create customizable request
const BaseURL = "https://hacker-news.firebaseio.com/v0/item/"

func main() {
	fmt.Println("Hackernews cli application that written in Go!")
	response, err := http.Get("https://hacker-news.firebaseio.com/v0/maxitem.json")
	if err != nil {
		log.Fatalln(err)
	}
	GetMaxID(*response)

}

// GetMaxID returns the ID of last posted item
func GetMaxID(response http.Response) {
	returnedData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var maxItem int
	json.Unmarshal(returnedData, &maxItem)
}

// GetLastItem brings the last posted item which can be in different types like story, job, comment etc.
func GetLastItem(maxID int) {
	lastItem := fmt.Sprintf("%s%d%s", BaseURL, maxID, ".json")

}
