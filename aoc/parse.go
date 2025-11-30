package aoc

import (
	"strconv"
	"strings"
)

// Lines splits the input into lines, trimming any trailing newline
func Lines(input string) []string {
	input = strings.TrimSpace(input)
	if input == "" {
		return []string{}
	}
	return strings.Split(input, "\n")
}

// Paragraphs splits the input by blank lines (double newlines)
func Paragraphs(input string) []string {
	input = strings.TrimSpace(input)
	if input == "" {
		return []string{}
	}
	return strings.Split(input, "\n\n")
}

// Ints parses all integers from a line (separated by any non-digit characters)
func Ints(s string) []int {
	var result []int
	var current strings.Builder
	inNumber := false
	isNegative := false

	for i, r := range s {
		if r == '-' && !inNumber {
			// Check if next char is a digit
			if i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
				isNegative = true
				continue
			}
		}
		if r >= '0' && r <= '9' {
			current.WriteRune(r)
			inNumber = true
		} else if inNumber {
			if num, err := strconv.Atoi(current.String()); err == nil {
				if isNegative {
					num = -num
				}
				result = append(result, num)
			}
			current.Reset()
			inNumber = false
			isNegative = false
		} else {
			isNegative = false
		}
	}

	// Don't forget the last number
	if inNumber {
		if num, err := strconv.Atoi(current.String()); err == nil {
			if isNegative {
				num = -num
			}
			result = append(result, num)
		}
	}

	return result
}

// MustInt converts a string to int, panicking on error
func MustInt(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic("MustInt: " + err.Error())
	}
	return n
}

// MustInt64 converts a string to int64, panicking on error
func MustInt64(s string) int64 {
	n, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	if err != nil {
		panic("MustInt64: " + err.Error())
	}
	return n
}

// Fields splits a string by whitespace and returns non-empty fields
func Fields(s string) []string {
	return strings.Fields(s)
}

// IntFields splits by whitespace and converts each field to int
func IntFields(s string) []int {
	fields := strings.Fields(s)
	result := make([]int, len(fields))
	for i, f := range fields {
		result[i] = MustInt(f)
	}
	return result
}

// Chars splits a string into individual characters
func Chars(s string) []rune {
	return []rune(s)
}

// Grid parses the input into a 2D grid of runes
func Grid(input string) [][]rune {
	lines := Lines(input)
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

// IntGrid parses the input into a 2D grid of single-digit integers
func IntGrid(input string) [][]int {
	lines := Lines(input)
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, r := range line {
			grid[i][j] = int(r - '0')
		}
	}
	return grid
}

// SplitTwo splits a string by the given separator and returns exactly two parts
// Panics if there aren't exactly two parts after splitting
func SplitTwo(s, sep string) (string, string) {
	parts := strings.SplitN(s, sep, 2)
	if len(parts) != 2 {
		panic("SplitTwo: expected exactly 2 parts, got " + strconv.Itoa(len(parts)))
	}
	return parts[0], parts[1]
}

// CommaSeparated splits by comma and trims whitespace from each element
func CommaSeparated(s string) []string {
	parts := strings.Split(s, ",")
	result := make([]string, len(parts))
	for i, p := range parts {
		result[i] = strings.TrimSpace(p)
	}
	return result
}

// CommaSeparatedInts splits by comma and parses each as int
func CommaSeparatedInts(s string) []int {
	parts := CommaSeparated(s)
	result := make([]int, len(parts))
	for i, p := range parts {
		result[i] = MustInt(p)
	}
	return result
}
