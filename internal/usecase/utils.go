package usecase

// import (
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// func GetRapiWorkingDir() (string, error) {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		return "", err
// 	}

// 	for {
// 		if _, err := os.Stat(filepath.Join(dir, RAPI_DIR, RAPI_CONFIG)); err == nil {
// 			break
// 		}

// 		parent := filepath.Dir(dir)
// 		if parent == dir {
// 			err = os.ErrNotExist
// 			break
// 		}
// 		dir = parent
// 	}
// 	return dir, err
// }

// func GetOriginName(origin string) (string, string) {
// 	// TODO: aliaの考慮
// 	parsed := strings.Split(origin, "/")
// 	for i, p := range parsed {
// 		if strings.Contains(p, ".") {
// 			path := strings.Join(parsed[i+1:], "/")
// 			return p + "/" + path, path
// 		}
// 	}
// 	return ORIGIN_DEFAULT_HOST + "/" + origin, origin
// }

// func OriginToUrl(origin string) string {
// 	return "https://" + origin
// }

// func downloadOrigin(origin string) {
// 	wd, err := GetRapiWorkingDir()
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}

// 	_, originAlias := GetOriginName(origin)

// 	packagePath := filepath.Join(wd, RAPI_DIR, RAPI_PACKAGE_DIR, originAlias)

// 	if _, err := os.Stat(packagePath); err != nil {
// 		os.MkdirAll(packagePath, 0755)
// 	}

// 	run.GitClone(OriginToUrl(origin), packagePath, 1)
// }

// func AddUseTemplate(origin string, template string, local string) {
// 	// TODO: aliasの考慮
// 	dep := getOriginDependency(origin)
// 	if dep == nil {
// 		addOriginDependency(origin)
// 		dep = getOriginDependency(origin)
// 	}
// 	if dep == nil {
// 		fmt.Fprintln(os.Stderr, "Failed to add dependency")
// 		os.Exit(1)
// 	}

// 	// BUG: 配列のポインタだと参照しない？
// 	dep.Template = append(dep.Template, cfg.RapiDependenciesConfig{
// 		Name:       template,
// 		Path:       local,
// 		Follow:     true,
// 		AutoUpdate: true,
// 		NoParam:    false,
// 	})

// 	fmt.Printf("XXXX %+v\n", dep.Template)

// 	err := cfg.SaveConfig()
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}
// }

// func getOriginDependency(origin string) *cfg.RapiDependency {
// 	for _, dep := range cfg.Config.Dependencies {
// 		if dep.Origin == origin {
// 			return &dep
// 		}
// 	}
// 	return nil
// }

// func addOriginDependency(origin string) {
// 	dep := getOriginDependency(origin)
// 	if dep != nil {
// 		return
// 	}
// 	originName, originAlias := GetOriginName(origin)
// 	cfg.Config.Dependencies = append(cfg.Config.Dependencies, cfg.RapiDependency{
// 		Origin:   originName,
// 		Alias:    originAlias,
// 		Template: []cfg.RapiDependenciesConfig{},
// 	})
// 	downloadOrigin(origin)
// }
