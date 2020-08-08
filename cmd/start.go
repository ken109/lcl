package cmd

import (
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ken109/lcl/util"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var notEmpty bool
var share bool

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start environment",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.PersistentFlags().BoolVar(&notEmpty, "not-empty", config.Option.Start.NotEmpty, "Execute even if the current directory is not empty")
	startCmd.PersistentFlags().BoolVar(&share, "share", config.Option.Start.Share, "Share in the local network")
}

func start(framework string, project string) {
	var host string
	emptyCheck(notEmpty)
	color.Green("Copying docker-compose.yml...")
	host = copyCompose(brewPrefix+"/etc/lcl/"+framework+".yml", project, share)
	color.Green("Creating database...")
	createDB(project, config.Mysql.User, config.Mysql.Password)

	color.Green("Starting...")
	composeUp()
	color.Yellow("URL: http://" + host)
	color.Green("Completed.")
}

func emptyCheck(notEmpty bool) {
	var fileCount = 0
	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		fileCount++
		return nil
	})

	if fileCount > 1 && !notEmpty {
		color.Red("Current directory is not empty.")
		os.Exit(1)
	}
}

func copyCompose(path string, project string, share bool) string {
	ft, err := os.Open(path)
	if err != nil {
		color.Red("Could not read " + path)
		os.Exit(1)
	}
	defer ft.Close()
	b, _ := ioutil.ReadAll(ft)
	var compose = string(b)
	var host string
	compose = strings.ReplaceAll(compose, "APP_NAME", project)
	if share {
		var ifconfig, err = exec.Command("ifconfig").Output()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		pt := regexp.MustCompile("inet (192.168.[0-9]{1,3}.[0-9]{1,3})")
		group := pt.FindSubmatch(ifconfig)

		if len(group) != 2 {
			color.Red("Could not get IP address")
			os.Exit(1)
		}

		var ip = strings.Split(string(group[1]), ".")
		compose = strings.ReplaceAll(compose, "HOST_NAME", project+"-"+ip[2]+ip[3])
		host = project + "-" + ip[2] + ip[3]
	} else {
		compose = strings.ReplaceAll(compose, "HOST_NAME", project)
		host = project
	}
	fc, err := os.Create("./docker-compose.yml")
	if err != nil {
		color.Red("Could not read template")
		fmt.Println(err)
		os.Exit(1)
	}
	defer fc.Close()
	fc.WriteString(compose)

	return host
}

func createDB(project string, user string, password string) {
	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS `" + project + "`")
	if err != nil {
		fmt.Println(err)
		color.Red("Could not create database")
		os.Exit(1)
	}
}

func composeUp() {
	err := util.TryCommand("docker-compose", "up", "-d", "--remove-orphans")
	if err != nil {
		color.Red("Could not start docker-compose")
		os.Exit(1)
	}
}
