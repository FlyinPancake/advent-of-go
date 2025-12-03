package day03

import (
	"testing"

	"github.com/flyinpancake/advent-of-go/aoc"
)

var exampleInput = `
987654321111111
811111111111119
234234234234278
818181911112111
`

func TestPart1(t *testing.T) {
	testCases := []aoc.TestCase{
		{
			Name:     "example",
			Input:    exampleInput,
			Expected: "357",
		},
	}
	aoc.RunPart1Tests(t, &Solution{}, testCases)
}

func TestPart2(t *testing.T) {
	testCases := []aoc.TestCase{
		{
			Name:     "example",
			Input:    exampleInput,
			Expected: "3121910778619",
		},
	}
	aoc.RunPart2Tests(t, &Solution{}, testCases)
}

func BenchmarkPart1(b *testing.B) {
	input, err := aoc.ReadInput(2025, 3)
	if err != nil {
		b.Skip("input not available")
	}
	aoc.BenchmarkPart1(b, &Solution{}, input)
}

func BenchmarkPart2(b *testing.B) {
	input, err := aoc.ReadInput(2025, 3)
	if err != nil {
		b.Skip("input not available")
	}
	aoc.BenchmarkPart2(b, &Solution{}, input)
}
