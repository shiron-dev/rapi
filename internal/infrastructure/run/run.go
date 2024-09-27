package run

import (
	"os/exec"
	"strings"
)

const (
	Git = iota
)

func GetCommandName(command int) string {
	switch command {
	case Git:
		return "git"
	}
	return ""
}

func RunOutErr(cmd *exec.Cmd) string {
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}

func RunOutFunc(cmd *exec.Cmd, f func(error) string) string {
	out, err := cmd.Output()
	if err != nil {
		return f(err)
	}
	return strings.TrimSpace(string(out))
}

func RunOutDef(cmd *exec.Cmd, def string) string {
	out, err := cmd.Output()
	if err != nil {
		return def
	}
	return strings.TrimSpace(string(out))
}

func RunErr(cmd *exec.Cmd) {
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func RunFunc(cmd *exec.Cmd, f func(error)) {
	err := cmd.Run()
	if err != nil {
		f(err)
	}
}
