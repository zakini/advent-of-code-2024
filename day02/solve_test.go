package day02_test

import (
	"testing"
	"zakini/advent-of-code-2024/day02"
	"zakini/advent-of-code-2024/utils"
)

func TestSolvePart1WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day02.SolvePart1, "example1.txt", 2)
}

func TestSolvePart2WithExample(t *testing.T) {
	t.SkipNow()
}
