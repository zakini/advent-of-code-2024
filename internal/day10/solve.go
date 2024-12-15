package day10

import (
	"fmt"
	"slices"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

const maxPathValue = 9

func SolvePart1(input string, debug bool) int {
	world, trailStartingPoints := parseInput(input)

	result := 0
	for _, start := range trailStartingPoints {
		paths := findPaths(world, []utils.Vector2{start})

		uniqueEndPoints := make([]utils.Vector2, 0, len(paths))
		for _, path := range paths {
			endPoint := path[len(path)-1]
			if !slices.Contains(uniqueEndPoints, endPoint) {
				uniqueEndPoints = append(uniqueEndPoints, endPoint)
			}
		}

		result += len(uniqueEndPoints)
	}

	return result
}

func SolvePart2(input string, debug bool) int {
	world, trailStartingPoints := parseInput(input)

	result := 0
	for _, start := range trailStartingPoints {
		paths := findPaths(world, []utils.Vector2{start})

		result += len(paths)
	}

	return result
}

func parseInput(input string) ([][]int, []utils.Vector2) {
	lines := strings.Split(input, "\n")

	world := make([][]int, len(lines))
	trailStartingPoints := make([]utils.Vector2, 0)
	for y, line := range lines {
		chars := utils.SplitIntoChars(line)
		world[y] = make([]int, len(chars))

		for x, char := range chars {
			world[y][x] = utils.ParseIntAndAssert(char)
			if world[y][x] == 0 {
				trailStartingPoints = append(trailStartingPoints, utils.Vector2{X: x, Y: y})
			}
		}
	}

	return world, trailStartingPoints
}

func findPaths(world [][]int, pathSoFar []utils.Vector2) [][]utils.Vector2 {
	utils.Assert(len(pathSoFar) <= 10, fmt.Sprintf("Path is too long: %v", pathSoFar))

	pathHead := pathSoFar[len(pathSoFar)-1]

	paths := make([][]utils.Vector2, 0)
	for _, point := range utils.FindSurroundingPoints(world, pathHead) {
		if world[point.Y][point.X] != world[pathHead.Y][pathHead.X]+1 {
			continue
		}

		path := make([]utils.Vector2, 0, len(pathSoFar)+1)
		path = append(path, pathSoFar...)
		path = append(path, point)

		if world[point.Y][point.X] >= maxPathValue {
			paths = append(paths, path)
		} else {
			paths = append(paths, findPaths(world, path)...)
		}
	}

	return paths
}
