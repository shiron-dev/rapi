package repository

import (
	"path/filepath"

	"github.com/shiron-dev/rapi/internal/infrastructure/infra"
	"github.com/shiron-dev/rapi/internal/infrastructure/model"
	"gopkg.in/yaml.v2"
)

const (
	RapiDirName    = ".rapi"
	ConfigFileName = "rapi.yaml"
)

type FilesRepositoryImpl struct {
	Files *infra.FilesInterfaceImpl

	Config *model.RapiConfig
}

func NewFilesRepositoryImpl() *FilesRepositoryImpl {
	return &FilesRepositoryImpl{}
}

func (c FilesRepositoryImpl) getRapiDir() (string, error) {
	wd, err := c.Files.GetWD()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, RapiDirName), nil
}

func (c FilesRepositoryImpl) newConfig() (*model.RapiConfig, error) {
	wd, err := c.Files.GetWD()
	if err != nil {
		return nil, err
	}

	return &model.RapiConfig{
		Package: model.RapiPackageConfig{
			Name:    filepath.Base(wd),
			Version: "1.0.0",
		},
		Rapi: model.RapiCLIConfig{
			RapiVersion: "1.0.0",
		},
		Dependencies: []model.RapiDependency{},
	}, nil
}

func (c FilesRepositoryImpl) LoadConfig() (*model.RapiConfig, error) {
	if c.Config != nil {
		return c.Config, nil
	}

	rapiPath, err := c.getRapiDir()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(rapiPath, ConfigFileName)
	data, err := c.Files.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config model.RapiConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	c.Config = &config

	return c.Config, nil
}

func (c FilesRepositoryImpl) SaveConfig() error {
	data, err := yaml.Marshal(c.Config)
	if err != nil {
		return err
	}

	rapiPath, err := c.getRapiDir()
	if err != nil {
		return err
	}
	path := filepath.Join(rapiPath, ConfigFileName)
	return c.Files.WriteFile(path, data, 0644)
}
