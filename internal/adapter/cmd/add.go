package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add [remote] [template] [local path]",
	Short:   "Add the specified template.",
	Run:     controller.RunAdd,
	Aliases: []string{"use"},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
