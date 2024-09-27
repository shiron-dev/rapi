package infra

import (
	"os"
)

type FilesInterfaceImpl struct{}

func NewFilesInterfaceImpl() FilesInterfaceImpl {
	return FilesInterfaceImpl{}
}

func (f FilesInterfaceImpl) GetWD() (string, error) {
	return os.Getwd()
}

func (f FilesInterfaceImpl) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (f FilesInterfaceImpl) WriteFile(path string, data []byte, perm os.FileMode) error {
	return os.WriteFile(path, data, perm)
}
