package core

import (
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/rs/zerolog/log"
)

func parseUserInput(args []string) (string, string, string) {
	// validate
	if len(args) != 3 {
		log.Fatal().Msgf("Please provide category, item and date")
	}

	// parse
	category := args[0]
	item := args[1]
	date := args[2] // expiration date

	// category: validate
	isValidCategory := slices.Contains(AppConfig.Categories, category)
	if !isValidCategory {
		log.Fatal().Msgf("Invalid category: %s. Avaliable categories: %s", category, AppConfig.Categories)
	}

	// date: validate
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal().Msgf("Invalid date: %s", date)
	}

	return category, item, date
}

func writeTask(category string, item string, date string) string {
	// create string to write
	text := fmt.Sprintf("%s - %s - %s\n", category, item, date)

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
	category, item, date := parseUserInput(args)

	text := writeTask(category, item, date)
	fmt.Printf("Added: %s\n", text)
}
