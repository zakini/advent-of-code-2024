package day10_test

import (
	"testing"
	"zakini/advent-of-code-2024/internal/day10"
	"zakini/advent-of-code-2024/internal/utils"
)

func TestSolvePart1WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day10.SolvePart1, "example1.txt", 36)
}

func TestSolvePart2WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day10.SolvePart2, "example1.txt", 81)
}
