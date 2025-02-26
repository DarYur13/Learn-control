package modules

import (
	"github.com/kelseyhightower/envconfig"
)

const PgModulePrefix = "PG"

type Pg struct {
	Host     string `envconfig:"HOST"`
	Port     string `envconfig:"PORT" default:"5432"`
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
	Database string `envconfig:"DATABASE" default:"learn-control"`
}

func LoadPg() (*Pg, error) {
	var s Pg
	err := envconfig.Process(PgModulePrefix, &s)

	if err != nil {
		return nil, err
	}

	return &s, nil
}
