package text

import (
	"errors"
	"os"
	"strings"
)

func GetText(filePath string) (string, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return "", errors.New("file can not be read")
	}

	text := string(b)

	trimmed := strings.TrimSpace(text)
	if len(trimmed) == 0 {
		return "", errors.New("file contains no text")
	}

	return trimmed, nil
}
