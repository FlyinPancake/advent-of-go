package day06

import (
	"testing"

	"github.com/flyinpancake/advent-of-go/aoc"
)

var exampleInput = `
123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +
`

func TestPart1(t *testing.T) {
	testCases := []aoc.TestCase{
		{
			Name:     "example",
			Input:    exampleInput,
			Expected: "4277556",
		},
	}
	aoc.RunPart1Tests(t, &Solution{}, testCases)
}

func TestPart2(t *testing.T) {
	diskExampleInput, _ := aoc.ReadExample(2025, 6, "")

	testCases := []aoc.TestCase{
		{
			Name:     "example",
			Input:    diskExampleInput,
			Expected: "3263827",
		},
	}
	aoc.RunPart2Tests(t, &Solution{}, testCases)
}

func BenchmarkPart1(b *testing.B) {
	input, err := aoc.ReadInput(2025, 6)
	if err != nil {
		b.Skip("input not available")
	}
	aoc.BenchmarkPart1(b, &Solution{}, input)
}

func BenchmarkPart2(b *testing.B) {
	input, err := aoc.ReadInput(2025, 6)
	if err != nil {
		b.Skip("input not available")
	}
	aoc.BenchmarkPart2(b, &Solution{}, input)
}
