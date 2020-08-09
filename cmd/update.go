package cmd

import (
	"github.com/fatih/color"
	"github.com/ken109/lcl/util"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update lcl docker images",
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("Updating...")

		_ = util.TryCommand("docker", "pull", "ken109/dns:latest")
		_ = util.TryCommand("docker", "pull", "ken109/django:latest")
		_ = util.TryCommand("docker", "pull", "ken109/laravel:latest")

		color.Green("Completed.")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
