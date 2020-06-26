package cmd

import (
	"github.com/spf13/cobra"
)

var laravelCmd = &cobra.Command{
	Use:     "laravel",
	Aliases: []string{"la"},
	Short:   "laravel: start, stop, staging",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(laravelCmd)
}
