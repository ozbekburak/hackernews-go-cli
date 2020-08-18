package client

import (
	"fmt"
	"html"

	"github.com/hackernews-go-cli/model"
	"github.com/mitchellh/go-wordwrap"

	"github.com/alexeyco/simpletable"
)

// ShowTable creates a table with related data
func ShowTable(item []model.Story, storyType string) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			// {Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "By"},
			{Align: simpletable.AlignCenter, Text: "Score"},
			{Align: simpletable.AlignCenter, Text: "Time"},
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "URL"},
		},
	}

	for _, x := range item {
		r := []*simpletable.Cell{
			// {Text: fmt.Sprintf("%d", x.ID)},
			{Text: x.By},
			{Text: fmt.Sprintf("%d", x.Score)},
			{Text: fmt.Sprintf("%v", x.FormattedTime(x.Time))},
			{Text: fmt.Sprintf("%v", wordwrap.WrapString(html.UnescapeString(x.Title), 75))},
			// WrapString library works for only whitespace, so it can not be used for URL.
			// Maybe custom url shortener fit in here!
			{Text: x.URL}}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleRounded)
	fmt.Println(table)
}
