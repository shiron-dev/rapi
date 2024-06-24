package run

import (
	"golang.org/x/sys/execabs"
)

func GitClone(origin string, path string) error {
	cmd := execabs.Command(GetCommandName(Git), "clone", origin, path)
	return cmd.Run()
}
