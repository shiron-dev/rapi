package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/shiron-dev/rapi/files"
	"github.com/shiron-dev/rapi/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Set up Rapi in the current directory",
	Run:   run,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func run(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Set up a Rapi in", wd)

	cfgPath := filepath.Join(wd, utils.RAPI_DIR, utils.RAPI_CONFIG)
	if _, err := os.Stat(cfgPath); err == nil {
		fmt.Println("Already initialized")
		os.Exit(0)
	}

	tmpl, err := template.New(cfgPath).Parse(string(files.RapiYaml))
	if err != nil {
		panic(err)
	}
	writer := new(strings.Builder)

	type PackageData struct {
		Name        string
		Author      string
		URL         string
		Version     string
		Description string
		License     string
	}
	packageData := &PackageData{
		Name:        filepath.Base(wd),
		Author:      "",
		URL:         "",
		Version:     "0.1.0",
		Description: "",
		License:     "",
	}

	tmpl.Execute(writer, packageData)

	fmt.Println("Wrote to", cfgPath)

	if err := os.MkdirAll(filepath.Join(wd, utils.RAPI_DIR), 0755); err != nil {
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

}
