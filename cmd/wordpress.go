package cmd

import (
	"github.com/spf13/cobra"
)

var wordpressCmd = &cobra.Command{
	Use:     "wordpress",
	Aliases: []string{"wp"},
	Short:   "wordpress: start, stop, staging",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(wordpressCmd)
}
