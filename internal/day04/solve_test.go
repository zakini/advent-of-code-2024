package day04_test

import (
	"testing"
	"zakini/advent-of-code-2024/internal/day04"
	"zakini/advent-of-code-2024/internal/utils"
)

func TestSolvePart1WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day04.SolvePart1, "example1.txt", 18)
}
