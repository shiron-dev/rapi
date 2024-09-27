package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/shiron-dev/rapi/internal/usecase"
	"github.com/shiron-dev/rapi/internal/usecase/cfg"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rapi",
	Short: "Rapi is a cli tool to manage templates",
	Long: `Rapi is a cli tool for managing templates.

Like a package manager, you can use only the parts you need from a scrap of templates hosted in another location.
When a dependent template is updated, Rapi will notify you of the update.`,
	PersistentPreRun: presistPreRun,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	wr, _ := usecase.GetRapiWorkingDir()
	cfgPath := filepath.Join(wr, usecase.RAPI_DIR, usecase.RAPI_CONFIG)

	cfgData, err := os.ReadFile(cfgPath)
	if err != nil {
		return
	}

	if err := cfg.LoadConfig(cfgData); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", cfgPath)
	} else {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}
}

func presistPreRun(cmd *cobra.Command, args []string) {
	if cmd.Name() == "init" || isHelpCommand(cmd, args) {
		return
	}

	wd, _ := os.Getwd()
	cfgPath := filepath.Join(wd, usecase.RAPI_DIR, usecase.RAPI_CONFIG)
	if _, err := os.Stat(cfgPath); err != nil {
		fmt.Fprintln(os.Stderr, "No config file found.")
		fmt.Fprintln(os.Stderr, "Please run `rapi init` to create a new config file.")
		os.Exit(1)
	}
}

func isHelpCommand(cmd *cobra.Command, args []string) bool {
	if cmd.Name() == "help" {
		return true
	}
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	if help, _ := cmd.Flags().GetBool("help"); help {
		return true
	}
	return false
}
