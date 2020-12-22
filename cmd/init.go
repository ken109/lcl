package cmd

import (
	"github.com/fatih/color"
	"github.com/ken109/lcl/util"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init lcl command",
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("Initializing...")

		_ = util.TryCommand("docker", "network", "create", "develop")
		_ = util.TryCommand("docker", "volume", "create", "mysql")
		_ = util.TryCommand("docker", "volume", "create", "mysql5")
		_ = util.TryCommand("docker", "volume", "create", "mongo")
		_ = util.TryCommand("docker", "pull", "ken109/dns:latest")
		_ = util.TryCommand("docker", "pull", "ken109/django:latest")
		_ = util.TryCommand("docker", "pull", "ken109/laravel:latest")

		color.Green("Completed.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
