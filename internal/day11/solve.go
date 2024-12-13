package day11

import (
	"strconv"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

const simulationStepCount = 25

func SolvePart1(input string, debug bool) int {
	stones := parseInput(input)

	for range simulationStepCount {
		stones = simulationStep(stones)
	}

	return len(stones)
}

func parseInput(input string) []int {
	stringStones := strings.Fields(input)

	stones := make([]int, len(stringStones))
	for i, char := range stringStones {
		stones[i] = utils.ParseIntAndAssert(char)
	}

	return stones
}

func simulationStep(stones []int) []int {
	newStones := make([]int, 0, len(stones))
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if stringStone := strconv.Itoa(stone); len(stringStone)%2 == 0 {
			chars := utils.SplitIntoChars(stringStone)
			newStones = append(newStones, utils.ParseIntAndAssert(strings.Join(chars[:len(chars)/2], "")))
			newStones = append(newStones, utils.ParseIntAndAssert(strings.Join(chars[len(chars)/2:], "")))
		} else {
			newStones = append(newStones, stone*2024)
		}
	}

	return newStones
}
