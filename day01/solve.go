package day01

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolvePart1(input string) int {
	list1, list2 := parseInput(input)

	slices.Sort(list1)
	slices.Sort(list2)

	var distance int

	for i := range list1 {
		distance += abs(list1[i] - list2[i])
	}

	return distance
}

func SolvePart2(input string) int {
	list1, list2 := parseInput(input)

	var similarity int

	for _, n := range list1 {
		count := len(filter(list2, func(m int) bool {
			return n == m
		}))

		similarity += n * count
	}

	return similarity
}

func parseInput(input string) ([]int, []int) {
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

	return list1, list2
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

func filter(arr []int, predicate func(int) bool) []int {
	out := make([]int, 0, len(arr))

	for _, el := range arr {
		if predicate(el) {
			out = append(out, el)
		}
	}

	return out
}
