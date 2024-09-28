package repository

import (
	"errors"
	"os"
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
	GetRapiDir() (string, error)
	MkdirAll(path string, perm os.FileMode) error
	LoadConfig(path string) (*domain.RapiConfig, error)
	SaveConfig(path string, config domain.RapiConfig) error
	WriteFileRapiDir(filename string, data []byte) (string, error)
}

type FilesRepositoryImpl struct {
	files infra.FilesInterface
}

var ErrorConfigNotFound = errors.New("config file not found")

func NewFilesRepository(files infra.FilesInterface) FilesRepository {
	return &FilesRepositoryImpl{files: files}
}

func (c *FilesRepositoryImpl) GetRapiDir() (string, error) {
	wd, err := c.files.GetWD()
	if err != nil {
		return "", err
	}

	for {
		if c.files.Exists(filepath.Join(wd, RapiDirName, ConfigFileName)) {
			break
		}

		parent := filepath.Dir(wd)
		if parent == wd {
			return "", ErrorConfigNotFound
		}
		wd = parent
	}

	return filepath.Join(wd, RapiDirName), nil
}

func (c *FilesRepositoryImpl) GetWD() (string, error) {
	return c.files.GetWD()
}

func (c *FilesRepositoryImpl) MkdirAll(path string, perm os.FileMode) error {
	return c.files.MkdirAll(path, perm)
}

func (c *FilesRepositoryImpl) LoadConfig(path string) (*domain.RapiConfig, error) {
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

func (c *FilesRepositoryImpl) SaveConfig(path string, config domain.RapiConfig) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	rapiPath, err := c.GetRapiDir()
	if err != nil {
		return err
	}

	err = c.files.MkdirAll(rapiPath, 0755)
	if err != nil {
		return err
	}

	return c.files.WriteFile(path, data, 0644)
}

func (c *FilesRepositoryImpl) WriteFileRapiDir(filename string, data []byte) (string, error) {
	rapiPath, err := c.GetRapiDir()
	if err != nil {
		return "", err
	}

	err = c.files.MkdirAll(rapiPath, 0755)
	if err != nil {
		return "", err
	}

	path := filepath.Join(rapiPath, filename)
	err = c.files.WriteFile(path, data, 0644)
	if err != nil {
		return "", err
	}

	return path, nil
}
