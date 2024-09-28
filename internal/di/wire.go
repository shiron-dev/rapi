//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/shiron-dev/rapi/internal/adapter/controller"
	"github.com/shiron-dev/rapi/internal/adapter/repository"
	"github.com/shiron-dev/rapi/internal/infrastructure/infra"
	"github.com/shiron-dev/rapi/internal/usecase"
)

// Adapter
var repositorySet = wire.NewSet(
	repository.NewFilesRepository,
	repository.NewLoggerRepository,
)

var controllerSet = wire.NewSet(
	controller.NewController,
	controller.NewCoreController,
)

// Infrastructure
var infrastructureSet = wire.NewSet(
	infra.NewFilesInterface,
	infra.NewLoggerInterface,
)

// Usecase
var usecaseSet = wire.NewSet(
	usecase.NewConfigUsecase,
	usecase.NewFilesUsecase,
	usecase.NewLoggerUsecase,
)

type ControllersSet struct {
	Controller controller.Controller
}

func InitializeControllerSet() (*ControllersSet, error) {
	wire.Build(
		repositorySet,
		controllerSet,
		infrastructureSet,
		usecaseSet,
		wire.Struct(new(ControllersSet), "*"),
	)
	return nil, nil
}
