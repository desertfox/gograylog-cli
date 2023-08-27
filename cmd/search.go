package cmd

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/desertfox/gograylog-cli/util"
	"github.com/spf13/cobra"
)

var (
	streamid  string
	frequency int
	limit     int
	fields    string
	quiet     bool
	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "search --flag value \"graylog search query\"",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("must pass query string")
				os.Exit(1)

			}
			s, err := util.ReadFromDisk(savePath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			c := util.BuildClient(s)

			q := util.BuildQuery(streamid, args[0], frequency, strings.Split(fields, ","), limit)

			b, err := c.Search(q)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(string(b))
			records, err := csv.NewReader(bytes.NewBuffer(b)).ReadAll()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if !quiet {
				fmt.Printf("Count: %d\n", len(records))
				if len(records) > 0 {
					fmt.Printf("First record: %s\n", records[1][:])
					fmt.Printf("Last  record: %s\n", records[len(records)-1][:])
				}
			}
		},
	}
)

func init() {
	searchCmd.PersistentFlags().StringVar(&streamid, "streamid", "", "graylog streamid")
	searchCmd.PersistentFlags().IntVar(&frequency, "frequency", 900, "search frequency in seconds")
	searchCmd.PersistentFlags().IntVar(&limit, "limit", 10000, "limit of messages to return")
	searchCmd.PersistentFlags().StringVar(&fields, "fields", "timestamp,message", "csv header fields to be returned")
	searchCmd.PersistentFlags().StringVar(&fields, "fields", "timestamp,message", "csv header fields to be returned")
	searchCmd.PersistentFlags().BoolVar(&quiet, "quiet", false, "suppress extra info")
	loginCmd.MarkPersistentFlagRequired("streamid")
}
