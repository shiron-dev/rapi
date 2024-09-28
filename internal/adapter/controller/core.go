package controller

import (
	"errors"
	"os"

	"github.com/shiron-dev/rapi/internal/adapter/repository"
	"github.com/shiron-dev/rapi/internal/usecase"
)

type CoreController interface {
	InitRapi() error
}

type CoreControllerImpl struct {
	config usecase.ConfigUsecase
	files  usecase.FilesUsecase
	logger usecase.LoggerUsecase
}

func NewCoreController(config usecase.ConfigUsecase, files usecase.FilesUsecase, logger usecase.LoggerUsecase) CoreController {
	return &CoreControllerImpl{config: config, files: files, logger: logger}
}

func (c *CoreControllerImpl) InitRapi() error {
	ok, err := c.config.ExistsRapiConfig()
	if err != nil && !errors.Is(err, repository.ErrorConfigNotFound) {
		c.logger.ErrorWithErr(err)
	}
	if ok {
		c.logger.Info("Already initialized\n")
		os.Exit(0)
	}

	path, err := c.files.MakeRapiDir()
	if err != nil {
		c.logger.ErrorWithErr(err)
	}
	c.logger.Info("Created a new Rapi directory at %s\n", path)

	_, err = c.config.MakeNewRapiConfig()
	if err != nil {
		c.logger.ErrorWithErr(err)
	}

	err = c.files.MakeIgnoreFile()
	if err != nil {
		c.logger.ErrorWithErr(err)
	}

	return nil
}
