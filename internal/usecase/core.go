package usecase

import (
	"os"
)

type CoreUsecase interface {
	InitRapi() error
}

type CoreUsecaseImpl struct {
	config ConfigUsecase
	files  FilesUsecase
	logger LoggerUsecase
}

func NewCoreUsecase(config ConfigUsecase, files FilesUsecase, logger LoggerUsecase) CoreUsecase {
	return &CoreUsecaseImpl{config: config, files: files, logger: logger}
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
