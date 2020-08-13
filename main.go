package main

import (
	"flag"

	"github.com/hackernews-go-cli/client"
)

func main() {
	var storyCount int
	var askCount int
	var jobCount int
	var showCount int
	//getLastItem := flag.Bool("last", false, "display last post, it can be story, comment, job, poll or ask")
	flag.IntVar(&storyCount, "top", 1, "display top stories, you can pass arguments until 500")
	flag.IntVar(&askCount, "ask", 1, "display ask stories, you can pass arguments until 200")
	flag.IntVar(&jobCount, "job", 1, "display job stories, you can pass arguments until 200")
	flag.IntVar(&showCount, "show", 1, "display job stories, you can pass arguments until 200")

	flag.Parse()

	//if *getLastItem {
	//	client.GetLastItem()
	//}

	if storyCount > 0 && storyCount <= 500 {
		client.GetStories(storyCount, "story")
	}
	if askCount > 0 && askCount <= 200 {
		client.GetStories(askCount, "ask")
	}
	if jobCount > 0 && jobCount <= 200 {
		client.GetStories(jobCount, "job")
	}
	if showCount > 0 && showCount <= 200 {
		client.GetStories(showCount, "show")
	}
}
