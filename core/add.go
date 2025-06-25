package core

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func parseUserInput(args []string) Item {
	// validate
	if len(args) < 3 {
		log.Fatal().Msgf("Please provide category, item and date")
	}

	// parse
	category := args[0]
	item := args[1]
	date := args[2] // expiration date

	quantity := 1
	var err error
	if len(args) == 4 {
		quantity, err = strconv.Atoi(args[3])
		if err != nil {
			log.Fatal().Msgf("Quantity must be an integer, currently got %s", args[3])
		}
	}

	// category: validate
	isValidCategory := slices.Contains(AppConfig.Categories, category)
	if !isValidCategory {
		log.Fatal().Msgf("Invalid category: %s. Avaliable categories: %s", category, AppConfig.Categories)
	}

	// date: validate
	_, err = time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal().Msgf("Invalid date: %s", date)
	}

	return Item{
		Category: category,
		Item:     item,
		Date:     date,
		Quantity: quantity,
	}
}

func writeTask(item Item) string {
	// create string to write
	text := fmt.Sprintf("%s - %s - %s - %v\n", item.Category, item.Item, item.Date, item.Quantity)

	// open file
	f, err := os.OpenFile(AppConfig.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Error().Msgf("Error closing file: %v", err)
		}
	}(f)

	// write to file
	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}

	// return value
	return text
}

func Add(args []string) {
	item := parseUserInput(args)

	text := writeTask(item)
	fmt.Printf("Added: %s\n", text)
}
