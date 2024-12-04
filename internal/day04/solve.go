package day04

import (
	"strings"
)

type point struct {
	x int
	y int
}

func SolvePart1(input string) int {
	lines := strings.Split(input, "\n")
	chars := make([][]string, len(lines))
	for i, line := range lines {
		chars[i] = splitIntoChars(line)
	}

	matches := 0

	for y, line := range chars {
		for x, char := range line {
			if char == "X" {
				matches += findXmasFromPoint(chars, point{x, y})
			}
		}
	}

	return matches
}

func splitIntoChars(value string) []string {
	chars := make([]string, 0, len(value))
	for _, r := range value {
		chars = append(chars, string(r))
	}
	return chars
}

func findXmasFromPoint(chars [][]string, xLocation point) int {
	xmases := 0

	mLocations := findInSurroundingChars(chars, xLocation, "M")
	for _, mLocation := range mLocations {
		direction := point{x: mLocation.x - xLocation.x, y: mLocation.y - xLocation.y}
		maybeALocation := point{x: mLocation.x + direction.x, y: mLocation.y + direction.y}

		if !inBounds(chars, maybeALocation.y) || !inBounds(chars[maybeALocation.y], maybeALocation.x) {
			continue
		}

		if chars[maybeALocation.y][maybeALocation.x] == "A" {
			maybeSLocation := point{x: maybeALocation.x + direction.x, y: maybeALocation.y + direction.y}

			if !inBounds(chars, maybeSLocation.y) || !inBounds(chars[maybeSLocation.y], maybeSLocation.x) {
				continue
			}

			if chars[maybeSLocation.y][maybeSLocation.x] == "S" {
				xmases++
			}
		}
	}

	return xmases
}

func findInSurroundingChars(chars [][]string, center point, target string) []point {
	matches := make([]point, 0, 8) // there are 8 chars to check around a center point

	for yOffset := -1; yOffset <= 1; yOffset++ {
		y := center.y + yOffset
		if !inBounds(chars, y) {
			continue
		}

		line := chars[y]

		for xOffset := -1; xOffset <= 1; xOffset++ {
			x := center.x + xOffset
			if !inBounds(line, x) {
				continue
			}
			if line[x] == target {
				matches = append(matches, point{x, y})
			}
		}
	}

	return matches
}

func inBounds[S ~[]E, E any](s S, i int) bool {
	return 0 <= i && i < len(s)
}
