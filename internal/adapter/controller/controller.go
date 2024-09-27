package controller

import (
	"os"

	"github.com/shiron-dev/rapi/internal/usecase"
)

type Controller interface {
	PresistPreRun()
}

type ControllerImpl struct {
	config usecase.ConfigUsecase
	logger usecase.LoggerUsecase
}

func NewController(config usecase.ConfigUsecase, logger usecase.LoggerUsecase) Controller {
	return &ControllerImpl{config: config, logger: logger}
}

func (c *ControllerImpl) PresistPreRun() {
	isExists, err := c.config.ExistsRapiConfig()
	if err != nil {
		c.logger.ErrorWithErr(err)
	}

	if isExists {
		c.logger.Error("No config file found.")
		c.logger.Error("Please run `rapi init` to create a new config file.")
		os.Exit(1)
	}
}
