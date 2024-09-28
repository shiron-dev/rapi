package usecase

import (
	"fmt"
	"path"

	"github.com/shiron-dev/rapi/configs"
	"github.com/shiron-dev/rapi/internal/adapter/repository"
)

type FilesUsecase interface {
	GetWD() (string, error)
	GetRapiDir() (string, error)
	MakeRapiDir() (string, error)
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

func (f *FilesUsecaseImpl) GetRapiDir() (string, error) {
	return f.files.GetRapiDir()
}

func (f *FilesUsecaseImpl) MakeRapiDir() (string, error) {
	wd, err := f.files.GetWD()
	if err != nil {
		return "", err
	}

	path := path.Join(wd, repository.RapiDirName)
	err = f.files.MkdirAll(path, 0755)
	if err != nil {
		return "", err
	}
	return path, nil
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
