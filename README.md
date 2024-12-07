# Advent of Code 2024
https://adventofcode.com/2024

This year, I'm trying [go](https://go.dev/) (v1.23.3)

## Setup
Just install go

## Running
Run `go run ./... dayXX partY <path to input file>`

### Debugging
To run with debug messages, run `go run ./... --debug dayXX partY <path to input
file>`  
**NOTE** the `--debug` flag _has_ to come before the other arguments, due to
limitations in the `flag` module

## Testing
This project has tests that run the solvers against the example inputs and
answers provided with each puzzle
1. Run `go test ./internal/...` to run tests for all puzzles
