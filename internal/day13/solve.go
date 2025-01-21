package day13

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

const buttonACost = 3
const buttonBCost = 1

func SolvePart1(input string, debug bool) int {
	equationPairs := parseInput(input)

	tokenCost := 0
	for _, equationPair := range equationPairs {
		if debug {
			fmt.Println(equationPair)
		}

		if tokens, err := solveEquations(equationPair, debug); err == nil {
			if debug {
				fmt.Printf("Token cost: %v\n", tokens)
			}
			tokenCost += tokens
		}

		if debug {
			fmt.Println()
		}
	}

	return tokenCost
}

func parseInput(input string) []SimultaneousEquationPair {
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
	equations := make([]SimultaneousEquationPair, 0, len(lines)/4)
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

			equations = append(
				equations,
				SimultaneousEquationPair{
					equationA: Equation{
						lhs: map[Unknown]int{
							UnknownA: buttonA.X,
							UnknownB: buttonB.X,
						},
						rhs: map[Unknown]int{
							Scalar: prizePosition.X,
						},
					},
					equationB: Equation{
						lhs: map[Unknown]int{
							UnknownA: buttonA.Y,
							UnknownB: buttonB.Y,
						},
						rhs: map[Unknown]int{
							Scalar: prizePosition.Y,
						},
					},
				},
			)
		} else {
			panic(fmt.Sprintf("Failed to parse line: %v", line))
		}
	}

	return equations
}

func solveEquations(equationPair SimultaneousEquationPair, debug bool) (int, error) {
	// multiply equation A by the coefficient of equation B's b term and vice versa
	bCoefficientInA, err := equationPair.equationA.findCoefficient(UnknownB)
	utils.AssertNoError(err)
	bCoefficientInB, err := equationPair.equationB.findCoefficient(UnknownB)
	utils.AssertNoError(err)

	rearrangePairForB := SimultaneousEquationPair{
		equationA: equationPair.equationA.multiply(bCoefficientInB),
		equationB: equationPair.equationB.multiply(bCoefficientInA),
	}

	if debug {
		fmt.Println(rearrangePairForB)
	}

	// rearrange both in terms of b
	rearrangedA, err := rearrangePairForB.equationA.rearrangeInTermsOf(UnknownB)
	utils.AssertNoError(err)
	rearrangedB, err := rearrangePairForB.equationB.rearrangeInTermsOf(UnknownB)
	utils.AssertNoError(err)

	rearrangePairForB = SimultaneousEquationPair{
		equationA: rearrangedA,
		equationB: rearrangedB,
	}

	if debug {
		fmt.Println(rearrangePairForB)
	}

	// make them equal each other
	aSideWithoutB, err := rearrangePairForB.equationA.eliminate(UnknownB)
	utils.AssertNoError(err)
	bSideWithoutB, err := rearrangePairForB.equationB.eliminate(UnknownB)
	utils.AssertNoError(err)

	aOnlyEquation := Equation{
		lhs: aSideWithoutB,
		rhs: bSideWithoutB,
	}

	if debug {
		fmt.Println(aOnlyEquation)
	}

	// rearrange in terms of a
	aOnlyEquation, err = aOnlyEquation.rearrangeInTermsOf(UnknownA)
	utils.AssertNoError(err)

	if debug {
		fmt.Println(aOnlyEquation)
	}

	// reduce to 1a
	aValue, err := aOnlyEquation.evaluate(UnknownA)
	err = handleEvaluationError(err, debug, "a does not have an integer value")
	if err != nil {
		return -1, err
	}

	if debug {
		fmt.Printf("a = %v\n", aValue)
	}

	// plug a into original equations
	substitutedA := SimultaneousEquationPair{
		equationA: equationPair.equationA.substitute(UnknownA, aValue),
		equationB: equationPair.equationB.substitute(UnknownA, aValue),
	}

	if debug {
		fmt.Println(substitutedA)
	}

	// rearrange both in terms of b
	rearrangedA, err = substitutedA.equationA.rearrangeInTermsOf(UnknownB)
	utils.AssertNoError(err)
	rearrangedB, err = substitutedA.equationB.rearrangeInTermsOf(UnknownB)
	utils.AssertNoError(err)

	rearrangePairForB = SimultaneousEquationPair{
		equationA: rearrangedA,
		equationB: rearrangedB,
	}

	if debug {
		fmt.Println(rearrangePairForB)
	}

	// reduce to 1b
	bValueFromA, err := rearrangePairForB.equationA.evaluate(UnknownB)
	err = handleEvaluationError(err, debug, "b does not have an integer value")
	if err != nil {
		return -1, err
	}
	bValueFromB, err := rearrangePairForB.equationB.evaluate(UnknownB)
	err = handleEvaluationError(err, debug, "b does not have an integer value")
	if err != nil {
		return -1, err
	}

	// if both bs are not the same, panic
	utils.Assert(bValueFromA == bValueFromB, "Equations have different values for B")
	bValue := bValueFromA

	if debug {
		fmt.Printf("b = %v\n", bValue)
	}

	// calculate token cost and return
	return (aValue * buttonACost) + (bValue * buttonBCost), nil
}

func handleEvaluationError(err error, debug bool, debugMessage string) error {
	if err == nil {
		return nil
	}

	var nie NonIntegerError
	if errors.As(err, &nie) {
		if debug {
			fmt.Println(debugMessage)
		}
		return errors.New("equations do not have an integer solution")
	} else {
		panic(err.Error())
	}
}
