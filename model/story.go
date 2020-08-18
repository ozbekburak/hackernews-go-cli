package model

import (
	// html.UnescapeString unescapes entities like "&lt;" to become "<".
	"strconv"
	"time"
)

// Story struct represents the model of our response.
type Story struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int64  `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

// FormattedTime function converts epoch time to human readable format
func (s Story) FormattedTime(epoch int64) string {
	strTime := strconv.FormatInt(epoch, 10)
	itemDate, err := strconv.ParseInt(strTime, 10, 64)
	if err != nil {
		panic(err)
	}
	t := time.Unix(itemDate, 0)
	return t.Format("15:04")
}
