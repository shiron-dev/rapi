package model

type RapiDependenciesConfig struct {
	Name       string
	Path       string
	Follow     bool
	AutoUpdate bool
	NoParam    bool
}
type RapiDependency struct {
	Origin   string
	Alias    string
	Template []RapiDependenciesConfig
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
	Package      RapiPackageConfig
	Rapi         RapiCLIConfig
	Dependencies []RapiDependency
}
