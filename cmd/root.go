package cmd

import (
	"fmt"
	"github.com/ken109/lcl/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var config util.Config
var brewPrefix = util.GetOutput("brew", "--prefix")

var rootCmd = &cobra.Command{
	Use:   "lcl",
	Short: "Build local environment",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "",
		fmt.Sprintf("config file (default is %s)", brewPrefix+"/etc/lcl/config.yaml"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile(brewPrefix + "/etc/lcl/config.yml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config file Read error")
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("config file Unmarshal error")
		fmt.Println(err)
		os.Exit(1)
	}
}
