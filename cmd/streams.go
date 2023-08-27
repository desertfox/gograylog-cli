package cmd

import (
	"fmt"
	"os"

	"github.com/desertfox/gograylog-cli/util"
	"github.com/spf13/cobra"
)

var (
	streamsCmd = &cobra.Command{
		Use: "streams",
		Run: func(cmd *cobra.Command, args []string) {
			s, err := util.ReadFromDisk(savePath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			c := util.BuildClient(s)

			b, err := c.Streams()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			data, err := util.PrettyString(b)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(data)
		},
	}
)
