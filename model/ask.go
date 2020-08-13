package model

// Ask struct represents the model of our response.
type Ask struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Text        string `json:"text"`
	Time        int64  `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}
