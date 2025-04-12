package modules

import (
	"github.com/kelseyhightower/envconfig"
)

const DocsGeneratorModulePrefix = "DOCS_GENERATOR"

type DocsGenerator struct {
	TamplatePath string `envconfig:"TEMPLATE_PATH"`
}

func LoadDocsGenerator() (*DocsGenerator, error) {
	var dg DocsGenerator
	err := envconfig.Process(DocsGeneratorModulePrefix, &dg)

	if err != nil {
		return nil, err
	}

	return &dg, nil
}
