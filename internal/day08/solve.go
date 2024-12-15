package day08

import (
	"fmt"
	"slices"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

type Antenna struct {
	position  utils.Vector2
	frequency string
}

func SolvePart1(input string, debug bool) int {
	worldBounds, antennae := parseInput(input)

	antiNodes := make([]utils.Vector2, 0)
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
				distance := b.position.Subtract(a.position)

				antiNodeA := a.position.Subtract(distance)
				antiNodeB := b.position.Add(distance)

				if pointInBounds(antiNodeA, worldBounds) {
					antiNodes = append(antiNodes, antiNodeA)
				}
				if pointInBounds(antiNodeB, worldBounds) {
					antiNodes = append(antiNodes, antiNodeB)
				}
			}
		}

		resolvedFrequencies = append(resolvedFrequencies, antenna.frequency)
	}

	// Remove duplicate anti-nodes
	antiNodes = utils.Filter(antiNodes, func(i int, antiNode utils.Vector2) bool {
		return i == slices.Index(antiNodes, antiNode)
	})

	if debug {
		printWorld(worldBounds, antennae, antiNodes)
	}

	return len(antiNodes)
}

func SolvePart2(input string, debug bool) int {
	worldBounds, antennae := parseInput(input)

	antiNodes := make([]utils.Vector2, 0)
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
				distance := b.position.Subtract(a.position)

				antiNode := a.position.Clone()
				for pointInBounds(antiNode, worldBounds) {
					antiNodes = append(antiNodes, antiNode)
					antiNode = antiNode.Add(distance)
				}

				antiNode = a.position.Clone().Subtract(distance)
				for pointInBounds(antiNode, worldBounds) {
					antiNodes = append(antiNodes, antiNode)
					antiNode = antiNode.Subtract(distance)
				}
			}
		}

		resolvedFrequencies = append(resolvedFrequencies, antenna.frequency)
	}

	// Remove duplicate anti-nodes
	antiNodes = utils.Filter(antiNodes, func(i int, antiNode utils.Vector2) bool {
		return i == slices.Index(antiNodes, antiNode)
	})

	if debug {
		printWorld(worldBounds, antennae, antiNodes)
	}

	return len(antiNodes)
}

func parseInput(input string) (utils.Vector2, []Antenna) {
	lines := strings.Split(input, "\n")

	worldBounds := utils.Vector2{X: len(lines[0]), Y: len(lines)}
	antennae := make([]Antenna, 0)

	for y, line := range lines {
		chars := utils.SplitIntoChars(line)
		for x, char := range chars {
			if char == "." {
				continue
			}

			antennae = append(antennae, Antenna{utils.Vector2{X: x, Y: y}, char})
		}
	}

	return worldBounds, antennae
}

func vector2Compare(a utils.Vector2, b utils.Vector2) int {
	if a.Y < b.Y {
		return -2
	} else if a.Y > b.Y {
		return +2
	} else {
		if a.X < b.X {
			return -1
		} else if a.X > b.X {
			return +1
		} else {
			return 0
		}
	}
}

func antennaPositionCompare(a Antenna, b Antenna) int {
	return vector2Compare(a.position, b.position)
}

func pointInBounds(point utils.Vector2, bounds utils.Vector2) bool {
	return 0 <= point.X && point.X < bounds.X &&
		0 <= point.Y && point.Y < bounds.Y
}

func printWorld(worldBounds utils.Vector2, antennae []Antenna, antiNodes []utils.Vector2) {
	slices.SortFunc(antennae, antennaPositionCompare)
	slices.SortFunc(antiNodes, vector2Compare)

	for y := range worldBounds.Y {
		for x := range worldBounds.X {
			i, antennaFound := slices.BinarySearchFunc(
				antennae,
				Antenna{position: utils.Vector2{X: x, Y: y}, frequency: ""}, // This search doesn't care about the frequency
				antennaPositionCompare,
			)
			_, antiNodeFound := slices.BinarySearchFunc(
				antiNodes,
				utils.Vector2{X: x, Y: y},
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
