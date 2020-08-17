package main

import (
	"flag"

	"github.com/hackernews-go-cli/client"
)

var (
	storyNum int
	askNum   int
	jobNum   int
	showNum  int
)

func main() {
	flag.IntVar(&storyNum, "top", 0, "display top stories, you can pass arguments until 500")
	flag.IntVar(&askNum, "ask", 0, "display ask stories, you can pass arguments until 200")
	flag.IntVar(&jobNum, "job", 0, "display job stories, you can pass arguments until 200")
	flag.IntVar(&showNum, "show", 0, "display job stories, you can pass arguments until 200")

	flag.Parse()

	if storyNum > 0 && storyNum <= 500 {
		client.GetStories(storyNum, "topstories")
	}
	if askNum > 0 && askNum <= 200 {
		client.GetStories(askNum, "askstories")
	}
	if jobNum > 0 && jobNum <= 200 {
		client.GetStories(jobNum, "jobstories")
	}
	if showNum > 0 && showNum <= 200 {
		client.GetStories(showNum, "showstories")
	}
}
