package day02

import (
	"strings"
	"zakini/advent-of-code-2024/utils"
)

const minDiff = 1
const maxDiff = 3

func SolvePart1(input string) int {
	lines := strings.Split(input, "\n")

	reports := make([][]int, len(lines))

	for i, line := range lines {
		elements := strings.Fields(line)

		for _, el := range elements {
			num := utils.ParseIntAndAssert(el)
			reports[i] = append(reports[i], num)
		}
	}

	safeCount := 0

	for _, report := range reports {
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
				break
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
				break
			}

			if i == len(report)-1 {
				// Reached the last level with no issues, report is safe
				safeCount++
			}
		}
	}

	return safeCount
}
