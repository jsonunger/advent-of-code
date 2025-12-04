package utils

import (
	"strconv"

	"github.com/pkg/errors"
)

func PanicWithMsg(err error, message string) {
	wrapped := errors.Wrap(err, message)
	if wrapped != nil {
		panic(wrapped)
	}
}

func ConvertStringToInt(str string) int {
	val, err := strconv.Atoi(str)
	PanicWithMsg(err, "converting "+str)
	return val
}

func ChunkString(s string, chunkSize int) []string {
	if chunkSize <= 0 {
		return []string{s}
	}

	runes := []rune(s) // Convert string to a slice of runes
	var chunks []string

	for i := 0; i < len(runes); i += chunkSize {
		end := i + chunkSize
		if end > len(runes) {
			end = len(runes)
		}
		chunks = append(chunks, string(runes[i:end]))
	}
	return chunks
}
