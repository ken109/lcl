package cmd

import (
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/ken109/lcl/util"
	"github.com/spf13/cobra"
)

var ssh string
var domain string

var stgCmd = &cobra.Command{
	Use:     "stg",
	Aliases: []string{"staging"},
	Short:   "Staging environment",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(stgCmd)

	stgCmd.PersistentFlags().StringVar(&ssh, "ssh", config.Option.Staging.Ssh, "SSH connection information")
	stgCmd.PersistentFlags().StringVar(&domain, "domain", config.Option.Staging.Domain, "Base domain")
}

func staging(framework string, keyFile string, project string) {
	if !util.Exists(keyFile) {
		color.Red("Current directory is not a " + framework + " project.")
		os.Exit(1)
	}
	if util.Exists("srv-" + project + ".tar.gz") {
		util.Remove("srv-" + project + ".tar.gz")
	}
	color.Green("Compressing...")
	tar(project)
	color.Green("Transferring...")
	scp(project)
	color.Green("Starting...")
	stg(framework, project)
	util.Remove("../srv-" + project + ".tar.gz")
	color.Blue("URL: https://" + project + "." + domain)
	color.Green("Completed.")
}

func tar(project string) {
	pwd := strings.Split(util.Pwd(), "/")
	wd := pwd[len(pwd)-1]

	util.Cd("..")
	if err := util.TryCommand("bash", "-c", "tar acf srv-"+project+".tar.gz -C "+wd+" ."); err != nil {
		color.Red("Could not create archive.")
		util.Cd(wd)
		os.Exit(1)
	}
	util.Cd(wd)
}

func scp(project string) {
	if err := util.TryCommand("scp", "../srv-"+project+".tar.gz", ssh+":/tmp"); err != nil {
		color.Red("Could not transfer.")
		os.Exit(1)
	}
}

func stg(framework string, project string) {
	if err := util.TryCommand(
		"bash", "-c",
		"echo srv stg "+framework+" "+project+" | ssh -t "+ssh,
	); err != nil {
		color.Red("Could not start.")
		os.Exit(1)
	}
}
