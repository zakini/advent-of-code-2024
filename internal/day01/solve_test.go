package day01_test

import (
	"testing"
	"zakini/advent-of-code-2024/internal/day01"
	"zakini/advent-of-code-2024/internal/utils"
)

func TestSolvePart1WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day01.SolvePart1, "example1.txt", 11)
}

func TestSolvePart2WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day01.SolvePart2, "example1.txt", 31)
}
