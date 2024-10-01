package controller

import (
	"errors"
	"fmt"
	"os"

	"github.com/shiron-dev/rapi/internal/adapter/repository"
	"github.com/shiron-dev/rapi/internal/usecase"
	"github.com/spf13/cobra"
)

type Controller interface {
	PresistPreRun()
	RunInit(cmd *cobra.Command, args []string)
	RunAdd(cmd *cobra.Command, args []string)
}

type ControllerImpl struct {
	core CoreController

	config usecase.ConfigUsecase
	logger usecase.LoggerUsecase
	files  usecase.FilesUsecase
	pkg    usecase.PackageUsecase
}

func NewController(core CoreController, config usecase.ConfigUsecase, logger usecase.LoggerUsecase, files usecase.FilesUsecase, pkg usecase.PackageUsecase) Controller {
	return &ControllerImpl{core: core, config: config, logger: logger, files: files, pkg: pkg}
}

func (c *ControllerImpl) PresistPreRun() {
	ok, err := c.config.ExistsRapiConfig()
	if err != nil && !errors.Is(err, repository.ErrorConfigNotFound) {
		c.logger.ErrorWithErr(err)
	}

	if !ok {
		c.logger.Error("No config file found.\n")
		c.logger.Error("Please run `rapi init` to create a new config file.\n")
		os.Exit(1)
	}

	path, err := c.files.GetRapiDir()
	if err != nil {
		c.logger.ErrorWithErr(err)
	}
	c.logger.Info("Found a Rapi directory at %s\n", path)
}

func (c *ControllerImpl) RunInit(cmd *cobra.Command, args []string) {
	wd, err := c.files.GetWD()
	if err != nil {
		c.logger.ErrorWithErr(err)
	}
	c.logger.Info("Set up a Rapi in %s\n", wd)

	c.core.InitRapi()
	c.logger.Info("Initialized\n")
}

func (c *ControllerImpl) RunAdd(cmd *cobra.Command, args []string) {
	const (
		all = iota
		auto
		local
	)
	mode, err := func(cmd *cobra.Command, args []string) (uint, error) {
		switch len(args) {
		case 0:
			return all, nil
		case 1:
			return auto, nil
		case 2:
			return local, nil
		}
		return 0, fmt.Errorf("invalid arguments")
	}(cmd, args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if mode == auto {
		config, err := c.config.GetRapiConfig()
		if err != nil {
			c.logger.ErrorWithErr(err)
			os.Exit(1)
		}
		err = c.pkg.MakeRapiDependencyObj(config, args[0], "", make(map[string]string))
		if err != nil {
			c.logger.Error("Error: %s\n", err.Error())
			os.Exit(1)
		} else {
			c.logger.Info("Added %s\n", args[0])
		}
		c.config.SaveRapiConfig(config)
	}
}
