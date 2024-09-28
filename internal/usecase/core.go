package usecase

import (
	"os"
	"path"

	"github.com/shiron-dev/rapi/internal/adapter/repository"
)

type CoreUsecase interface {
	InitRapi() error
}

type CoreUsecaseImpl struct {
	config ConfigUsecase
	files  FilesUsecase
	logger LoggerUsecase

	filesRepo repository.FilesRepository
}

func NewCoreUsecase(config ConfigUsecase, files FilesUsecase, logger LoggerUsecase, filesRepo repository.FilesRepository) CoreUsecase {
	return &CoreUsecaseImpl{config: config, files: files, logger: logger, filesRepo: filesRepo}
}

func (c *CoreUsecaseImpl) InitRapi() error {
	ok, err := c.config.ExistsRapiConfig()
	if err != nil {
		c.logger.ErrorWithErr(err)
	}
	if ok {
		c.logger.Info("Already initialized\n")
		os.Exit(0)
	}

	wd, err := c.files.GetWD()
	if err != nil {
		return err
	}
	path := path.Join(wd, repository.RapiDirName)
	c.filesRepo.MkdirAll(path, 0755)

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
