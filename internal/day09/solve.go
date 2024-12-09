package day09

import (
	"fmt"
	"slices"
	"strconv"
	"zakini/advent-of-code-2024/internal/utils"
)

const emptyBlock = -1

func SolvePart1(input string, debug bool) int {
	blockMap, _ := parseInput(input)
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

	return calculateChecksum(blockMap)
}

func SolvePart2(input string, debug bool) int {
	blockMap, blockLengthMap := parseInput(input)
	if debug {
		printBlockMap(blockMap)
	}

	for i := len(blockLengthMap) - 1; i >= 0; i-- {
		targetRunIndex := slices.Index(blockMap, i)
		targetRunLength := blockLengthMap[i]

		emptyRunIndex := -1
		maybeEmptyRunIndex := -1
		emptyRunLength := 0
		for j, block := range blockMap[:targetRunIndex] {
			if block != emptyBlock {
				emptyRunLength = 0
				maybeEmptyRunIndex = j + 1
				continue
			}

			emptyRunLength++

			if emptyRunLength == targetRunLength {
				emptyRunIndex = maybeEmptyRunIndex
				break
			}
		}

		if emptyRunIndex == -1 {
			continue
		}

		for i := 0; i < emptyRunLength; i++ {
			blockMap[emptyRunIndex+i], blockMap[targetRunIndex+i] = blockMap[targetRunIndex+i], blockMap[emptyRunIndex+i]
		}
	}

	if debug {
		printBlockMap(blockMap)
	}

	return calculateChecksum(blockMap)
}

func parseInput(input string) ([]int, []int) {
	chars := utils.SplitIntoChars(input)

	blockMap := make([]int, 0)
	blockLengthMap := make([]int, 0)
	blockId := 0
	for i, char := range chars {
		length := utils.ParseIntAndAssert(char)

		if i%2 == 0 {
			blockMap = append(blockMap, slices.Repeat([]int{blockId}, length)...)
			blockLengthMap = append(blockLengthMap, length)
			blockId++
		} else {
			blockMap = append(blockMap, slices.Repeat([]int{emptyBlock}, length)...)
		}
	}

	return blockMap, blockLengthMap
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

func calculateChecksum(blockMap []int) int {
	checksum := 0
	for i, blockId := range blockMap {
		if blockId == emptyBlock {
			continue
		}

		checksum += i * blockId
	}

	return checksum
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
