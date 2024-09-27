package controller

import (
	"fmt"
	"os"

	"github.com/shiron-dev/rapi/internal/usecase"
)

type ControllerImpl struct {
	config *usecase.ConfigUsecaseImpl
}

func NewControllerImpl() *ControllerImpl {
	return &ControllerImpl{}
}

func (c *ControllerImpl) PresistPreRun() {
	isExists, err := c.config.ExistsRapiConfig()
	if isExists {
		fmt.Fprintln(os.Stderr, "No config file found.")
		fmt.Fprintln(os.Stderr, "Please run `rapi init` to create a new config file.")
		os.Exit(1)
	}
}
