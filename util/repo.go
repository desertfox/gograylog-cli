package util

import (
	"encoding/json"
	"os"

	"github.com/desertfox/gograylog"
)

type Session struct {
	Host    string            `json:"host"`
	Session gograylog.Session `json:"session"`
}

func SaveToDisk(path, host string, session gograylog.Session) error {
	cliSession := Session{host, session}

	b, err := json.Marshal(cliSession)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return err
}

func ReadFromDisk(path string) (Session, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return Session{}, err
	}
	var session Session
	err = json.Unmarshal(file, &session)

	return session, nil
}
