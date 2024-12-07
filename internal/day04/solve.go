package day04

import (
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

type point struct {
	x int
	y int
}

func SolvePart1(input string, debug bool) int {
	chars := parseInput(input)
	wordCoords := findWord(chars, "XMAS")
	return len(wordCoords)
}

func SolvePart2(input string, debug bool) int {
	chars := parseInput(input)
	mases := findWord(chars, "MAS")

	diagonalMases := make([][]point, 0, len(mases))
	for _, mas := range mases {
		if mas[0].x != mas[1].x && mas[0].y != mas[1].y {
			diagonalMases = append(diagonalMases, mas)
		}
	}

	xMasCount := 0

	for i, mas1 := range diagonalMases {
		for _, mas2 := range diagonalMases[i+1:] {
			mas1APoint := mas1[1]
			mas2APoint := mas2[1]
			if mas1APoint == mas2APoint {
				xMasCount++
			}
		}
	}

	return xMasCount
}

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\n")
	chars := make([][]string, len(lines))
	for i, line := range lines {
		chars[i] = utils.SplitIntoChars(line)
	}

	return chars
}

func findWord(chars [][]string, word string) [][]point {
	matches := make([][]point, 0)

	wordChars := utils.SplitIntoChars(word)
	var startPoints []point
	for y, line := range chars {
		for x, char := range line {
			if char == wordChars[0] {
				startPoints = append(startPoints, point{x, y})
			}
		}
	}

	for _, start := range startPoints {
		nextPoints := findInSurroundingChars(chars, start, wordChars[1])
		for _, next := range nextPoints {
			direction := point{x: next.x - start.x, y: next.y - start.y}
			currentLocation := next

			wordCoords := make([]point, len(wordChars))
			wordCoords[0] = start
			wordCoords[1] = next

			found := true
			for i, char := range wordChars[2:] {
				if findCharInDirection(chars, currentLocation, char, direction) {
					currentLocation = point{x: next.x + direction.x, y: next.y + direction.y}
					wordCoords[i+2] = currentLocation
				} else {
					found = false
					break
				}
			}

			if found {
				matches = append(matches, wordCoords)
			}
		}
	}

	return matches
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

func findCharInDirection(chars [][]string, from point, target string, direction point) bool {
	maybeCharLocation := point{x: from.x + direction.x, y: from.y + direction.y}

	if !inBounds(chars, maybeCharLocation.y) || !inBounds(chars[maybeCharLocation.y], maybeCharLocation.x) {
		return false
	}

	return chars[maybeCharLocation.y][maybeCharLocation.x] == target
}

func inBounds[S ~[]E, E any](s S, i int) bool {
	return 0 <= i && i < len(s)
}
