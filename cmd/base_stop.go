package cmd

import (
	"github.com/fatih/color"
	"github.com/ken109/lcl/util"
	"github.com/spf13/cobra"
)

var baseStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "base: stop",
	Run: func(cmd *cobra.Command, args []string) {
		var dir = util.Pwd()
		util.Cd(brewPrefix + "/etc/lcl/base")
		color.Green("Stopping...")
		if err := util.TryCommand("docker-compose", "down"); err != nil {
			color.Red("Failed to stop.")
		} else {
			color.Green("Successfully stopped.")
		}
		util.Cd(dir)
	},
}

func init() {
	baseCmd.AddCommand(baseStopCmd)
}
