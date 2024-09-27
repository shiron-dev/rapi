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

type FilesRepository interface {
	GetWD() (string, error)
	LoadConfig() (*domain.RapiConfig, error)
	SaveConfig(config domain.RapiConfig) error
}

type FilesRepositoryImpl struct {
	files infra.FilesInterface
}

var ErrorConfigNotFound = errors.New("config file not found")

func NewFilesRepository(files infra.FilesInterface) FilesRepository {
	return &FilesRepositoryImpl{files: files}
}

func (c *FilesRepositoryImpl) getRapiDir() (string, error) {
	wd, err := c.files.GetWD()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, RapiDirName), nil
}

func (c *FilesRepositoryImpl) GetWD() (string, error) {
	return c.files.GetWD()
}

func (c *FilesRepositoryImpl) LoadConfig() (*domain.RapiConfig, error) {
	rapiPath, err := c.getRapiDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(rapiPath, ConfigFileName)
	if !c.files.Exists(path) {
		return nil, ErrorConfigNotFound
	}

	data, err := c.files.ReadFile(path)
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

	err = c.files.MkdirAll(rapiPath, 0755)
	if err != nil {
		return err
	}

	path := filepath.Join(rapiPath, ConfigFileName)
	return c.files.WriteFile(path, data, 0644)
}
