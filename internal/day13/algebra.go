package day13

import (
	"errors"
	"fmt"
	"strings"
)

type Unknown int

const (
	UnknownA Unknown = iota
	UnknownB
	Scalar
)

func (unknown Unknown) String() string {
	switch unknown {
	case UnknownA:
		return "a"
	case UnknownB:
		return "b"
	case Scalar:
		return ""
	}

	panic("Invalid unknown")
}

type Equation struct {
	lhs map[Unknown]int
	rhs map[Unknown]int
}

func (eq Equation) String() string {
	lhs := make([]string, 0, len(eq.lhs))
	for unknown, coefficient := range eq.lhs {
		if coefficient == 1 {
			lhs = append(lhs, unknown.String())
		} else if coefficient != 0 {
			lhs = append(lhs, fmt.Sprintf("%v%v", coefficient, unknown))
		}
	}

	if len(lhs) == 0 {
		lhs = append(lhs, "0")
	}

	rhs := make([]string, 0, len(eq.rhs))
	for unknown, coefficient := range eq.rhs {
		if coefficient != 0 {
			rhs = append(rhs, fmt.Sprintf("%v%v", coefficient, unknown))
		}
	}

	if len(rhs) == 0 {
		rhs = append(rhs, "0")
	}

	return fmt.Sprintf(
		"%v = %v",
		strings.Join(lhs, " + "),
		strings.Join(rhs, " + "),
	)
}

func (eq Equation) add(coefficient int, unknown Unknown) Equation {
	lhs := eq.lhs
	lhs[unknown] += coefficient

	rhs := eq.rhs
	rhs[unknown] += coefficient

	return Equation{
		lhs: lhs,
		rhs: rhs,
	}
}

func (eq Equation) subtract(coefficient int, unknown Unknown) Equation {
	return eq.add(-coefficient, unknown)
}

func (eq Equation) multiply(value int) Equation {
	return Equation{
		lhs: map[Unknown]int{
			UnknownA: eq.lhs[UnknownA] * value,
			UnknownB: eq.lhs[UnknownB] * value,
			Scalar:   eq.lhs[Scalar] * value,
		},
		rhs: map[Unknown]int{
			UnknownA: eq.rhs[UnknownA] * value,
			UnknownB: eq.rhs[UnknownB] * value,
			Scalar:   eq.rhs[Scalar] * value,
		},
	}
}

func (eq Equation) divide(value int) (Equation, error) {
	// Check if any division would result in a non integer
	// For the purposes of this puzzle, we only care about integer solutions
	modulusSum := 0

	modulusSum += eq.lhs[UnknownA] % value
	modulusSum += eq.lhs[UnknownB] % value
	modulusSum += eq.lhs[Scalar] % value
	modulusSum += eq.rhs[UnknownA] % value
	modulusSum += eq.rhs[UnknownB] % value
	modulusSum += eq.rhs[Scalar] % value

	if modulusSum != 0 {
		return Equation{}, errors.New("division would result in non-integer equation")
	}

	divided := Equation{
		lhs: map[Unknown]int{
			UnknownA: eq.lhs[UnknownA] / value,
			UnknownB: eq.lhs[UnknownB] / value,
			Scalar:   eq.lhs[Scalar] / value,
		},
		rhs: map[Unknown]int{
			UnknownA: eq.rhs[UnknownA] / value,
			UnknownB: eq.rhs[UnknownB] / value,
			Scalar:   eq.rhs[Scalar] / value,
		},
	}

	return divided, nil
}

func (eq Equation) findCoefficient(unknown Unknown) (int, error) {
	lhsCoefficient := eq.lhs[unknown]
	rhsCoefficient := eq.rhs[unknown]

	if lhsCoefficient == 0 && rhsCoefficient == 0 {
		return 0, fmt.Errorf("Equation has no %v term", unknown)
	}

	if lhsCoefficient != 0 && rhsCoefficient != 0 {
		return 0, fmt.Errorf("Equation has %v on both sides", unknown)
	}

	if lhsCoefficient != 0 {
		return lhsCoefficient, nil
	}

	return rhsCoefficient, nil
}

func (eq Equation) rearrangeInTermsOf(unknown Unknown) (Equation, error) {
	lhsCoefficient := eq.lhs[unknown]
	rhsCoefficient := eq.rhs[unknown]

	if lhsCoefficient == 0 && rhsCoefficient == 0 {
		return Equation{}, fmt.Errorf("Equation has no %v term", unknown)
	}

	if lhsCoefficient == 0 {
		eq = eq.flip()
	}

	for _, otherUnknown := range []Unknown{UnknownA, UnknownB, Scalar} {
		if otherUnknown == unknown {
			eq = eq.subtract(eq.rhs[otherUnknown], otherUnknown)
		} else {
			eq = eq.subtract(eq.lhs[otherUnknown], otherUnknown)
		}
	}

	return eq, nil
}

func (eq Equation) flip() Equation {
	return Equation{
		lhs: eq.rhs,
		rhs: eq.lhs,
	}
}

func (eq Equation) eliminate(unknown Unknown) (map[Unknown]int, error) {
	lhsCoefficient := eq.lhs[unknown]
	rhsCoefficient := eq.rhs[unknown]

	if lhsCoefficient == 0 && rhsCoefficient == 0 {
		return map[Unknown]int{}, fmt.Errorf("Equation has no %v term", unknown)
	}

	if lhsCoefficient == 0 {
		return eq.lhs, nil
	}

	return eq.rhs, nil
}

func (eq Equation) evaluate(unknown Unknown) (int, error) {
	leftTerms := 0
	unknownLeftTerms := make([]Unknown, 0, 2)
	for unknown, coefficient := range eq.lhs {
		if coefficient != 0 {
			leftTerms++

			if unknown != Scalar {
				unknownLeftTerms = append(unknownLeftTerms, unknown)
			}
		}
	}

	rightTerms := 0
	unknownRightTerms := make([]Unknown, 0, 2)
	for unknown, coefficient := range eq.rhs {
		if coefficient != 0 {
			rightTerms++

			if unknown != Scalar {
				unknownRightTerms = append(unknownRightTerms, unknown)
			}
		}
	}

	// There's probably a way to simplify more complex equations, but for our purposes, we shouldn't need to
	if leftTerms > 1 || rightTerms > 1 {
		return 0, errors.New("Equation has too many terms to evaluate")
	}

	if len(unknownLeftTerms) == 1 && len(unknownRightTerms) == 1 {
		return 0, errors.New("Equation has not been properly rearranged")
	}

	if len(unknownLeftTerms) == 0 && len(unknownRightTerms) == 0 {
		return 0, errors.New("Equation is equivalent to a constant")
	}

	var result Equation
	var err error
	if len(unknownLeftTerms) == 1 {
		result, err = eq.divide(eq.lhs[unknownLeftTerms[0]])
	} else {
		result, err = eq.divide(eq.rhs[unknownRightTerms[0]])
	}

	if err != nil {
		return 0, NonIntegerError{}
	}

	return result.findCoefficient(Scalar)
}

func (eq Equation) substitute(unknown Unknown, value int) Equation {
	lhs := eq.lhs
	lhs[Scalar] = eq.lhs[Scalar] + (eq.lhs[unknown] * value)
	lhs[unknown] = 0

	rhs := eq.rhs
	rhs[Scalar] = eq.rhs[Scalar] + (eq.rhs[unknown] * value)
	rhs[unknown] = 0

	return Equation{
		lhs: lhs,
		rhs: rhs,
	}
}

type SimultaneousEquationPair struct {
	equationA Equation
	equationB Equation
}

func (eqs SimultaneousEquationPair) String() string {
	return fmt.Sprintf("%v | %v", eqs.equationA, eqs.equationB)
}

type NonIntegerError struct{}

func (err NonIntegerError) Error() string {
	return "Equation cannot be simplified to an integer value"
}
