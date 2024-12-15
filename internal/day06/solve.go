package day06

import (
	"fmt"
	"slices"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

type Node struct {
	obstacle bool
	visited  []Direction // The direction(s) the guard was facing when they visited this node
}

var nullPoint = utils.Vector2{X: -1, Y: -1}

type Direction int

const (
	Unknown Direction = iota
	Up
	Right
	Down
	Left
)

type OrientatedPosition struct {
	position  utils.Vector2
	direction Direction
}

var nullPosition = OrientatedPosition{position: utils.Vector2{X: -1, Y: -1}, direction: Unknown}

func SolvePart1(input string, debug bool) int {
	world, guardPosition := parseInput(input)
	runSimulation(&world, &guardPosition, debug)

	visited := 0
	for _, row := range world {
		for _, node := range row {
			if len(node.visited) > 0 {
				visited++
			}
		}
	}

	return visited
}

func SolvePart2(input string, debug bool) int {
	currentWorld, guardPosition := parseInput(input)
	initialGuardPosition := OrientatedPosition{
		position:  utils.Vector2{X: guardPosition.position.X, Y: guardPosition.position.Y},
		direction: guardPosition.direction,
	}
	runSimulation(&currentWorld, &guardPosition, debug)

	loopingAlternateWorlds := 0

	for y, row := range currentWorld {
		for x, node := range row {
			// Ignore guard's initial position (can't place an obstacle here)
			if initialGuardPosition.position.X == x && initialGuardPosition.position.Y == y {
				continue
			}
			// No point trying to place obstacles on nodes that the guard never steps on
			if len(node.visited) <= 0 {
				continue
			}

			if debug {
				fmt.Printf("Trying obstacle at (%v, %v)\n", x, y)
			}
			alternateWorld, alternateGuardPosition := parseInput(input)
			alternateWorld[y][x].obstacle = true
			runSimulation(&alternateWorld, &alternateGuardPosition, debug)

			if pointIsInWorld(&alternateWorld, alternateGuardPosition.position) {
				if debug {
					fmt.Println("Loop found!")
				}

				loopingAlternateWorlds++
			}
		}
	}

	return loopingAlternateWorlds
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
				visited:  make([]Direction, 0, 4),
			}

			if char == "^" {
				utils.Assert(
					guardPosition == nullPosition,
					fmt.Sprintf("Found duplicate guard positions: %v | %v", guardPosition, OrientatedPosition{utils.Vector2{X: x, Y: y}, Up}),
				)
				guardPosition = OrientatedPosition{utils.Vector2{X: x, Y: y}, Up}
				world[y][x].visited = append(world[y][x].visited, guardPosition.direction)
			}
		}
	}

	utils.Assert(guardPosition != nullPosition, "Guard position not found")

	return world, guardPosition
}

func runSimulation(world *[][]Node, guardPosition *OrientatedPosition, debug bool) {
	if debug {
		printWorld(world, guardPosition)
	}

	for pointIsInWorld(world, guardPosition.position) {
		looped := simulationStep(world, guardPosition)
		if looped {
			break
		}
		if debug {
			printWorld(world, guardPosition)
		}
	}
}

func simulationStep(world *[][]Node, guardPosition *OrientatedPosition) bool {
	targetPoint := nullPoint

	switch guardPosition.direction {
	case Up:
		targetPoint = utils.Vector2{X: guardPosition.position.X, Y: guardPosition.position.Y - 1}
	case Right:
		targetPoint = utils.Vector2{X: guardPosition.position.X + 1, Y: guardPosition.position.Y}
	case Down:
		targetPoint = utils.Vector2{X: guardPosition.position.X, Y: guardPosition.position.Y + 1}
	case Left:
		targetPoint = utils.Vector2{X: guardPosition.position.X - 1, Y: guardPosition.position.Y}
	default:
		panic(fmt.Sprintf("Invalid guard direction: %v", guardPosition.direction))
	}

	targetIsInWorld := pointIsInWorld(world, targetPoint)

	if !targetIsInWorld || !(*world)[targetPoint.Y][targetPoint.X].obstacle {
		// Move 1 step forward
		guardPosition.position = targetPoint
		if targetIsInWorld {
			if slices.Contains((*world)[targetPoint.Y][targetPoint.X].visited, guardPosition.direction) {
				// Guard stepped onto a node they've already visited while facing the same direction
				// They're hit a loop
				return true
			}
			(*world)[targetPoint.Y][targetPoint.X].visited = append((*world)[targetPoint.Y][targetPoint.X].visited, guardPosition.direction)
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

		(*world)[guardPosition.position.Y][guardPosition.position.X].visited = append((*world)[guardPosition.position.Y][guardPosition.position.X].visited, guardPosition.direction)
	}

	return false
}

func pointIsInWorld(world *[][]Node, point utils.Vector2) bool {
	if point.Y < 0 || len(*world)-1 < point.Y {
		return false
	}

	return 0 <= point.X && point.X < len((*world)[point.Y])
}

var singleVisitPathDrawingMap = map[Direction]string{
	Up:    "|",
	Down:  "|",
	Left:  "-",
	Right: "-",
}
var doubleVisitPathDrawingMap = map[Direction]map[Direction]string{
	Up: {
		Right: "+",
		Down:  "|",
		Left:  "+",
	},
	Right: {
		Up:   "+",
		Down: "+",
		Left: "-",
	},
	Down: {
		Up:    "|",
		Right: "+",
		Left:  "+",
	},
	Left: {
		Up:    "+",
		Right: "-",
		Down:  "+",
	},
}

func printWorld(world *[][]Node, guardPosition *OrientatedPosition) {
	for y, row := range *world {
		for x, node := range row {
			if node.obstacle {
				fmt.Print("#")
			} else if guardPosition.position.X == x && guardPosition.position.Y == y {
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
			} else if len(node.visited) > 0 {
				switch len(node.visited) {
				case 1:
					fmt.Print(singleVisitPathDrawingMap[node.visited[0]])
				case 2:
					fmt.Print(doubleVisitPathDrawingMap[node.visited[0]][node.visited[1]])
				case 3, 4:
					fmt.Print("+")
				default:
					panic(fmt.Sprintf("Unexpected number of visit directions for node at (%v, %v): %v", x, y, len(node.visited)))
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
