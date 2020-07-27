package model

import "fmt"

// Item struct represents the model of our response.
type Item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Text        string `json:"text"`
	Parts       []int  `json:"parts"`
	Parent      int    `json:"parent"`
	Time        int64  `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

// FormattedStory formats our response to show meaningful output to user
func (i Item) FormattedStory() string {
	story := fmt.Sprintf(
		"\nBy: %s\nID: %d\nComment Count: %d\nScore: %d\nTitle: %s\nURL: %s\n",
		i.By, i.ID, i.Descendants, i.Score, i.Title, i.URL)
	return story
}
