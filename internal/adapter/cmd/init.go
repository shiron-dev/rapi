package cmd

import (
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Set up Rapi in the current directory",
	Run:   controller.RunInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
