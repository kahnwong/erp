package core

import (
	cliBase "github.com/kahnwong/cli-base"
)

type Config struct {
	Categories []string `json:"categories"`
}

var AppConfig = cliBase.ReadYaml[Config]("~/.config/erp/config.yaml") // init
