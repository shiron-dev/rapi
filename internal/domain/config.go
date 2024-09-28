package domain

type RapiDependencyOrigin struct {
	Origin  string
	Version string
}

type RapiDependencyPath struct {
	Paths map[string]*RapiDependencyPath
	Files map[string]*RapiDependencyPathTemplate
}

type RapiDependencyPathTemplate struct {
	Origin string
	Url    string
}

type RapiDependency struct {
	Templates map[string]*RapiDependencyOrigin
	Paths     *RapiDependencyPath
}

type RapiPackageConfig struct {
	Name        string
	Author      string
	URL         string
	Version     string
	Description string
	License     string
}
type RapiCLIConfig struct {
	RapiVersion string
}

type RapiConfig struct {
	Package      *RapiPackageConfig
	Rapi         *RapiCLIConfig
	Dependencies *RapiDependency
}

const OriginDefaultHost = "github.com"

func NewRapiConfig(rapiName string) *RapiConfig {
	return &RapiConfig{
		Package: &RapiPackageConfig{
			Name:    rapiName,
			Version: "0.0.1",
		},
		Rapi: &RapiCLIConfig{
			RapiVersion: RapiCLIVersion,
		},
		Dependencies: &RapiDependency{
			Templates: map[string]*RapiDependencyOrigin{},
			Paths: &RapiDependencyPath{
				Paths: map[string]*RapiDependencyPath{},
				Files: map[string]*RapiDependencyPathTemplate{},
			},
		},
	}
}
