package usecase

import (
	"path"

	"github.com/kataras/iris/v12/x/errors"
	"github.com/shiron-dev/rapi/internal/adapter/repository"
	"github.com/shiron-dev/rapi/internal/domain"
)

type ConfigUsecaseImpl struct {
	Files *repository.FilesRepositoryImpl
}

func NewConfigUsecaseImpl(files *repository.FilesRepositoryImpl) *ConfigUsecaseImpl {
	return &ConfigUsecaseImpl{Files: files}
}

func (c *ConfigUsecaseImpl) GetOrNewRapiConfig() (*domain.RapiConfig, error) {
	config, err := c.Files.LoadConfig()
	if err != nil {
		if errors.Is(err, repository.ConfigNotFoundError) {

		} else {
			return nil, err
		}
	}

	if config == nil {
		wd, err := c.Files.GetWD()
		if err != nil {
			return nil, err
		}
		config = domain.NewRapiConfig(path.Base(wd))
		err = c.Files.SaveConfig(*config)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}
