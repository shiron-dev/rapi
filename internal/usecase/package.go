package usecase

import (
	"errors"
	"regexp"
	"strings"

	"github.com/shiron-dev/rapi/internal/domain"
)

type PackageUsecase interface {
	MakeRapiDependencyObj(config *domain.RapiConfig, name string, path string, params map[string]string) error
}

type PackageUsecaseImpl struct {
}

var ErrorInvalidDependencyName = errors.New("Invalid dependency name")
var ErrorDuplicateDependency = errors.New("Duplicate dependency")
var ErrorUnknownDependency = errors.New("Unknown dependency")

func NewPackageUsecase() PackageUsecase {
	return &PackageUsecaseImpl{}
}

func removeScheme(name string) string {
	re := regexp.MustCompile(`^[a-zA-Z]+://`)
	return re.ReplaceAllString(name, "")
}

func (p *PackageUsecaseImpl) getOriginDepsName(name string) (string, string, error) {
	name = removeScheme(name)
	name = strings.TrimPrefix(name, "/")
	name = strings.TrimSuffix(name, "/")

	parts := strings.Split(name, "/")
	if !strings.Contains(parts[0], ".") {
		parts = append([]string{domain.OriginDefaultHost}, parts...)
	}

	if len(parts) < 3 {
		return "", "", ErrorInvalidDependencyName
	}

	prefix := strings.Join(parts[:3], "/")
	suffix := strings.Join(parts[3:], "/")

	return prefix, suffix, nil
}

func (p *PackageUsecaseImpl) MakeRapiDependencyObj(config *domain.RapiConfig, name string, path string, params map[string]string) error {
	origin, oPath, err := p.getOriginDepsName(name)
	if err != nil {
		return ErrorInvalidDependencyName
	}
	if oPath == "" {
		return ErrorInvalidDependencyName
	}

	if path == "" {
		path = oPath
	}

	deps := config.Dependencies.Templates[origin]
	if deps == nil {
		deps = &domain.RapiDependencyOrigin{
			Origin: origin,
		}
		config.Dependencies.Templates[origin] = deps
	}

	parts := strings.Split(path, "/")
	depsPath := config.Dependencies.Paths
	for i, part := range parts {
		if i != len(parts)-1 {
			if depsPath.Paths[part] == nil {
				depsPath.Paths[part] = &domain.RapiDependencyPath{}
			}
			depsPath = depsPath.Paths[part]
		} else {
			if depsPath.Files[part] != nil {
				return ErrorDuplicateDependency
			}

			depsPath.Files[part] = &domain.RapiDependencyPathTemplate{
				Origin: origin,
				Url:    path,
			}
			return nil
		}
	}

	return ErrorUnknownDependency
}
