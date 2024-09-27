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
var cmdSet = wire.NewSet()

var repositorySet = wire.NewSet(
	repository.NewFilesRepositoryImpl,
	repository.NewLoggerRepositoryImpl,
)

var controllerSet = wire.NewSet(
	controller.NewControllerImpl,
)

// Infrastructure
var infrastructureSet = wire.NewSet(
	infra.NewFilesInterfaceImpl,
	infra.NewLoggerInterfaceImpl,
)

// Usecase
var usecaseSet = wire.NewSet(
	usecase.NewConfigUsecaseImpl,
)

type ControllerSet struct {
	ControllerImpl *controller.ControllerImpl
}

func InitializeControllerSet() (*ControllerSet, error) {
	wire.Build(
		cmdSet,
		repositorySet,
		controllerSet,
		infrastructureSet,
		usecaseSet,
		wire.Struct(new(ControllerSet), "*"),
	)
	return &ControllerSet{}, nil
}
