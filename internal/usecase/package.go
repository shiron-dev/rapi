package usecase

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shiron-dev/rapi/internal/infrastructure/run"
	"github.com/shiron-dev/rapi/internal/usecase/cfg"
)

func downloadOrigin(origin string) {
	wd, err := GetRapiWorkingDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, originAlias := GetOriginName(origin)

	packagePath := filepath.Join(wd, RAPI_DIR, RAPI_PACKAGE_DIR, originAlias)

	if _, err := os.Stat(packagePath); err != nil {
		os.MkdirAll(packagePath, 0755)
	}

	run.GitClone(OriginToUrl(origin), packagePath, 1)
}

func AddUseTemplate(origin string, template string, local string) {
	// TODO: aliasの考慮
	dep := getOriginDependency(origin)
	if dep == nil {
		addOriginDependency(origin)
		dep = getOriginDependency(origin)
	}
	if dep == nil {
		fmt.Fprintln(os.Stderr, "Failed to add dependency")
		os.Exit(1)
	}

	// BUG: 配列のポインタだと参照しない？
	dep.Template = append(dep.Template, cfg.RapiDependenciesConfig{
		Name:       template,
		Path:       local,
		Follow:     true,
		AutoUpdate: true,
		NoParam:    false,
	})

	fmt.Printf("XXXX %+v\n", dep.Template)

	err := cfg.SaveConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getOriginDependency(origin string) *cfg.RapiDependency {
	for _, dep := range cfg.Config.Dependencies {
		if dep.Origin == origin {
			return &dep
		}
	}
	return nil
}

func addOriginDependency(origin string) {
	dep := getOriginDependency(origin)
	if dep != nil {
		return
	}
	originName, originAlias := GetOriginName(origin)
	cfg.Config.Dependencies = append(cfg.Config.Dependencies, cfg.RapiDependency{
		Origin:   originName,
		Alias:    originAlias,
		Template: []cfg.RapiDependenciesConfig{},
	})
	downloadOrigin(origin)
}
