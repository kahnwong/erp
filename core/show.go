package core

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rs/zerolog/log"
)

var (
	Green   = color.New(color.FgHiGreen).SprintFunc()
	Red     = color.New(color.FgRed).SprintFunc()
	Yellow  = color.New(color.FgYellow).SprintFunc()
	Blue    = color.New(color.FgBlue).SprintFunc()
	Cyan    = color.New(color.FgCyan).SprintFunc()
	Magenta = color.New(color.FgHiMagenta).SprintFunc()
)

func readData() ([]Item, error) {
	var erpData []Item

	// read data
	file, err := openFile(AppConfig.Path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Error().Err(err).Msg("failed to close file")
		}
	}(file)

	s := bufio.NewScanner(file)

	// main
	for s.Scan() {
		line := s.Text()
		line = strings.TrimRight(line, "\r\n")
		data := strings.Split(line, " - ")

		if len(data) < 4 {
			continue // skip malformed lines
		}

		category := data[0]
		item := data[1]
		date := data[2]
		quantity, err := strconv.Atoi(data[3])
		if err != nil {
			return nil, fmt.Errorf("invalid quantity in line '%s': %w", line, err)
		}

		// append
		erpData = append(erpData, Item{category, item, date, quantity})
	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return erpData, nil
}

func Show() error {
	// parse data
	data, err := readData()
	if err != nil {
		return fmt.Errorf("failed to read data: %w", err)
	}

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
		dateObject, err := time.Parse("2006-01-02", date)
		if err != nil {
			return fmt.Errorf("invalid date format '%s': %w", date, err)
		}
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
	return nil
}
