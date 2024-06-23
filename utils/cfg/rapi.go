package cfg

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/shiron-dev/rapi/utils"
)

type RapiDependenciesConfig struct {
	Path       string
	Follow     bool
	AutoUpdate bool
	NoParam    bool
}
type RapiNameDependencies map[string]RapiDependenciesConfig
type RapiDependency struct {
	Alias     string
	Templates RapiNameDependencies
}

type RapiPackageConfig struct {
	Name        string
	Author      string
	URL         string
	Version     string
	Description string
	License     string
	Recipe      string
}
type RapiCLIConfig struct {
	RapiVersion string
}
type RapiConfig struct {
	Package      RapiPackageConfig
	Rapi         RapiCLIConfig
	Dependencies map[string]RapiDependency
}

var Config *RapiConfig = nil

func LoadConfig(data []byte) error {
	var config RapiConfig
	err := yaml.Unmarshal(data, &config)
	if err == nil {
		Config = &config
	}
	return err
}

func NewConfig() (*RapiConfig, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return &RapiConfig{
		Package: RapiPackageConfig{
			Name:    filepath.Base(wd),
			Version: "0.1.0",
			Recipe:  filepath.Join(utils.RAPI_DIR, utils.RAPI_RECIPE_DIR),
		},
		Rapi: RapiCLIConfig{
			RapiVersion: "0.1.0", // TODO: get version from rapi
		},
	}, nil
}
