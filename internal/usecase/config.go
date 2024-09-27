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
	config, err := c.files.LoadConfig()
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
	err = c.files.SaveConfig(*config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *ConfigUsecaseImpl) GetRapiConfig() (*domain.RapiConfig, error) {
	config, err := c.files.LoadConfig()
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
		err = c.files.SaveConfig(*config)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}
