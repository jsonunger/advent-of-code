package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func ReadFile(fileName string) []string {
	dir, err := os.Getwd()
	PanicWithMsg(err, "getting working directory")
	year, _ := os.LookupEnv("YEAR")
	day, _ := os.LookupEnv("DAY")

	fullPath := fmt.Sprintf("%s/%s/day%s/%s", dir, year, day, fileName)
	data, err := ioutil.ReadFile(fullPath)
	PanicWithMsg(err, fmt.Sprintf("unable to read %s", fullPath))

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
