package cmd

import (
	"fmt"
	"os"

	"github.com/shiron-dev/rapi/core"
	"github.com/shiron-dev/rapi/utils"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use [remote] [template] [local path]",
	Short: "Use the specified template.",
	Run:   runUse,
}

func init() {
	rootCmd.AddCommand(useCmd)
}

func runUse(cmd *cobra.Command, args []string) {
	const (
		all = iota
		auto
		local
	)
	mode, err := func() (uint, error) {
		switch len(args) {
		case 0:
			return all, nil
		case 2:
			return auto, nil
		case 3:
			return local, nil
		}
		return 0, fmt.Errorf("invalid arguments")
	}()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(mode)

	originPath, originAlias := utils.GetOriginName(args[0])

	if mode == local {
		core.AddUseTemplate(originPath, args[1], args[2])
	}
	println(originPath, originAlias)
}
