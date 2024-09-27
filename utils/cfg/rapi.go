package cfg

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/shiron-dev/rapi/utils"
)

type RapiDependenciesConfig struct {
	Name       string
	Path       string
	Follow     bool
	AutoUpdate bool
	NoParam    bool
}
type RapiDependency struct {
	Origin   string
	Alias    string
	Template []RapiDependenciesConfig
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
	Dependencies []RapiDependency
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

func SaveConfig() error {
	data, err := yaml.Marshal(Config)
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	cfgPath := filepath.Join(wd, utils.RAPI_DIR, utils.RAPI_CONFIG)
	return os.WriteFile(cfgPath, data, 0644)
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
