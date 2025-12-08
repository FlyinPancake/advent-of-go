package day08

import (
	"testing"

	"github.com/flyinpancake/advent-of-go/aoc"
)

var exampleInput = `
162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
`

func TestPart1(t *testing.T) {
	testCases := []aoc.TestCase{
		{
			Name:     "example",
			Input:    exampleInput,
			Expected: "40",
		},
	}
	aoc.RunPart1Tests(t, &Solution{}, testCases)
}

func TestPart2(t *testing.T) {
	testCases := []aoc.TestCase{
		{
			Name:     "example",
			Input:    exampleInput,
			Expected: "expected_result",
		},
	}
	aoc.RunPart2Tests(t, &Solution{}, testCases)
}

func BenchmarkPart1(b *testing.B) {
	input, err := aoc.ReadInput(2025, 8)
	if err != nil {
		b.Skip("input not available")
	}
	aoc.BenchmarkPart1(b, &Solution{}, input)
}

func BenchmarkPart2(b *testing.B) {
	input, err := aoc.ReadInput(2025, 8)
	if err != nil {
		b.Skip("input not available")
	}
	aoc.BenchmarkPart2(b, &Solution{}, input)
}
