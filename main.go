package main

import (
	"fmt"
	"os"
	"zakini/advent-of-code-2024/day01"
	"zakini/advent-of-code-2024/day02"
	"zakini/advent-of-code-2024/day03"
	"zakini/advent-of-code-2024/utils"
)

var solverMap = map[string]map[string]utils.Solver{
	"day01": {"part1": day01.SolvePart1, "part2": day01.SolvePart2},
	"day02": {"part1": day02.SolvePart1},
	"day03": {"part1": day03.SolvePart1, "part2": day03.SolvePart2},
}

func main() {
	solver, filePath := parseArgs()
	fileContents := utils.LoadInputFile(filePath)
	result := solver(fileContents)

	fmt.Printf("Result: %d\n", result)
}

func parseArgs() (utils.Solver, string) {
	args := os.Args[1:]

	// TODO show usage message
	utils.Assert(len(args) >= 3, "Missing required arguments")

	day := args[0]
	part := args[1]
	filePath := args[2]

	partsMap, dayValid := solverMap[day]
	// TODO show valid options for day
	utils.Assert(dayValid, "Invalid day provided")

	solver, partValid := partsMap[part]
	// TODO show valid options for part
	utils.Assert(partValid, "Invalid part provided")

	return solver, filePath
}
