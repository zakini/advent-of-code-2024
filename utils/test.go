package utils

import (
	"testing"
)

func TestAgainstExample(t *testing.T, solver Solver, filePath string, expected int) {
	input := LoadInputFile(filePath)
	actual := solver(input)

	if expected != actual {
		t.Fatalf("expected: %d | actual: %d", expected, actual)
	}
}
