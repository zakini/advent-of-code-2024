package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Solver func(string, bool) int

func LoadInputFile(filePath string) string {
	fileData, err := os.ReadFile(filePath)
	AssertNoErrorWithMessage(err, fmt.Sprintf("Could not read input file: %s", filePath))

	return strings.TrimSpace(string(fileData))
}

func Assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func AssertNoError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func AssertNoErrorWithMessage(err error, message string) {
	Assert(err == nil, message)
}

func ParseIntAndAssert(str string) int {
	num, err := strconv.Atoi(str)
	AssertNoErrorWithMessage(err, fmt.Sprintf("Failed to parse number %s", str))
	return num
}

// bruh how is this not in the standard lib
func Abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func SplitIntoChars(value string) []string {
	chars := make([]string, 0, len(value))
	for _, r := range value {
		chars = append(chars, string(r))
	}
	return chars
}

func Filter[T any](arr []T, predicate func(int, T) bool) []T {
	out := make([]T, 0, len(arr))

	for i, el := range arr {
		if predicate(i, el) {
			out = append(out, el)
		}
	}

	return out
}

func FindSurroundingPoints[T any](world [][]T, centre Vector2) []Vector2 {
	return FindSurroundingPointsFunc(world, centre, func(_ T) bool { return true })
}

func FindSurroundingPointsFunc[T any](world [][]T, centre Vector2, predicate func(T) bool) []Vector2 {
	possiblePoints := [...]Vector2{
		{centre.X - 1, centre.Y},
		{centre.X + 1, centre.Y},
		{centre.X, centre.Y - 1},
		{centre.X, centre.Y + 1},
	}

	points := make([]Vector2, 0, len(possiblePoints))
	for _, point := range possiblePoints {
		if 0 <= point.Y && point.Y < len(world) && 0 <= point.X && point.X < len(world[point.Y]) && predicate(world[point.Y][point.X]) {
			points = append(points, point)
		}
	}

	return points
}
