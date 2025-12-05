package day05

import (
	"testing"

	"github.com/flyinpancake/advent-of-go/aoc"
)

var exampleInput = `
3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

func TestPart1(t *testing.T) {
	testCases := []aoc.TestCase{
		{
			Name:     "example",
			Input:    exampleInput,
			Expected: "3",
		},
	}
	aoc.RunPart1Tests(t, &Solution{}, testCases)
}

func TestPart2(t *testing.T) {
	testCases := []aoc.TestCase{
		{
			Name:     "example",
			Input:    exampleInput,
			Expected: "14",
		},
	}
	aoc.RunPart2Tests(t, &Solution{}, testCases)
}

func BenchmarkPart1(b *testing.B) {
	input, err := aoc.ReadInput(2025, 5)
	if err != nil {
		b.Skip("input not available")
	}
	aoc.BenchmarkPart1(b, &Solution{}, input)
}

func BenchmarkPart2(b *testing.B) {
	input, err := aoc.ReadInput(2025, 5)
	if err != nil {
		b.Skip("input not available")
	}
	aoc.BenchmarkPart2(b, &Solution{}, input)
}
