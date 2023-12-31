package cmd

import (
	"fmt"
	"os"
	"syscall"

	"github.com/desertfox/gograylog"
	"github.com/desertfox/gograylog-cli/util"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	host     string
	username string
	loginCmd = &cobra.Command{
		Use: "login",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Please enter the password for user %v on host %v:\n", username, host)
			bytepw, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			c := util.BuildClient(util.Session{
				Host:    host,
				Session: gograylog.Session{},
			})

			err = c.Login(username, string(bytepw))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if err := util.SaveToDisk(savePath, host, *c.Session); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	loginCmd.PersistentFlags().StringVar(&host, "host", "", "host")
	loginCmd.PersistentFlags().StringVar(&username, "username", "", "username")
	loginCmd.MarkPersistentFlagRequired("host")
	loginCmd.MarkPersistentFlagRequired("username")
}
