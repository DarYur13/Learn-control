package modules

import (
	"github.com/kelseyhightower/envconfig"
)

const LogModulePrefix = "LOG"

type Log struct {
	Level    string `envconfig:"LEVEL" default:"info"`
	FilePath string `envconfig:"FILE" default:"stdout"`
}

func LoadLog() (*Log, error) {
	var s Log
	err := envconfig.Process(LogModulePrefix, &s)

	if err != nil {
		return nil, err
	}

	return &s, nil
}
