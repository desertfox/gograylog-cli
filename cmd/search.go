package cmd

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/desertfox/gograylog"
	"github.com/desertfox/gograylog-cli/util"
	"github.com/spf13/cobra"
)

var (
	streamid  string
	frequency int
	limit     int
	fields    string
	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "search --flag value \"graylog search query\"",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("must pass query string")
				os.Exit(1)

			}
			h, t, err := util.ReadFromDisk(savePath)
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
				Fields:      strings.Split(fields, ","),
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
	searchCmd.PersistentFlags().StringVar(&streamid, "streamid", "", "graylog streamid")
	searchCmd.PersistentFlags().IntVar(&frequency, "frequency", 900, "search frequency in seconds")
	searchCmd.PersistentFlags().IntVar(&limit, "limit", 10000, "limit of messages to return")
	searchCmd.PersistentFlags().StringVar(&fields, "fields", "timestamp,message", "csv header fields to be returned")
	loginCmd.MarkPersistentFlagRequired("streamid")
}
