package day06_test

import (
	"testing"
	"zakini/advent-of-code-2024/internal/day06"
	"zakini/advent-of-code-2024/internal/utils"
)

func TestSolvePart1WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day06.SolvePart1, "example1.txt", 41)
}
