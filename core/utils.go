package core

import (
	"os"

	"github.com/rs/zerolog/log"
)

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal().Msg("Error opening file")
	}

	return file
}
