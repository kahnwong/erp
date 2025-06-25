package core

import (
	"bufio"
	"os"
	"sort"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Item struct {
	Category string
	Item     string
	Date     string
}

func readData() []Item {
	var erpData []Item

	// read data
	file := openFile(AppConfig.Path)
	s := bufio.NewScanner(file)

	// main
	for s.Scan() {
		line := s.Text()
		line = strings.TrimRight(line, "\r\n")
		data := strings.Split(line, " - ")

		category := data[0]
		item := data[1]
		date := data[2]

		// append
		erpData = append(erpData, Item{category, item, date})
	}

	return erpData
}

func Show() {
	// parse data
	data := readData()

	// sort by expiration date
	sort.Slice(data, func(i, j int) bool {
		return data[i].Date < data[j].Date
	})

	// render table
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Category", "Item", "Date Expired"})

	for _, item := range data {
		t.AppendRows([]table.Row{
			{item.Category, item.Item, item.Date},
		})
	}
	t.Render()
}
