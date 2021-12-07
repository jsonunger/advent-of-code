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
