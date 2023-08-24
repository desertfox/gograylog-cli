package token

import (
	"fmt"
	"os"
	"strings"
)

func SaveToDisk(host, token, path string) error {
	str := fmt.Sprintf("%v\n%v", host, token)

	err := os.WriteFile(path, []byte(str), 0644)

	return err
}

func ReadFromDisk(path string) (string, string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}
	data := strings.Split(string(file), "\n")

	return data[0], data[1], nil
}
