package cfg

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type RapiDependenciesConfig struct {
	path   string
	follow bool
}
type RapiNameDependencies map[string]RapiDependenciesConfig
type RapiDependencies map[string]RapiNameDependencies

type RapiPackageConfig struct {
	Name        string
	Author      string
	URL         string
	Version     string
	Description string
	License     string
}
type RapiCLIConfig struct {
	RapiVersion string
}
type RapiConfig struct {
	Package      RapiPackageConfig
	Rapi         RapiCLIConfig
	Dependencies RapiDependencies
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
		},
		Rapi: RapiCLIConfig{
			RapiVersion: "0.1.0", // TODO: get version from rapi
		},
	}, nil
}
