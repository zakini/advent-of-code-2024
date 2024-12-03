package day03_test

import (
	"testing"
	"zakini/advent-of-code-2024/internal/day03"
	"zakini/advent-of-code-2024/internal/utils"
)

func TestSolvePart1WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day03.SolvePart1, "example1.txt", 161)
}

func TestSolvePart2WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day03.SolvePart2, "example2.txt", 48)
}
