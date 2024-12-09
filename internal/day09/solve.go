package day09

import (
	"fmt"
	"slices"
	"strconv"
	"zakini/advent-of-code-2024/internal/utils"
)

const emptyBlock = -1

func SolvePart1(input string, debug bool) int {
	blockMap := parseInput(input)
	if debug {
		printBlockMap(blockMap)
	}

	emptyIndex := slices.Index(blockMap, emptyBlock)
	filledIndex := lastIndexFunc(blockMap, func(b int) bool {
		return b != emptyBlock
	})
	for emptyIndex < filledIndex {
		blockMap[emptyIndex], blockMap[filledIndex] = blockMap[filledIndex], blockMap[emptyIndex]

		for blockMap[emptyIndex] != emptyBlock {
			emptyIndex++
		}
		for blockMap[filledIndex] == emptyBlock {
			filledIndex--
		}
	}

	if debug {
		printBlockMap(blockMap)
	}

	checksum := 0
	for i, blockId := range blockMap {
		if blockId == emptyBlock {
			break
		}

		checksum += i * blockId
	}

	return checksum
}

func parseInput(input string) []int {
	chars := utils.SplitIntoChars(input)

	blockMap := make([]int, 0)
	blockId := 0
	for i, char := range chars {
		value := utils.ParseIntAndAssert(char)

		if i%2 == 0 {
			blockMap = append(blockMap, slices.Repeat([]int{blockId}, value)...)
			blockId++
		} else {
			blockMap = append(blockMap, slices.Repeat([]int{emptyBlock}, value)...)
		}
	}

	return blockMap
}

func lastIndexFunc(s []int, f func(int) bool) int {
	index := -1

	for i, e := range s {
		if f(e) {
			index = i
		}
	}

	return index
}

func printBlockMap(blockMap []int) {
	for _, block := range blockMap {
		if block == emptyBlock {
			fmt.Print(".")
		} else {
			fmt.Print(strconv.Itoa(block))
		}
	}

	fmt.Println()
}
