package utils

import (
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func ReadFile(fileName string) []string {
	data, err := os.ReadFile(fileName)
	PanicWithMsg(err, "getting working directory")

	lines := strings.Split(string(data), "\n")
	trimmedLines := make([]string, len(lines))
	for i, line := range lines {
		trimmedLines[i] = strings.TrimSpace(line)
	}
	return trimmedLines
}

func PanicWithMsg(err error, message string) {
	wrapped := errors.Wrap(err, message)
	if wrapped != nil {
		panic(wrapped)
	}
}

func ConvertStringsToInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, str := range strs {
		val, err := strconv.Atoi(str)
		PanicWithMsg(err, "converting "+str)
		ints[i] = val
	}
	return ints
}
