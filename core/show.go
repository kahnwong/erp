package core

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	Green   = color.New(color.FgHiGreen).SprintFunc()
	Red     = color.New(color.FgRed).SprintFunc()
	Yellow  = color.New(color.FgYellow).SprintFunc()
	Blue    = color.New(color.FgBlue).SprintFunc()
	Cyan    = color.New(color.FgCyan).SprintFunc()
	Magenta = color.New(color.FgHiMagenta).SprintFunc()
)

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
		quantity, _ := strconv.Atoi(data[3])

		// append
		erpData = append(erpData, Item{category, item, date, quantity})
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
	t.AppendHeader(table.Row{"Category", "Item", "Date Expired", "Quantity"})

	for _, item := range data {
		/// date: assign color
		date := item.Date
		today := time.Now()
		dateObject, _ := time.Parse("2006-01-02", date)
		timeDiff := dateObject.Sub(today)

		if timeDiff.Hours() < 14*24 { // 2 weeks
			date = Red(date)
		} else if timeDiff.Hours() < 30*24 { // 1 month
			date = Yellow(date)
		} else {
			date = Green(date)
		}

		t.AppendRows([]table.Row{
			{Blue(item.Category), Cyan(item.Item), date, Magenta(item.Quantity)},
		})
	}
	t.Render()
}
