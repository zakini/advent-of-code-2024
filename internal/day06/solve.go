package day06

import (
	"fmt"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

type Node struct {
	obstacle bool
	visited  bool
}

type Point struct {
	x int
	y int
}

var nullPoint = Point{-1, -1}

type Direction int

const (
	Unknown Direction = iota
	Up
	Right
	Down
	Left
)

type OrientatedPosition struct {
	position  Point
	direction Direction
}

var nullPosition = OrientatedPosition{Point{-1, -1}, Unknown}

func SolvePart1(input string, debug bool) int {
	world, guardPosition := parseInput(input)
	if debug {
		printWorld(&world, &guardPosition)
	}

	for pointIsInWorld(&world, guardPosition.position) {
		simulationStep(&world, &guardPosition)
		if debug {
			printWorld(&world, &guardPosition)
		}
	}

	visited := 0
	for _, row := range world {
		for _, node := range row {
			if node.visited {
				visited++
			}
		}
	}

	return visited
}

func parseInput(input string) ([][]Node, OrientatedPosition) {
	lines := strings.Split(input, "\n")

	world := make([][]Node, len(lines))
	guardPosition := nullPosition

	for y, line := range lines {
		chars := utils.SplitIntoChars(line)
		world[y] = make([]Node, len(chars))
		for x, char := range chars {
			world[y][x] = Node{
				obstacle: char == "#",
				visited:  false,
			}

			if char == "^" {
				utils.Assert(
					guardPosition == nullPosition,
					fmt.Sprintf("Found duplicate guard positions: %v | %v", guardPosition, OrientatedPosition{Point{x, y}, Up}),
				)
				guardPosition = OrientatedPosition{Point{x, y}, Up}
				world[y][x].visited = true
			}
		}
	}

	utils.Assert(guardPosition != nullPosition, "Guard position not found")

	return world, guardPosition
}

func simulationStep(world *[][]Node, guardPosition *OrientatedPosition) {
	targetPoint := nullPoint

	switch guardPosition.direction {
	case Up:
		targetPoint = Point{guardPosition.position.x, guardPosition.position.y - 1}
	case Right:
		targetPoint = Point{guardPosition.position.x + 1, guardPosition.position.y}
	case Down:
		targetPoint = Point{guardPosition.position.x, guardPosition.position.y + 1}
	case Left:
		targetPoint = Point{guardPosition.position.x - 1, guardPosition.position.y}
	default:
		panic(fmt.Sprintf("Invalid guard direction: %v", guardPosition.direction))
	}

	targetIsInWorld := pointIsInWorld(world, targetPoint)

	if !targetIsInWorld || !(*world)[targetPoint.y][targetPoint.x].obstacle {
		// Move 1 step forward
		guardPosition.position = targetPoint
		if targetIsInWorld {
			(*world)[targetPoint.y][targetPoint.x].visited = true
		}
	} else {
		// Rotate right, don't move
		switch guardPosition.direction {
		case Up:
			guardPosition.direction = Right
		case Right:
			guardPosition.direction = Down
		case Down:
			guardPosition.direction = Left
		case Left:
			guardPosition.direction = Up
		default:
			panic(fmt.Sprintf("Invalid guard direction: %v", guardPosition.direction))
		}
	}
}

func pointIsInWorld(world *[][]Node, point Point) bool {
	if point.y < 0 || len(*world)-1 < point.y {
		return false
	}

	return 0 <= point.x && point.x < len((*world)[point.y])
}

func printWorld(world *[][]Node, guardPosition *OrientatedPosition) {
	for y, row := range *world {
		for x, node := range row {
			if node.obstacle {
				fmt.Print("#")
			} else if guardPosition.position.x == x && guardPosition.position.y == y {
				switch guardPosition.direction {
				case Up:
					fmt.Print("^")
				case Right:
					fmt.Print(">")
				case Down:
					fmt.Print("V")
				case Left:
					fmt.Print("<")
				default:
					panic(fmt.Sprintf("Invalid guard direction: %v", guardPosition.direction))
				}
			} else if node.visited {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
