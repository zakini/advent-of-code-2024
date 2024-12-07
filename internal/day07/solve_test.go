package day07_test

import (
	"testing"
	"zakini/advent-of-code-2024/internal/day07"
	"zakini/advent-of-code-2024/internal/utils"
)

func TestSolvePart1WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day07.SolvePart1, "example1.txt", 3749)
}

func TestSolvePart2WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day07.SolvePart2, "example1.txt", 11387)
}
