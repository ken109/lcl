package cmd

import (
	"github.com/fatih/color"
	"github.com/ken109/lcl/util"
	"github.com/spf13/cobra"
)

var baseStartCmd = &cobra.Command{
	Use:   "start",
	Short: "base: start",
	Run: func(cmd *cobra.Command, args []string) {
		var dir = util.Pwd()
		util.Cd(brewPrefix + "/etc/lcl/base")
		color.Green("Starting...")

		composeArgs := append([]string{"up", "-d"}, args...)
		if err := util.TryCommand("docker-compose", composeArgs...); err != nil {
			color.Red("Failed to start.")
		} else {
			color.Green("Successfully started.")
		}
		util.Cd(dir)
	},
}

func init() {
	baseCmd.AddCommand(baseStartCmd)
}
