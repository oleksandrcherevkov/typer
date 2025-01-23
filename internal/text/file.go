package text

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func GetText(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", errors.New("file can not be opened")
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	s := scan.Scan()
	if !s {
		return "", errors.New("file contains no text")
	}

	trimmed := strings.TrimSpace(scan.Text())
	if len(trimmed) == 0 {
		return "", errors.New("file contains no text")
	}

	return trimmed, nil
}
