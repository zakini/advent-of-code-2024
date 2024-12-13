package main

import (
	"flag"
	"fmt"
	"zakini/advent-of-code-2024/internal/day01"
	"zakini/advent-of-code-2024/internal/day02"
	"zakini/advent-of-code-2024/internal/day03"
	"zakini/advent-of-code-2024/internal/day04"
	"zakini/advent-of-code-2024/internal/day05"
	"zakini/advent-of-code-2024/internal/day06"
	"zakini/advent-of-code-2024/internal/day07"
	"zakini/advent-of-code-2024/internal/day08"
	"zakini/advent-of-code-2024/internal/day09"
	"zakini/advent-of-code-2024/internal/day10"
	"zakini/advent-of-code-2024/internal/day11"
	"zakini/advent-of-code-2024/internal/utils"
)

var solverMap = map[string]map[string]utils.Solver{
	"day01": {"part1": day01.SolvePart1, "part2": day01.SolvePart2},
	"day02": {"part1": day02.SolvePart1, "part2": day02.SolvePart2},
	"day03": {"part1": day03.SolvePart1, "part2": day03.SolvePart2},
	"day04": {"part1": day04.SolvePart1, "part2": day04.SolvePart2},
	"day05": {"part1": day05.SolvePart1, "part2": day05.SolvePart2},
	"day06": {"part1": day06.SolvePart1, "part2": day06.SolvePart2},
	"day07": {"part1": day07.SolvePart1, "part2": day07.SolvePart2},
	"day08": {"part1": day08.SolvePart1, "part2": day08.SolvePart2},
	"day09": {"part1": day09.SolvePart1, "part2": day09.SolvePart2},
	"day10": {"part1": day10.SolvePart1, "part2": day10.SolvePart2},
	"day11": {"part1": day11.SolvePart1},
}

func main() {
	solver, filePath, debug := parseArgs()
	fileContents := utils.LoadInputFile(filePath)
	result := solver(fileContents, debug)

	fmt.Printf("Result: %d\n", result)
}

func parseArgs() (utils.Solver, string, bool) {
	debugFlag := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()
	args := flag.Args()

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

	return solver, filePath, *debugFlag
}
