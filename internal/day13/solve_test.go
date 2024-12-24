package day13_test

import (
	"testing"
	"zakini/advent-of-code-2024/internal/day13"
	"zakini/advent-of-code-2024/internal/utils"
)

func TestSolvePart1WithExample1(t *testing.T) {
	utils.TestAgainstExample(t, day13.SolvePart1, "example1.txt", 480)
}
