package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetTopStories returns top stories
func GetTopStories(num int) {
	//var ItemResponse model.Item
	//itemresponselist := []model.Item{}
	//topStoriesList := []model.Item{}
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
			// itemURL := fmt.Sprintf("%s%d%s", BaseURL, topStories[i], ".json")
			// item, err := http.Get(itemURL)

			// if err != nil {
			// 	log.Fatalln(err)
			// }
			// defer item.Body.Close()
			// itemBody, err := ioutil.ReadAll(item.Body)
			// if err != nil {
			// 	log.Fatalln(err)
			// }
			// error := json.Unmarshal(itemBody, &ItemResponse)
			// if error != nil {
			// 	fmt.Println("Something bad happened: ", error)
			// }
			//element := GetItem(topStories[i][0])
			//topStoriesList = append(topStoriesList, GetItem(topStories[i]))
			GetItem(topStories[i])
			fmt.Println(topStories[i])
			fmt.Printf("%T", topStories[i])
		}
		//ShowTable(topStoriesList)
	}
}
