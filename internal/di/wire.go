//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/shiron-dev/rapi/internal/adapter/repository"
	"github.com/shiron-dev/rapi/internal/infrastructure/infra"
)

// Adapter
var cmdSet = wire.NewSet()

var repositorySet = wire.NewSet(
	repository.NewFilesRepositoryImpl,
)

// Infrastructure
var infrastructureSet = wire.NewSet(
	infra.NewFilesInterfaceImpl,
)

// Usecase
var usecaseSet = wire.NewSet()
