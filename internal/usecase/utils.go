package usecase

import (
	"os"
	"path/filepath"
	"strings"
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

func GetOriginName(origin string) (string, string) {
	// TODO: aliaの考慮
	parsed := strings.Split(origin, "/")
	for i, p := range parsed {
		if strings.Contains(p, ".") {
			path := strings.Join(parsed[i+1:], "/")
			return p + "/" + path, path
		}
	}
	return ORIGIN_DEFAULT_HOST + "/" + origin, origin
}

func OriginToUrl(origin string) string {
	return "https://" + origin
}
