package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Set up Rapi in the current directory",
	Run:   runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Set up a Rapi in", wd)

	cfgPath := filepath.Join(wd, usecase.RAPI_DIR, usecase.RAPI_CONFIG)
	if _, err := os.Stat(cfgPath); err == nil {
		fmt.Println("Already initialized")
		os.Exit(0)
	}

	tmpl, err := template.New(cfgPath).Parse(string(usecase.RapiYaml))
	if err != nil {
		panic(err)
	}
	writer := new(strings.Builder)

	newCfg, err := cfg.NewConfig()
	if err != nil {
		panic(err)
	}
	tmpl.Execute(writer, newCfg)

	fmt.Println("Wrote to", cfgPath)

	if err := os.MkdirAll(filepath.Join(wd, usecase.RAPI_DIR), 0755); err != nil {
		panic(err)
	}
	f, err := os.Create(cfgPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.Write([]byte(writer.String())); err != nil {
		panic(err)
	}
	println("\n" + writer.String())

	ignorePath := filepath.Join(wd, usecase.RAPI_DIR, ".gitignore")
	f, err = os.Create(ignorePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.Write([]byte(usecase.GitIgnore)); err != nil {
		panic(err)
	}

}
