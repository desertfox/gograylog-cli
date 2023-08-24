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
	streamid  string
	frequency int
	limit     int
	searchCmd = &cobra.Command{
		Use: "search",
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

			q := gograylog.Query{
				StreamID:    streamid,
				QueryString: args[0],
				Frequency:   frequency,
				Fields:      []string{"source"}, //TODO
				Limit:       limit,
			}

			b, err := c.Search(q)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(string(b))
		},
	}
)

func init() {
	searchCmd.PersistentFlags().StringVar(&streamid, "streamid", "", "streamid")
	searchCmd.PersistentFlags().IntVar(&frequency, "frequency", 900, "frequency")
	searchCmd.PersistentFlags().IntVar(&limit, "limit", 10000, "limit")
	loginCmd.MarkPersistentFlagRequired("streamid")
}
