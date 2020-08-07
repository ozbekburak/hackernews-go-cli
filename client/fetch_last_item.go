package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

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
	lastitem := GetItem(int64(maxID))
	ShowTable(lastitem)
}
