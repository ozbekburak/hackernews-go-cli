package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetTopStories returns top 100 story
func GetTopStories(num int, response http.Response) {
	topStory, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var topStories []int64
	json.Unmarshal(topStory, &topStories)
	for i := 0; i < num; i++ {
		fmt.Println(topStories[i])
	}
}
