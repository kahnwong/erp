package core

import (
	cliBase "github.com/kahnwong/cli-base"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Path       string   `json:"path"`
	Categories []string `json:"categories"`
}

var AppConfig *Config

func init() {
	var err error
	AppConfig, err = cliBase.ReadYaml[Config]("~/.config/erp/config.yaml")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read config")
	}
}

// models
type Item struct {
	Category string
	Item     string
	Date     string
	Quantity int
}
