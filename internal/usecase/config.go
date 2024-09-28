package usecase

import (
	"errors"
	"path"

	"github.com/shiron-dev/rapi/internal/adapter/repository"
	"github.com/shiron-dev/rapi/internal/domain"
)

type ConfigUsecase interface {
	ExistsRapiConfig() (bool, error)
	MakeNewRapiConfig() (*domain.RapiConfig, error)
	GetRapiConfig() (*domain.RapiConfig, error)
}

type ConfigUsecaseImpl struct {
	files repository.FilesRepository
}

func NewConfigUsecase(files repository.FilesRepository) ConfigUsecase {
	return &ConfigUsecaseImpl{files: files}
}

func (c *ConfigUsecaseImpl) ExistsRapiConfig() (bool, error) {
	rapiPath, err := c.files.GetRapiDir()
	if err != nil {
		return false, err
	}
	cfgPath := path.Join(rapiPath, repository.ConfigFileName)

	config, err := c.files.LoadConfig(cfgPath)
	if err != nil {
		if errors.Is(err, repository.ErrorConfigNotFound) {
			return false, nil
		}
		return false, err
	}

	return config != nil, nil
}

func (c *ConfigUsecaseImpl) MakeNewRapiConfig() (*domain.RapiConfig, error) {
	wd, err := c.files.GetWD()
	if err != nil {
		return nil, err
	}
	config := domain.NewRapiConfig(path.Base(wd))
	path := path.Join(wd, repository.RapiDirName, repository.ConfigFileName)
	err = c.files.SaveConfig(path, *config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *ConfigUsecaseImpl) GetRapiConfig() (*domain.RapiConfig, error) {
	rapiDir, err := c.files.GetRapiDir()
	if err != nil {
		return nil, err
	}
	cfgPath := path.Join(rapiDir, repository.ConfigFileName)

	config, err := c.files.LoadConfig(cfgPath)
	if err != nil {
		if errors.Is(err, repository.ErrorConfigNotFound) {

		} else {
			return nil, err
		}
	}

	if config == nil {
		wd, err := c.files.GetWD()
		if err != nil {
			return nil, err
		}
		config = domain.NewRapiConfig(path.Base(wd))
		err = c.files.SaveConfig(cfgPath, *config)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}
