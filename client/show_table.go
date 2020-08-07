package client

import (
	"fmt"
	"html"

	"github.com/alexeyco/simpletable"
	"github.com/hackernews-go-cli/model"
)

// ShowTable creates a table with related data
func ShowTable(item []model.Item) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Type"},
			{Align: simpletable.AlignCenter, Text: "By"},
			{Align: simpletable.AlignCenter, Text: "Score"},
			{Align: simpletable.AlignCenter, Text: "Time"},
			{Align: simpletable.AlignCenter, Text: "Text"},
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "URL"},
		},
	}
	for _, x := range item {
		r := []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", x.ID)},
			{Text: x.Type},
			{Text: x.By},
			{Text: fmt.Sprintf("%d", x.Score)},
			{Text: fmt.Sprintf("%v", x.FormattedTime(x.Time))},
			{Text: fmt.Sprintf("%v", html.UnescapeString(x.Text))},
			{Text: x.Title},
			{Text: x.URL},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleRounded)
	fmt.Println(table)
}
