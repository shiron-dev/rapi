// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"github.com/shiron-dev/rapi/internal/adapter/controller"
	"github.com/shiron-dev/rapi/internal/adapter/repository"
	"github.com/shiron-dev/rapi/internal/infrastructure/infra"
	"github.com/shiron-dev/rapi/internal/usecase"
)

// Injectors from wire.go:

func InitializeControllerSet() (*ControllersSet, error) {
	filesInterface := infra.NewFilesInterface()
	filesRepository := repository.NewFilesRepository(filesInterface)
	configUsecase := usecase.NewConfigUsecase(filesRepository)
	loggerInterface := infra.NewLoggerInterface()
	loggerRepository := repository.NewLoggerRepository(loggerInterface)
	loggerUsecase := usecase.NewLoggerUsecase(loggerRepository)
	controllerController := controller.NewController(configUsecase, loggerUsecase)
	controllersSet := &ControllersSet{
		Controller: controllerController,
	}
	return controllersSet, nil
}

// wire.go:

// Adapter
var repositorySet = wire.NewSet(repository.NewFilesRepository, repository.NewLoggerRepository)

var controllerSet = wire.NewSet(controller.NewController)

// Infrastructure
var infrastructureSet = wire.NewSet(infra.NewFilesInterface, infra.NewLoggerInterface)

// Usecase
var usecaseSet = wire.NewSet(usecase.NewConfigUsecase, usecase.NewLoggerUsecase)

type ControllersSet struct {
	Controller controller.Controller
}