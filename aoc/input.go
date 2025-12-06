package aoc

import (
	"fmt"
	"os"
	"path/filepath"
)

// ReadInput reads the input file for a given year and day
func ReadInput(year, day int) (string, error) {
	// Try multiple possible locations for the input file
	paths := []string{
		filepath.Join("inputs", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d.txt", day)),
		filepath.Join("inputs", fmt.Sprintf("day%02d.txt", day)),
		filepath.Join(fmt.Sprintf("day%02d", day), "input.txt"),
	}

	var lastErr error
	for _, path := range paths {
		content, err := os.ReadFile(path)
		if err == nil {
			return string(content), nil
		}
		lastErr = err
	}

	return "", fmt.Errorf("could not find input file for year %d day %d: %w", year, day, lastErr)
}

// ReadInputString reads the input and returns it as a trimmed string
func ReadInputString(year, day int) (string, error) {
	input, err := ReadInput(year, day)
	if err != nil {
		return "", err
	}
	return TrimInput(input), nil
}

// ReadExample reads an example input file for a given year and day
func ReadExample(year, day int, name string) (string, error) {
	if name == "" {
		name = "example"
	}

	paths := []string{
		filepath.Join("inputs", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d_%s.txt", day, name)),
		filepath.Join("..", "..", "..", "inputs", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d_%s.txt", day, name)),
		filepath.Join("inputs", fmt.Sprintf("day%02d_%s.txt", day, name)),
		filepath.Join(fmt.Sprintf("day%02d", day), fmt.Sprintf("%s.txt", name)),
	}

	var lastErr error
	for _, path := range paths {
		content, err := os.ReadFile(path)
		if err == nil {
			return string(content), nil
		}
		lastErr = err
	}

	return "", fmt.Errorf("could not find example file for year %d day %d: %w", year, day, lastErr)
}

// TrimInput removes leading and trailing whitespace from input
func TrimInput(input string) string {
	// Trim leading/trailing whitespace but preserve internal structure
	start := 0
	end := len(input)

	// Trim leading whitespace
	for start < end && (input[start] == ' ' || input[start] == '\t' || input[start] == '\n' || input[start] == '\r') {
		start++
	}

	// Trim trailing whitespace
	for end > start && (input[end-1] == ' ' || input[end-1] == '\t' || input[end-1] == '\n' || input[end-1] == '\r') {
		end--
	}

	return input[start:end]
}
