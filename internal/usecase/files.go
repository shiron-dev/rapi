package usecase

import (
	"fmt"

	"github.com/shiron-dev/rapi/configs"
	"github.com/shiron-dev/rapi/internal/adapter/repository"
)

type FilesUsecase interface {
	GetWD() (string, error)
	MakeIgnoreFile() error
}

type FilesUsecaseImpl struct {
	files repository.FilesRepository
}

func NewFilesUsecase(files repository.FilesRepository) FilesUsecase {
	return &FilesUsecaseImpl{files: files}
}

func (f *FilesUsecaseImpl) GetWD() (string, error) {
	return f.files.GetWD()
}

func (f *FilesUsecaseImpl) MakeIgnoreFile() error {
	fmt.Println(configs.GitIgnore)
	fmt.Println("MakeIgnoreFile")
	_, err := f.files.WriteFileRapiDir("/.gitignore", configs.GitIgnore)
	if err != nil {
		return err
	}
	return nil
}
