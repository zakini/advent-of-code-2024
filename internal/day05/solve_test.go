package day05_test

import (
	"testing"
	"zakini/advent-of-code-2024/internal/day05"
	"zakini/advent-of-code-2024/internal/utils"
)

func TestSolvePart1WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day05.SolvePart1, "example1.txt", 143)
}

func TestSolvePart2WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day05.SolvePart2, "example1.txt", 123)
}
