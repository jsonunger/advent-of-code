package utils

import (
	"os"
	"strings"
)

func ReadFile(fileName string) string {
	data, err := os.ReadFile(fileName)
	PanicWithMsg(err, "getting working directory")

	return strings.TrimSpace(string(data))
}

func ReadLines(fileName string) []string {
	lines := strings.Split(ReadFile(fileName), "\n")
	trimmedLines := make([]string, len(lines))
	for i, line := range lines {
		trimmedLines[i] = strings.TrimSpace(line)
	}
	return trimmedLines
}
