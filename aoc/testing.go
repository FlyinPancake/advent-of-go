package aoc

import (
	"fmt"
	"testing"
)

// TestCase represents a single test case for a solution
type TestCase struct {
	Name     string
	Input    string
	Expected string
}

// RunPart1Tests runs all Part 1 test cases for a solution
func RunPart1Tests(t *testing.T, solution Solution, testCases []TestCase) {
	t.Helper()
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := solution.Part1(tc.Input)
			if result != tc.Expected {
				t.Errorf("Part1() = %q, want %q", result, tc.Expected)
			}
		})
	}
}

// RunPart2Tests runs all Part 2 test cases for a solution
func RunPart2Tests(t *testing.T, solution Solution, testCases []TestCase) {
	t.Helper()
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := solution.Part2(tc.Input)
			if result != tc.Expected {
				t.Errorf("Part2() = %q, want %q", result, tc.Expected)
			}
		})
	}
}

// BenchmarkPart1 creates a benchmark function for Part 1
func BenchmarkPart1(b *testing.B, solution Solution, input string) {
	b.Helper()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solution.Part1(input)
	}
}

// BenchmarkPart2 creates a benchmark function for Part 2
func BenchmarkPart2(b *testing.B, solution Solution, input string) {
	b.Helper()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solution.Part2(input)
	}
}

// AssertEqual is a simple assertion helper for testing
func AssertEqual(t *testing.T, got, want interface{}, msgPrefix string) {
	t.Helper()
	if got != want {
		t.Errorf("%s: got %v, want %v", msgPrefix, got, want)
	}
}

// TestSolution provides a convenient way to test a full solution with example input
type TestSolution struct {
	Solution     Solution
	ExampleInput string
	Part1Result  string
	Part2Result  string
}

// Run runs both parts and prints results for quick verification
func (ts *TestSolution) Run() {
	fmt.Println("=== Testing Solution ===")
	if ts.Part1Result != "" {
		result := ts.Solution.Part1(ts.ExampleInput)
		status := "✓"
		if result != ts.Part1Result {
			status = "✗"
		}
		fmt.Printf("Part 1: %s (expected: %s) %s\n", result, ts.Part1Result, status)
	}
	if ts.Part2Result != "" {
		result := ts.Solution.Part2(ts.ExampleInput)
		status := "✓"
		if result != ts.Part2Result {
			status = "✗"
		}
		fmt.Printf("Part 2: %s (expected: %s) %s\n", result, ts.Part2Result, status)
	}
}
