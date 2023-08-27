package util

import (
	"crypto/tls"
	"net/http"

	"github.com/desertfox/gograylog"
)

func BuildClient(s Session) gograylog.Client {
	return gograylog.Client{
		Host:    s.Host,
		Session: &s.Session,
		HttpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

func BuildQuery(streamid, query string, frequency int, fields []string, limit int) gograylog.Query {
	return gograylog.Query{
		StreamID:    streamid,
		QueryString: query,
		Frequency:   frequency,
		Fields:      fields,
		Limit:       limit,
	}
}
