package day02

import (
	"slices"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

const minDiff = 1
const maxDiff = 3

func SolvePart1(input string) int {
	reports := parseInput(input)

	safeCount := 0

	for _, report := range reports {
		if reportIsSafe(report) {
			safeCount++
		}
	}

	return safeCount
}

func SolvePart2(input string) int {
	reports := parseInput(input)

	safeCount := 0

	for _, report := range reports {
		if reportIsSafe(report) {
			safeCount++
		} else {
			// Try all possible reports that have 1 level removed
			for i := range report {
				copiedReport := slices.Clone(report)
				dampenedReport := append(copiedReport[:i], copiedReport[i+1:]...)
				if reportIsSafe(dampenedReport) {
					safeCount++
					break
				}
			}
		}
	}

	return safeCount
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")

	reports := make([][]int, len(lines))

	for i, line := range lines {
		elements := strings.Fields(line)

		for _, el := range elements {
			num := utils.ParseIntAndAssert(el)
			reports[i] = append(reports[i], num)
		}
	}

	return reports
}

func reportIsSafe(report []int) bool {
	direction := 0

	for i, level := range report {
		if i == 0 {
			continue
		}

		previousLevel := report[i-1]
		diff := level - previousLevel
		absDiff := utils.Abs(diff)

		if absDiff < minDiff || maxDiff < absDiff {
			// Out of range, report is unsafe
			return false
		}

		var diffDirection int
		if diff < 0 {
			diffDirection = -1
		} else if diff > 0 {
			diffDirection = 1
		} else {
			diffDirection = 0
		}

		if direction == 0 {
			direction = diffDirection
		} else if direction != diffDirection {
			// Direction has changed, report is unsafe
			return false
		}
	}

	return true
}
