package repository

import (
	"errors"
	"path/filepath"

	"github.com/shiron-dev/rapi/internal/domain"
	"github.com/shiron-dev/rapi/internal/infrastructure/infra"
	"gopkg.in/yaml.v2"
)

const (
	RapiDirName    = ".rapi"
	ConfigFileName = "rapi.yaml"
)

type FilesRepositoryImpl struct {
	Files *infra.FilesInterfaceImpl
}

var ConfigNotFoundError = errors.New("config file not found")

func NewFilesRepositoryImpl(files *infra.FilesInterfaceImpl) *FilesRepositoryImpl {
	return &FilesRepositoryImpl{Files: files}
}

func (c *FilesRepositoryImpl) getRapiDir() (string, error) {
	wd, err := c.Files.GetWD()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, RapiDirName), nil
}

func (c *FilesRepositoryImpl) GetWD() (string, error) {
	return c.Files.GetWD()
}

func (c *FilesRepositoryImpl) LoadConfig() (*domain.RapiConfig, error) {
	rapiPath, err := c.getRapiDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(rapiPath, ConfigFileName)
	if !c.Files.Exists(path) {
		return nil, ConfigNotFoundError
	}

	data, err := c.Files.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config domain.RapiConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *FilesRepositoryImpl) SaveConfig(config domain.RapiConfig) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	rapiPath, err := c.getRapiDir()
	if err != nil {
		return err
	}

	err = c.Files.MkdirAll(rapiPath, 0755)
	if err != nil {
		return err
	}

	path := filepath.Join(rapiPath, ConfigFileName)
	return c.Files.WriteFile(path, data, 0644)
}
