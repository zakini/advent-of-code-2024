package day08

import (
	"fmt"
	"slices"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

type Vector2 struct {
	x int
	y int
}

type Antenna struct {
	position  Vector2
	frequency string
}

func (a Vector2) add(b Vector2) Vector2 {
	return Vector2{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

func (a Vector2) subtract(b Vector2) Vector2 {
	return Vector2{
		x: a.x - b.x,
		y: a.y - b.y,
	}
}

func SolvePart1(input string, debug bool) int {
	worldBounds, antennae := parseInput(input)

	antiNodes := make([]Vector2, 0)
	resolvedFrequencies := make([]string, 0)
	for _, antenna := range antennae {
		if slices.Contains(resolvedFrequencies, antenna.frequency) {
			// Already checked antennae with this frequency, skip this antenna
			continue
		}

		targetAntennae := utils.Filter(antennae, func(_ int, a Antenna) bool {
			return a.frequency == antenna.frequency
		})

		for i, a := range targetAntennae {
			for _, b := range targetAntennae[i+1:] {
				distance := b.position.subtract(a.position)
				antiNodes = append(
					antiNodes,
					a.position.subtract(distance),
					b.position.add(distance),
				)
			}
		}

		resolvedFrequencies = append(resolvedFrequencies, antenna.frequency)
	}

	// Remove anti-nodes that are out of bounds
	antiNodes = utils.Filter(antiNodes, func(_ int, antiNode Vector2) bool {
		return 0 <= antiNode.x && antiNode.x < worldBounds.x &&
			0 <= antiNode.y && antiNode.y < worldBounds.y
	})
	// Remove duplicate anti-nodes
	antiNodes = utils.Filter(antiNodes, func(i int, antiNode Vector2) bool {
		return i == slices.Index(antiNodes, antiNode)
	})

	if debug {
		printWorld(worldBounds, antennae, antiNodes)
	}

	return len(antiNodes)
}

func parseInput(input string) (Vector2, []Antenna) {
	lines := strings.Split(input, "\n")

	worldBounds := Vector2{len(lines[0]), len(lines)}
	antennae := make([]Antenna, 0)

	for y, line := range lines {
		chars := utils.SplitIntoChars(line)
		for x, char := range chars {
			if char == "." {
				continue
			}

			antennae = append(antennae, Antenna{Vector2{x, y}, char})
		}
	}

	return worldBounds, antennae
}

func vector2Compare(a Vector2, b Vector2) int {
	if a.y < b.y {
		return -2
	} else if a.y > b.y {
		return +2
	} else {
		if a.x < b.x {
			return -1
		} else if a.x > b.x {
			return +1
		} else {
			return 0
		}
	}
}

func antennaPositionCompare(a Antenna, b Antenna) int {
	return vector2Compare(a.position, b.position)
}

func printWorld(worldBounds Vector2, antennae []Antenna, antiNodes []Vector2) {
	slices.SortFunc(antennae, antennaPositionCompare)
	slices.SortFunc(antiNodes, vector2Compare)

	for y := range worldBounds.y {
		for x := range worldBounds.x {
			i, antennaFound := slices.BinarySearchFunc(
				antennae,
				Antenna{Vector2{x, y}, ""}, // This search doesn't care about the frequency
				antennaPositionCompare,
			)
			_, antiNodeFound := slices.BinarySearchFunc(
				antiNodes,
				Vector2{x, y},
				vector2Compare,
			)

			if antennaFound {
				fmt.Print(antennae[i].frequency)
			} else if antiNodeFound {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
