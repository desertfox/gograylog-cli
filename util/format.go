package util

import (
	"bytes"
	"encoding/json"
)

func PrettyString(b []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, b, "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
