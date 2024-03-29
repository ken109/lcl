package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var laStartCmd = &cobra.Command{
	Use:     "la [project name]",
	Aliases: []string{"laravel"},
	Short:   "laravel: start",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a project name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		start("laravel", args[0])
	},
}

func init() {
	startCmd.AddCommand(laStartCmd)
}
