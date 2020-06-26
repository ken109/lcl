package cmd

import (
	"github.com/spf13/cobra"
)

var djangoCmd = &cobra.Command{
	Use:     "django",
	Aliases: []string{"dj"},
	Short:   "django: start, stop, staging",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(djangoCmd)
}
