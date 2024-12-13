package day11_test

import (
	"testing"
	"zakini/advent-of-code-2024/internal/day11"
	"zakini/advent-of-code-2024/internal/utils"
)

func TestSolvePart1WithExample(t *testing.T) {
	utils.TestAgainstExample(t, day11.SolvePart1, "example1.txt", 55312)
}

// NOTE there's no example result given for part 2 of this puzzle, so we can't add a test for part 2
