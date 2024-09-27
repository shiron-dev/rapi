package controller

import (
	"os"

	"github.com/shiron-dev/rapi/internal/usecase"
)

type Controller interface {
	PresistPreRun()
	RunInit()
}

type ControllerImpl struct {
	core   usecase.CoreUsecase
	config usecase.ConfigUsecase
	logger usecase.LoggerUsecase
	files  usecase.FilesUsecase
}

func NewController(core usecase.CoreUsecase, config usecase.ConfigUsecase, logger usecase.LoggerUsecase, files usecase.FilesUsecase) Controller {
	return &ControllerImpl{core: core, config: config, logger: logger, files: files}
}

func (c *ControllerImpl) PresistPreRun() {
	ok, err := c.config.ExistsRapiConfig()
	if err != nil {
		c.logger.ErrorWithErr(err)
	}

	if ok {
		c.logger.Error("No config file found.\n")
		c.logger.Error("Please run `rapi init` to create a new config file.\n")
		os.Exit(1)
	}
}

func (c *ControllerImpl) RunInit() {
	wd, err := c.files.GetWD()
	if err != nil {
		c.logger.ErrorWithErr(err)
	}
	c.logger.Info("Set up a Rapi in %s\n", wd)

	c.core.InitRapi()
	c.logger.Info("Initialized\n")
}
