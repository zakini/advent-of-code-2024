package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Solver func(string) int

func LoadInputFile(filePath string) string {
	fileData, err := os.ReadFile(filePath)
	AssertNoError(err, fmt.Sprintf("Could not read input file: %s", filePath))

	return strings.TrimSpace(string(fileData))
}

func Assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func AssertNoError(err error, message string) {
	Assert(err == nil, message)
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