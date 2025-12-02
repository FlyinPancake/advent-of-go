package day02

import (
	"testing"

	"github.com/flyinpancake/advent-of-go/aoc"
)

func TestValidateIdRepeatsTwice(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"repeated twice - 1212", "1212", false},
		{"repeated twice - 1010", "1010", false},
		{"valid - 101", "101", true},
		{"single digit repeated", "1111", false},
		{"no repeat", "1234", true},
		{"triple repeat", "123123123", true},
		{"partial repeat valid", "12123", true},
		{"two digits same", "11", false},
		{"single digit", "1", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validateIdRepeatsTwice(tt.input)
			if result != tt.expected {
				t.Errorf("validateId(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

var exampleInput = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestPart1(t *testing.T) {
	testCases := []aoc.TestCase{
		{
			Name:     "example",
			Input:    exampleInput,
			Expected: "1227775554",
		},
	}
	aoc.RunPart1Tests(t, &Solution{}, testCases)
}

func TestPart2(t *testing.T) {
	testCases := []aoc.TestCase{
		{
			Name:     "example",
			Input:    exampleInput,
			Expected: "4174379265",
		},
	}
	aoc.RunPart2Tests(t, &Solution{}, testCases)
}

func BenchmarkPart1(b *testing.B) {
	input, err := aoc.ReadInput(2025, 2)
	if err != nil {
		b.Skip("input not available")
	}
	aoc.BenchmarkPart1(b, &Solution{}, input)
}

func BenchmarkPart2(b *testing.B) {
	input, err := aoc.ReadInput(2025, 2)
	if err != nil {
		b.Skip("input not available")
	}
	aoc.BenchmarkPart2(b, &Solution{}, input)
}
