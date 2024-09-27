package files

import _ "embed"

//go:embed rapi.yaml
var RapiYaml []byte

// go:embed .gitignore
var GitIgnore []byte
