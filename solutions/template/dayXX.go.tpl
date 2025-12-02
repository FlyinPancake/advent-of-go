package day{{.DayPadded}}

import (
	"github.com/flyinpancake/advent-of-go/aoc"
)

func init() {
	aoc.Register({{.Year}}, {{.Day}}, &Solution{})
}

// Solution implements aoc.Solution for Year 2024 Day 02
type Solution struct{}

// Part1 solves part 1 of Day 02
func (s *Solution) Part1(input string) string {
	lines := aoc.Lines(input)
	_ = lines // Use the parsed input

	// TODO: Implement Part 1
	return "not implemented"
}

// Part2 solves part 2 of Day 02
func (s *Solution) Part2(input string) string {
	lines := aoc.Lines(input)
	_ = lines // Use the parsed input

	// TODO: Implement Part 2
	return "not implemented"
}
