package main

import (
	"flag"
	"fmt"

	"github.com/hackernews-go-cli/client"
)

var (
	storyNum int
	askNum   int
	jobNum   int
	showNum  int
)

func main() {
	// maybe it can be done with sub-commands
	flag.IntVar(&storyNum, "top", 0, "display top stories, you can pass arguments until 500")
	flag.IntVar(&askNum, "ask", 0, "display ask stories, you can pass arguments until 200")
	flag.IntVar(&jobNum, "job", 0, "display job stories, you can pass arguments until 200")
	flag.IntVar(&showNum, "show", 0, "display job stories, you can pass arguments until 200")

	flag.Parse()

	if storyNum > 0 && storyNum <= 500 {
		client.FetchStories(storyNum, "topstories")
	}
	if askNum > 0 && askNum <= 200 {
		client.FetchStories(askNum, "askstories")
	}
	if jobNum > 0 && jobNum <= 200 {
		client.FetchStories(jobNum, "jobstories")
	}
	if showNum > 0 && showNum <= 200 {
		client.FetchStories(showNum, "showstories")
	} else if (storyNum <= 0 || storyNum > 500) || (askNum <= 0 || askNum > 200) || (jobNum <= 0 || jobNum > 200) || (showNum <= 0 || showNum > 200) {
		fmt.Println("No argument found or you exceed the max number! use --help to learn more")
	}

}
