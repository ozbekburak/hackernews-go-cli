package main

import (
	"flag"

	"github.com/hackernews-go-cli/client"
)

func main() {
	var storyCount int
	getLastItem := flag.Bool("last", false, "display last post, it can be story, comment, job, poll or ask")
	flag.IntVar(&storyCount, "top", 1, "display top stories, you can pass arguments until 500")
	flag.Parse()

	if *getLastItem {
		client.GetLastItem()
	}

	if storyCount > 0 && storyCount <= 500 {
		if !*getLastItem {
			client.GetTopStories(storyCount)
		}
	}
}
