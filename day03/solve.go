package day03

import (
	"fmt"
	"regexp"
	"zakini/advent-of-code-2024/utils"
)

func SolvePart1(input string) int {
	parser := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := parser.FindAllStringSubmatch(input, -1)

	utils.Assert(matches != nil, "Failed to parse input")

	result := 0
	for _, match := range matches {
		utils.Assert(len(match) == 3, fmt.Sprintf("Unexpected match: %v", match))
		numA := utils.ParseIntAndAssert(match[1])
		numB := utils.ParseIntAndAssert(match[2])
		result += numA * numB
	}

	return result
}

func SolvePart2(input string) int {
	parser := regexp.MustCompile(`(mul)\((\d{1,3}),(\d{1,3})\)|(do)\(\)|(don't)\(\)`)
	matches := parser.FindAllStringSubmatch(input, -1)

	utils.Assert(matches != nil, "Failed to parse input")

	result := 0
	mulEnabled := true
	for _, match := range matches {
		utils.Assert(len(match) >= 2, fmt.Sprintf("Unexpected match: %v", match))
		maybeMulName := match[1]
		maybeDoName := match[4]
		maybeDontName := match[5]

		switch {
		case maybeMulName == "mul":
			if mulEnabled {
				utils.Assert(len(match) >= 4, fmt.Sprintf("Unexpected match: %v", match))
				numA := utils.ParseIntAndAssert(match[2])
				numB := utils.ParseIntAndAssert(match[3])
				result += numA * numB
			}
		case maybeDoName == "do":
			mulEnabled = true
		case maybeDontName == "don't":
			mulEnabled = false
		default:
			panic(fmt.Sprintf("No functions matched: %v", match))
		}
	}

	return result
}
