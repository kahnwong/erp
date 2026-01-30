package core

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

func parseUserInput(args []string) (Item, error) {
	// validate
	if len(args) < 3 {
		return Item{}, fmt.Errorf("please provide category, item and date")
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
			return Item{}, fmt.Errorf("quantity must be an integer, currently got %s: %w", args[3], err)
		}
	}

	// category: validate
	isValidCategory := slices.Contains(AppConfig.Categories, category)
	if !isValidCategory {
		return Item{}, fmt.Errorf("invalid category: %s. available categories: %v", category, AppConfig.Categories)
	}

	// date: validate
	_, err = time.Parse("2006-01-02", date)
	if err != nil {
		return Item{}, fmt.Errorf("invalid date: %s: %w", date, err)
	}

	return Item{
		Category: category,
		Item:     item,
		Date:     date,
		Quantity: quantity,
	}, nil
}

func writeTask(item Item) (string, error) {
	// create string to write
	text := fmt.Sprintf("%s - %s - %s - %v\n", item.Category, item.Item, item.Date, item.Quantity)

	// open file
	f, err := os.OpenFile(AppConfig.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer func(f *os.File) {
		if closeErr := f.Close(); closeErr != nil {
			err = fmt.Errorf("error closing file: %w", closeErr)
		}
	}(f)

	// write to file
	if _, err = f.WriteString(text); err != nil {
		return "", fmt.Errorf("failed to write to file: %w", err)
	}

	// return value
	return text, nil
}

func Add(args []string) error {
	item, err := parseUserInput(args)
	if err != nil {
		return err
	}

	text, err := writeTask(item)
	if err != nil {
		return err
	}

	fmt.Printf("Added: %s", text)
	return nil
}
