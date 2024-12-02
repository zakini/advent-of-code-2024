package utils

import (
	"fmt"
	"strconv"
)

func Assert(condition bool, message string) {
	if condition {
		panic(message)
	}
}

func AssertNoError(err error, message string) {
	Assert(err != nil, message)
}

func ParseIntAndAssert(str string) int {
	num, err := strconv.Atoi(str)
	AssertNoError(err, fmt.Sprintf("Failed to parse number %s", str))
	return num
}

// bruh how is this not in the standard lib
func Abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}
