package day01

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolvePart1(input string) int {
	lines := strings.Split(input, "\n")

	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))
	var parseError error

	for i, line := range lines {
		nums := strings.Fields(line)

		list1[i], parseError = strconv.Atoi(nums[0])
		assertNoError(parseError, fmt.Sprintf("Failed to parse number %s", nums[0]))
		list2[i], parseError = strconv.Atoi(nums[1])
		assertNoError(parseError, fmt.Sprintf("Failed to parse number %s", nums[1]))
	}

	slices.Sort(list1)
	slices.Sort(list2)

	var result int

	for i := range lines {
		result += abs(list1[i] - list2[i])
	}

	return result
}

func assert(condition bool, message string) {
	if condition {
		panic(message)
	}
}

func assertNoError(err error, message string) {
	assert(err != nil, message)
}

// bruh how is this not in the standard lib
func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}
