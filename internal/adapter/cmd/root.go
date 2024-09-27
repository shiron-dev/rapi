package cmd

import (
	"os"

	"github.com/spf13/cobra"
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
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func presistPreRun(cmd *cobra.Command, args []string) {
	if cmd.Name() == "init" || isHelpCommand(cmd, args) {
		return
	}

	controller.PresistPreRun()
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
