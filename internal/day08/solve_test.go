package day08_test

import (
	"testing"
	"zakini/advent-of-code-2024/internal/day08"
	"zakini/advent-of-code-2024/internal/utils"
)

func TestSolvePart1WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day08.SolvePart1, "example1.txt", 14)
}
