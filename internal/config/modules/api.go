package modules

import (
	"github.com/kelseyhightower/envconfig"
)

const ApiModulePrefix = "API"

type Api struct {
	HttpPort string `envconfig:"HTTP_PORT"`
	GRPCPort string `envconfig:"GRPC_PORT"`
	Host     string `envconfig:"HOST"`
}

func LoadApi() (*Api, error) {
	var s Api
	err := envconfig.Process(ApiModulePrefix, &s)

	if err != nil {
		return nil, err
	}

	return &s, nil
}
