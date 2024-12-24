package day13

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

type ClawMachine struct {
	buttonA       utils.Vector2
	buttonB       utils.Vector2
	prizePosition utils.Vector2
}

type Node struct {
	position utils.Vector2
	cost     int
}

const buttonACost = 3
const buttonBCost = 1

func SolvePart1(input string, debug bool) int {
	machines := parseInput(input)

	tokenCost := 0
	for _, machine := range machines {
		if tokens, err := solveMachine(machine); err == nil {
			tokenCost += tokens
		}
	}

	return tokenCost
}

func parseInput(input string) []ClawMachine {
	buttonAMatcher := regexp.MustCompile(`^Button A: X\+(\d+), Y\+(\d+)$`)
	buttonBMatcher := regexp.MustCompile(`^Button B: X\+(\d+), Y\+(\d+)$`)
	prizeMatcher := regexp.MustCompile(`^Prize: X=(\d+), Y=(\d+)$`)

	var buttonA utils.Vector2
	var buttonB utils.Vector2
	var prizePosition utils.Vector2

	lines := strings.Split(input, "\n")
	// Re-add a blank line so that we don't have to handle the last claw machine
	// as a special case
	lines = append(lines, "")
	machines := make([]ClawMachine, 0, (len(lines)+1)/4)
	for _, line := range lines {
		if matches := buttonAMatcher.FindStringSubmatch(line); matches != nil {
			buttonA.X = utils.ParseIntAndAssert(matches[1])
			buttonA.Y = utils.ParseIntAndAssert(matches[2])
		} else if matches := buttonBMatcher.FindStringSubmatch(line); matches != nil {
			buttonB.X = utils.ParseIntAndAssert(matches[1])
			buttonB.Y = utils.ParseIntAndAssert(matches[2])
		} else if matches := prizeMatcher.FindStringSubmatch(line); matches != nil {
			prizePosition.X = utils.ParseIntAndAssert(matches[1])
			prizePosition.Y = utils.ParseIntAndAssert(matches[2])
		} else if line == "" {
			utils.Assert(
				buttonA != (utils.Vector2{}) &&
					buttonB != (utils.Vector2{}) &&
					prizePosition != (utils.Vector2{}),
				"Claw machine is missing some elements",
			)

			machines = append(
				machines,
				ClawMachine{buttonA, buttonB, prizePosition},
			)
		} else {
			panic(fmt.Sprintf("Failed to parse line: %v", line))
		}
	}

	return machines
}

func solveMachine(machine ClawMachine) (int, error) {
	frontier := []Node{
		{position: utils.Vector2{X: 0, Y: 0}, cost: 0},
	}

	bestTokenCost := math.MaxInt
	checks := 0
	const maxChecks = 1_000_000
	for len(frontier) > 0 {
		var current Node
		current, frontier = frontier[0], frontier[1:]

		checks++
		utils.Assert(checks <= maxChecks, "Reached max checks")

		if current.position.X > machine.prizePosition.X ||
			current.position.Y > machine.prizePosition.Y {
			// Buttons always increase X and Y, so if we ever go past either
			// coordinate of the prize position, we're never coming back, so we
			// can ignore this node
			continue
		}

		if current.position.X == machine.prizePosition.X &&
			current.position.Y == machine.prizePosition.Y &&
			current.cost < bestTokenCost {
			// Found a better route to the prize
			bestTokenCost = current.cost
		} else {
			buttonANeighbour := Node{
				position: current.position.Add(machine.buttonA),
				cost:     current.cost + buttonACost,
			}
			if !slices.Contains(frontier, buttonANeighbour) {
				frontier = append(frontier, buttonANeighbour)
			}

			buttonBNeighbour := Node{
				position: current.position.Add(machine.buttonB),
				cost:     current.cost + buttonBCost,
			}
			if !slices.Contains(frontier, buttonBNeighbour) {
				frontier = append(frontier, buttonBNeighbour)
			}
		}
	}

	if bestTokenCost == math.MaxInt {
		return -1, errors.New("machine is unsolvable")
	}

	return bestTokenCost, nil
}
