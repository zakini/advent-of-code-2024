package day12_test

import (
	"testing"
	"zakini/advent-of-code-2024/internal/day12"
	"zakini/advent-of-code-2024/internal/utils"
)

func TestSolvePart1WithExample1(t *testing.T) {
	utils.TestAgainstExample(t, day12.SolvePart1, "example1.txt", 140)
}

func TestSolvePart1WithExample2(t *testing.T) {
	utils.TestAgainstExample(t, day12.SolvePart1, "example2.txt", 772)
}

func TestSolvePart1WithExample3(t *testing.T) {
	utils.TestAgainstExample(t, day12.SolvePart1, "example3.txt", 1930)
}
