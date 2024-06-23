package utils

import (
	"os"
	"path/filepath"
)

func GetRapiWorkingDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, RAPI_DIR, RAPI_CONFIG)); err == nil {
			break
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			err = os.ErrNotExist
			break
		}
		dir = parent
	}
	return dir, err
}
