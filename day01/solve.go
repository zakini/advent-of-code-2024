package day01

import (
	"slices"
	"strings"
	"zakini/advent-of-code-2024/utils"
)

func SolvePart1(input string) int {
	list1, list2 := parseInput(input)

	slices.Sort(list1)
	slices.Sort(list2)

	var distance int

	for i := range list1 {
		distance += utils.Abs(list1[i] - list2[i])
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

	for i, line := range lines {
		nums := strings.Fields(line)

		list1[i] = utils.ParseIntAndAssert(nums[0])
		list2[i] = utils.ParseIntAndAssert(nums[1])
	}

	return list1, list2
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
