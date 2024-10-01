package run

import (
	"golang.org/x/sys/execabs"
)

func GitClone(origin string, path string, depth uint) error {
	cmd := execabs.Command(GetCommandName(Git), "clone", origin, path, "--depth", string(depth))
	return cmd.Run()
}
