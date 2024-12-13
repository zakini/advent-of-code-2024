package day10

import (
	"fmt"
	"slices"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

type Point struct {
	x int
	y int
}

const maxPathValue = 9

func SolvePart1(input string, debug bool) int {
	world, trailStartingPoints := parseInput(input)

	result := 0
	for _, start := range trailStartingPoints {
		paths := findPaths(world, []Point{start})

		uniqueEndPoints := make([]Point, 0, len(paths))
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
		paths := findPaths(world, []Point{start})

		result += len(paths)
	}

	return result
}

func parseInput(input string) ([][]int, []Point) {
	lines := strings.Split(input, "\n")

	world := make([][]int, len(lines))
	trailStartingPoints := make([]Point, 0)
	for y, line := range lines {
		chars := utils.SplitIntoChars(line)
		world[y] = make([]int, len(chars))

		for x, char := range chars {
			world[y][x] = utils.ParseIntAndAssert(char)
			if world[y][x] == 0 {
				trailStartingPoints = append(trailStartingPoints, Point{x, y})
			}
		}
	}

	return world, trailStartingPoints
}

func findPaths(world [][]int, pathSoFar []Point) [][]Point {
	utils.Assert(len(pathSoFar) <= 10, fmt.Sprintf("Path is too long: %v", pathSoFar))

	pathHead := pathSoFar[len(pathSoFar)-1]

	paths := make([][]Point, 0)
	for _, point := range surroundingPoints(world, pathHead) {
		if world[point.y][point.x] != world[pathHead.y][pathHead.x]+1 {
			continue
		}

		path := make([]Point, 0, len(pathSoFar)+1)
		path = append(path, pathSoFar...)
		path = append(path, point)

		if world[point.y][point.x] >= maxPathValue {
			paths = append(paths, path)
		} else {
			paths = append(paths, findPaths(world, path)...)
		}
	}

	return paths
}

func surroundingPoints(world [][]int, centre Point) []Point {
	possiblePoints := [...]Point{
		{centre.x - 1, centre.y},
		{centre.x + 1, centre.y},
		{centre.x, centre.y - 1},
		{centre.x, centre.y + 1},
	}

	points := make([]Point, 0, len(possiblePoints))
	for _, point := range possiblePoints {
		if 0 <= point.y && point.y < len(world) && 0 <= point.x && point.x < len(world[point.y]) {
			points = append(points, point)
		}
	}

	return points
}
