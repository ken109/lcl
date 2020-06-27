package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var startCmd = &cobra.Command{
	Use:   "start [project name]",
	Short: "Start environment",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a project name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var host string
		if emptyOnly, err := cmd.PersistentFlags().GetBool("empty-only"); err == nil {
			EmptyCheck(emptyOnly)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
		if share, err := cmd.PersistentFlags().GetBool("share"); err == nil {
			color.Green("Copying docker-compose.yml...")
			host = CopyCompose(brewPrefix+"/etc/lcl/"+cmd.Parent().Name()+".yml", args[0], share)
		}
		color.Green("Creating database...")
		CreateDB(args[0], config.Mysql.User, config.Mysql.Password)

		color.Green("Starting...")
		ComposeUp()
		color.Blue("URL: http://" + host)
		color.Green("Completed.")
	},
}

func init() {
	wordpressCmd.AddCommand(startCmd)
	laravelCmd.AddCommand(startCmd)
	djangoCmd.AddCommand(startCmd)

	startCmd.PersistentFlags().Bool("empty-only", config.Option.Start.EmptyOnly, "Stop processing if the current directory is not empty")
	startCmd.PersistentFlags().Bool("share", config.Option.Start.Share, "Share in the office")
}

func EmptyCheck(emptyOnly bool) {
	var fileCount = 0
	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		fileCount++
		return nil
	})

	if fileCount > 1 && emptyOnly {
		color.Red("Current directory is not empty.")
		os.Exit(1)
	}
}

func CopyCompose(path string, project string, share bool) string {
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
		log.Fatal(err) //ファイルが開けなかったときエラー出力
	}
	defer fc.Close()
	fc.WriteString(compose)

	return host
}

func CreateDB(project string, user string, password string) {
	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS `" + project + "`")
	if err != nil {
		color.Red("Could not create database")
		os.Exit(1)
	}
}

func ComposeUp() {
	err := exec.Command("docker-compose", "up", "-d", "--remove-orphans").Run()
	if err != nil {
		color.Red("Could not start docker-compose")
		os.Exit(1)
	}
}
