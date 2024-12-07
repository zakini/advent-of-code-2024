package day07

import (
	"fmt"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

type Operation int

const (
	Invalid Operation = iota
	Add
	Multiply
	Concat
)

type Calibration struct {
	target   int
	operands []int
}

func (op Operation) String() string {
	switch op {
	case Add:
		return "+"
	case Multiply:
		return "✕"
	case Concat:
		return "||"
	default:
		return "?"
	}
}

func SolvePart1(input string, debug bool) int {
	return solve(input, []Operation{Add, Multiply}, debug)
}

func SolvePart2(input string, debug bool) int {
	return solve(input, []Operation{Add, Multiply, Concat}, debug)
}

func solve(input string, possibleOperations []Operation, debug bool) int {
	calibrations := parseInput(input)

	totalCalibrationResult := 0
	for _, calibration := range calibrations {
		operationLists := generateOperationLists(possibleOperations, len(calibration.operands)-1)

		for _, operationList := range operationLists {
			result := applyOperationList(calibration.operands, operationList)
			if debug {
				printOperationListApplication(calibration, operationList, result)
			}

			if calibration.target == result {
				totalCalibrationResult += result
				break
			}
		}
	}

	return totalCalibrationResult
}

func parseInput(input string) []Calibration {
	lines := strings.Split(input, "\n")

	calibrations := make([]Calibration, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ":")
		target := utils.ParseIntAndAssert(parts[0])
		operandStrings := strings.Fields(parts[1])
		operands := make([]int, len(operandStrings))

		for j, operandString := range operandStrings {
			operands[j] = utils.ParseIntAndAssert(operandString)
		}

		calibrations[i] = Calibration{target, operands}
	}

	return calibrations
}

func generateOperationLists(possibleOperations []Operation, length int) [][]Operation {
	utils.Assert(length >= 0, "Cannot operation lists with a negative length")

	if length == 0 {
		// empty list
		return [][]Operation{}
	}

	operationLists := make([][]Operation, 0)

	if length == 1 {
		for _, op := range possibleOperations {
			operationLists = append(operationLists, []Operation{op})
		}
	} else {
		for _, op := range possibleOperations {
			subLists := generateOperationLists(possibleOperations, length-1)
			for i := range subLists {
				subLists[i] = append(subLists[i], op)
			}
			operationLists = append(operationLists, subLists...)
		}
	}

	return operationLists
}

func applyOperationList(operands []int, operationList []Operation) int {
	result := 0
	for i, op := range operationList {
		if i == 0 {
			result = applyOperation(operands[i], operands[i+1], op)
		} else {
			result = applyOperation(result, operands[i+1], op)
		}
	}

	return result
}

func applyOperation(a int, b int, op Operation) int {
	switch op {
	case Add:
		return a + b
	case Multiply:
		return a * b
	case Concat:
		return utils.ParseIntAndAssert(fmt.Sprintf("%v%v", a, b))
	default:
		panic(fmt.Sprintf("Attempted to apply invalid operation: %v", op))
	}
}

func printOperationListApplication(calibration Calibration, operationList []Operation, result int) {
	fmt.Print(calibration.operands[0])
	for i, op := range operationList {
		fmt.Printf(" %v %v", op, calibration.operands[i+1])
	}

	if calibration.target == result {
		fmt.Print(" == ")
	} else {
		fmt.Print(" != ")
	}

	fmt.Println(result)
}
