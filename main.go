package main

import (
	"fmt"
	"os"
	"strings"
	"zakini/advent-of-code-2024/day01"
	"zakini/advent-of-code-2024/day02"
	"zakini/advent-of-code-2024/day03"
)

type Solver func(string) int

var solverMap = map[string]map[string]Solver{
	"day01": {"part1": day01.SolvePart1, "part2": day01.SolvePart2},
	"day02": {"part1": day02.SolvePart1},
	"day03": {"part1": day03.SolvePart1},
}

func main() {
	solver, filePath := parseArgs()
	fileContents := loadInputFile(filePath)
	result := solver(fileContents)

	fmt.Printf("Result: %d\n", result)
}

func parseArgs() (Solver, string) {
	args := os.Args[1:]

	// TODO show usage message
	exitIf(len(args) < 3, "Missing required arguments")

	day := args[0]
	part := args[1]
	filePath := args[2]

	partsMap, dayValid := solverMap[day]
	// TODO show valid options for day
	exitUnless(dayValid, "Invalid day provided")

	solver, partValid := partsMap[part]
	// TODO show valid options for part
	exitUnless(partValid, "Invalid part provided")

	return solver, filePath
}

func loadInputFile(filePath string) string {
	fileData, err := os.ReadFile(filePath)
	exitIfErr(err, fmt.Sprintf("Could not read input file: %s", filePath))

	return strings.TrimSpace(string(fileData))
}

func exitIf(condition bool, message string) {
	if condition {
		fmt.Fprintln(os.Stderr, message)
		os.Exit(1)
	}
}

func exitUnless(condition bool, message string) {
	exitIf(!condition, message)
}

func exitIfErr(err error, message string) {
	exitIf(err != nil, message)
}
