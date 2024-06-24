package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shiron-dev/rapi/utils"
	"github.com/shiron-dev/rapi/utils/cfg"
	"github.com/shiron-dev/rapi/utils/run"
)

func DownloadOrigin(origin string) {
	wd, err := utils.GetRapiWorkingDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, originAlias := utils.GetOriginName(origin)

	packagePath := filepath.Join(wd, utils.RAPI_DIR, utils.RAPI_PACKAGE_DIR, originAlias)

	if _, err := os.Stat(packagePath); err != nil {
		os.MkdirAll(packagePath, 0755)
	}

	run.GitClone(utils.OriginToUrl(origin), packagePath)
}

func AddUseTemplate(origin string, template string, local string) {
	DownloadOrigin(origin)
	fmt.Println("AddUseTemplate")
	// _, originAlias := utils.GetOriginName(origin)

	fmt.Printf("%+v\n", cfg.Config.Dependencies)
	cfg.Config.Dependencies[origin].Template[template] = cfg.RapiDependenciesConfig{
		Path:       local,
		Follow:     true,
		AutoUpdate: true,
		NoParam:    false,
	}

	err := cfg.SaveConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
