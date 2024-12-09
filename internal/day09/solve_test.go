package day09_test

import (
	"testing"
	"zakini/advent-of-code-2024/internal/day09"
	"zakini/advent-of-code-2024/internal/utils"
)

func TestSolvePart1WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day09.SolvePart1, "example1.txt", 1928)
}
