package codefile

import (
	"bufio"
	"os"
	"strings"
)

// ScanFile reads a file and returns its lines as a slice of strings
func ScanFile(filePath string, maxLines int) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
		if len(lines) >= maxLines {
			break
		}
	}

	return lines, scanner.Err()
}
