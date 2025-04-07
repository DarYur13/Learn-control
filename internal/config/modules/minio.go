package modules

import (
	"github.com/kelseyhightower/envconfig"
)

const MinioModulePrefix = "MINIO"

type Minio struct {
	Host     string `envconfig:"HOST"`
	Port     string `envconfig:"PORT" default:"9000"`
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
}

func LoadMinio() (*Minio, error) {
	var m Minio
	err := envconfig.Process(MinioModulePrefix, &m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}
