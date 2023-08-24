package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	savePath string
	rootCmd  = &cobra.Command{
		Use:   "gg",
		Short: "GoGraylog CLI",
		Long:  ``,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&savePath, "savePath", "/tmp/.gograylog", "Save path for persistant session data")

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(streamsCmd)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
