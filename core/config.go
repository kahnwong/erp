package core

import (
	cliBase "github.com/kahnwong/cli-base"
)

type Config struct {
	Path       string   `json:"path"`
	Categories []string `json:"categories"`
}

var AppConfig = cliBase.ReadYaml[Config]("~/.config/erp/config.yaml") // init

// models
type Item struct {
	Category string
	Item     string
	Date     string
	Quantity int
}
