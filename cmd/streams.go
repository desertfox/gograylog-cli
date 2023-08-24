package cmd

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/desertfox/gograylog"
	"github.com/desertfox/gograylog-cli/token"
	"github.com/spf13/cobra"
)

var (
	streamsCmd = &cobra.Command{
		Use: "streams",
		Run: func(cmd *cobra.Command, args []string) {
			h, t, err := token.ReadFromDisk(savePath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			c := gograylog.Client{
				Host:  h,
				Token: t,
				HttpClient: &http.Client{
					Transport: &http.Transport{
						TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
					},
				},
			}

			b, err := c.Streams()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(string(b))
		},
	}
)
