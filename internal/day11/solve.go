package day11

import (
	"fmt"
	"strconv"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

func SolvePart1(input string, debug bool) int {
	stoneMap := parseInput(input)
	return simulate(stoneMap, 25)
}

func SolvePart2(input string, debug bool) int {
	stoneMap := parseInput(input)
	return simulate(stoneMap, 75)
}

func parseInput(input string) map[int]int {
	stringStones := strings.Fields(input)

	stoneMap := make(map[int]int)
	for _, char := range stringStones {
		stone := utils.ParseIntAndAssert(char)
		stoneMap[stone]++
	}

	return stoneMap
}

func simulate(stoneMap map[int]int, steps int) int {
	for range steps {
		stoneMap = simulationStep(stoneMap)
		fmt.Print(".")
	}

	fmt.Println()

	result := 0
	for _, count := range stoneMap {
		result += count
	}

	return result
}

func simulationStep(stoneMap map[int]int) map[int]int {
	newStoneMap := make(map[int]int)

	for stone, count := range stoneMap {
		newStones := make([]int, 0, 2)

		if stone == 0 {
			newStones = append(newStones, 1)
		} else if stringStone := strconv.Itoa(stone); len(stringStone)%2 == 0 {
			chars := utils.SplitIntoChars(stringStone)
			newStones = append(newStones, utils.ParseIntAndAssert(strings.Join(chars[:len(chars)/2], "")))
			newStones = append(newStones, utils.ParseIntAndAssert(strings.Join(chars[len(chars)/2:], "")))
		} else {
			newStones = append(newStones, stone*2024)
		}

		for _, stone := range newStones {
			newStoneMap[stone] += count
		}
	}

	return newStoneMap
}
