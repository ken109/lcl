package cmd

import (
	"github.com/fatih/color"
	"github.com/ken109/lcl/util"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
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
	if util.Exists(project + ".tar.gz") {
		util.Remove(project + ".tar.gz")
	}
	color.Green("Compressing...")
	tar(project)
	color.Green("Transferring...")
	scp(project)
	color.Green("Starting...")
	stg(framework, project)
	util.Remove(project + ".tar.gz")
	color.Blue("URL: https://" + project + "." + domain)
	color.Green("Completed.")
}

func tar(project string) {
	if err := exec.Command("bash", "-c", "tar acf "+project+".tar.gz ./*").Run(); err != nil {
		color.Red("Could not create archive.")
		os.Exit(1)
	}
}

func scp(project string) {
	if err := exec.Command("scp", project+".tar.gz", ssh+":~/.tmp").Run(); err != nil {
		color.Red("Could not transfer.")
		os.Exit(1)
	}
}

func stg(framework string, project string) {
	if err := exec.Command("bash", "-c",
		"echo srv start "+framework+" "+project+" | ssh -t "+ssh).Run(); err != nil {
		color.Red("Could not start.")
		os.Exit(1)
	}
}
