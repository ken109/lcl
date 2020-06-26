package cmd

import (
	"github.com/spf13/cobra"
)

var baseCmd = &cobra.Command{
	Use:   "base",
	Short: "base(nginx-proxy, mysql, redis, dns): start, stop",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(baseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// baseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// baseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
