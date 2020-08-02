package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/alexeyco/simpletable"

	"github.com/hackernews-go-cli/client"
)

var (
	// data variable includes "id, by, url"
	data = [][]interface{}{
		{24028351, "ngram", "https://sliderulemuseum.com/SR_Course.htm"},
		{24028510, "bilinualcom", "https://www.gitenberg.org/"},
		{24027487, "DerCommodore", "https://gpu.rocks/"},
		{24027663, "kkm", "https://www.inkandswitch.com/local-first.html"},
		{24028289, "zarfius", "https://www.dylanwilson.net/what-i-learned-about-failing-from-my-5-year-indie-game-dev-project/"},
	}
)

func main() {
	var storyCount int
	getLastItem := flag.Bool("last", false, "display last post, it can be story, comment, job, poll or ask")
	flag.IntVar(&storyCount, "top", 0, "display top stories, you can pass arguments until 500")
	flag.Parse()

	if *getLastItem {
		client.GetLastItem()
	}

	if storyCount > 0 && storyCount <= 500 {
		client.GetTopStories(storyCount)
	} else {
		log.Fatalln("You can fetch max 500 item!")
	}
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "By"},
			{Align: simpletable.AlignCenter, Text: "URL"},
		},
	}
	for _, row := range data {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row[0].(int))},
			{Text: row[1].(string)},
			{Align: simpletable.AlignRight, Text: row[2].(string)},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	fmt.Println(table)
}
