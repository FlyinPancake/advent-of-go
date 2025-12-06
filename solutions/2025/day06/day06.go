package day06

import (
	"strconv"
	"strings"

	"github.com/flyinpancake/advent-of-go/aoc"
)

func init() {
	aoc.Register(2025, 6, &Solution{})
}

// Solution implements aoc.Solution for Year 2024 Day 02
type Solution struct{}

// Part1 solves part 1 of Day 02
func (s *Solution) Part1(input string) string {
	grid := [][]string{}

	lines := aoc.Lines(input)
	for _, line := range lines {
		lineFields := aoc.Fields(line)
		grid = append(grid, lineFields)
	}

	cols := len(grid[0])
	rows := len(grid)

	grandTotal := 0
	for ii := 0; ii < cols; ii++ {
		operand := grid[rows-1][ii]

		total := 0
		if operand == "*" {
			total = 1
		}
		for jj := 0; jj < rows-1; jj++ {
			if operand == "*" {
				total *= aoc.MustInt(grid[jj][ii])
			} else {
				total += aoc.MustInt(grid[jj][ii])
			}
		}
		grandTotal += total
	}

	return strconv.Itoa(grandTotal)
}

// Part2 solves part 2 of Day 02
func (s *Solution) Part2(input string) string {
	grid := [][]string{}
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	colWidths := []int{}

	operandLine := lines[len(lines)-1]

	cur := 0
	for _, v := range operandLine {
		if v == ' ' {
			cur++
		} else if cur != 0 {
			colWidths = append(colWidths, cur)
			cur = 0
		}
	}
	if cur != 0 {
		colWidths = append(colWidths, cur)
		cur = 0
	}

	colWidths[len(colWidths)-1]++

	for _, line := range lines {
		lineFields := []string{}

		prevWidth := 0
		for _, colWidth := range colWidths {
			lineFields = append(lineFields, line[prevWidth:prevWidth+colWidth])
			prevWidth += colWidth + 1
		}

		grid = append(grid, lineFields)
	}

	cols := len(colWidths)
	rows := len(grid)

	grandTotal := 0
	for ii := 0; ii < cols; ii++ {
		operand := strings.TrimSpace(grid[rows-1][ii])

		nums := []string{}
		maxLen := 0

		for jj := 0; jj < rows-1; jj++ {
			cur := grid[jj][ii]
			if len(cur) > maxLen {
				maxLen = len(cur)
			}
			nums = append(nums, cur)
		}

		// fmt.Printf("nums=%#v\n", nums)

		innerGrid := [][]rune{}
		for _, n := range nums {
			innerGrid = append(innerGrid, aoc.Chars(n))
		}

		finalNumbers := []int{}

		for col := maxLen - 1; col >= 0; col-- {
			n := ""
			for row := range innerGrid {
				n += string(innerGrid[row][col])
			}
			finalNumbers = append(finalNumbers, aoc.MustInt(n))
		}

		total := 0
		if operand == "*" {
			total = 1
		}
		for _, n := range finalNumbers {
			if operand == "*" {
				total *= n
			} else {
				total += n
			}
		}
		// fmt.Printf("%v=%d\n", finalNumbers, total)

		grandTotal += total
	}

	return strconv.Itoa(grandTotal)
}
