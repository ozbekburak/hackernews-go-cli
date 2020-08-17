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
	flag.IntVar(&storyNum, "top", 1, "display top stories, you can pass arguments until 500")
	flag.IntVar(&askNum, "ask", 1, "display ask stories, you can pass arguments until 200")
	flag.IntVar(&jobNum, "job", 1, "display job stories, you can pass arguments until 200")
	flag.IntVar(&showNum, "show", 1, "display job stories, you can pass arguments until 200")

	flag.Parse()

	if storyNum > 0 && storyNum <= 500 {

		client.FetchStories(storyNum, "topstories")
		fmt.Println(storyNum)
	}
	if askNum > 0 && askNum <= 200 {
		client.FetchStories(askNum, "askstories")
	}
	if jobNum > 0 && jobNum <= 200 {
		client.FetchStories(jobNum, "jobstories")
	}
	if showNum > 0 && showNum <= 200 {
		client.FetchStories(showNum, "showstories")
	}

}
