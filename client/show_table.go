package client

import (
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/hackernews-go-cli/model"
)

// ShowTable creates a table with related data
func ShowTable(item model.Item) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Type"},
			{Align: simpletable.AlignCenter, Text: "By"},
			{Align: simpletable.AlignCenter, Text: "Descendants"},
			{Align: simpletable.AlignCenter, Text: "Score"},
			{Align: simpletable.AlignCenter, Text: "Text"},
			{Align: simpletable.AlignCenter, Text: "Parent"},
			{Align: simpletable.AlignCenter, Text: "Time"},
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "URL"},
		},
	}

	r := []*simpletable.Cell{
		{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", item.ID)},
		{Text: item.Type},
		{Text: item.By},
		{Text: fmt.Sprintf("%d", item.Descendants)},
		{Text: fmt.Sprintf("%d", item.Score)},
		{Text: item.Text},
		{Text: fmt.Sprintf("%d", item.Parent)},
		{Text: fmt.Sprintf("%v", item.FormattedTime(item.Time))},
		{Text: item.Title},
		{Align: simpletable.AlignRight, Text: item.URL},
	}

	table.Body.Cells = append(table.Body.Cells, r)

	table.SetStyle(simpletable.StyleRounded)
	fmt.Println(table)
}
