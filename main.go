package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hackernews cli application that written in Go!")
	response, err := http.Get("https://hacker-news.firebaseio.com/v0/maxitem.json")
	if err != nil {
		log.Fatalln(err)
	}
	GetLastItem(*response)

}

// GetLastItem returns the ID of last posted item
func GetLastItem(response http.Response) {
	returnedData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var maxItem int
	json.Unmarshal(returnedData, &maxItem)
	fmt.Println(maxItem)
}
