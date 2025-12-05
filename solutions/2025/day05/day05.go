package day05

import (
	"fmt"
	"strconv"

	"github.com/flyinpancake/advent-of-go/aoc"
)

func init() {
	aoc.Register(2025, 5, &Solution{})
}

// Solution implements aoc.Solution for Year 2024 Day 02
type Solution struct{}

// Part1 solves part 1 of Day 02
func (s *Solution) Part1(input string) string {
	freshRangesStr, itemIdsStr := aoc.SplitTwo(input, "\n\n")

	freshRanges := [][]int{}

	for _, line := range aoc.Lines(freshRangesStr) {
		lowerStr, upperStr := aoc.SplitTwo(line, "-")
		lower := aoc.MustInt(lowerStr)
		upper := aoc.MustInt(upperStr)
		freshRanges = append(freshRanges, []int{lower, upper})
	}

	freshItemsCount := 0

	for _, line := range aoc.Lines(itemIdsStr) {
		itemId := aoc.MustInt(line)
		isFresh := false
		for _, fr := range freshRanges {
			if itemId >= fr[0] && itemId <= fr[1] {
				isFresh = true
				break
			}
		}
		if isFresh {
			freshItemsCount++
		}
	}

	return strconv.Itoa(freshItemsCount)
}

func mergeRanges(a, b aoc.Range) (aoc.Range, aoc.Range) {
	// check if either of the ranges is empty
	if a.IsEmpty() || b.IsEmpty() {
		return a, b
	}

	// check if b is wholly outside of a
	if b.Start > a.End || b.End < a.Start {
		// fmt.Printf("a outside of b -- a=%v; b=%v\n", a, b)

		return a, b
	}

	// check if either one contains the other
	//   check if a contains b
	if a.Start <= b.Start && a.End >= b.End {
		// fmt.Printf("a contains b -- a=%v; b=%v\n", a, b)
		return a, aoc.EmptyRange()
	}
	//   check if b contains a
	if b.Start <= a.Start && b.End >= a.End {
		// fmt.Printf("b contains a -- a=%v; b=%v\n", a, b)
		return aoc.EmptyRange(), b
	}

	// check if b overlaps the end of a
	if a.End >= b.Start && a.End < b.End {
		// fmt.Printf("b overlaps the end of a -- a=%v; b=%v\n", a, b)
		return aoc.EmptyRange(), aoc.NewInclusiveRange(a.Start, b.End)
	}

	// check if b overlaps the start of a
	if b.Start <= a.Start {
		// fmt.Printf("b overlaps the start of a -- a=%v; b=%v\n", a, b)

		return aoc.EmptyRange(), aoc.NewInclusiveRange(b.Start, a.End)
	}

	panic(fmt.Sprintf("not all cases are handled -- a=%v; b=%v\n", a, b))
}

// Part2 solves part 2 of Day 02
func (s *Solution) Part2(input string) string {
	freshRangesStr, _ := aoc.SplitTwo(input, "\n\n")

	freshIdRanges := []aoc.Range{}

	lines := aoc.Lines(freshRangesStr)

	for _, line := range lines {
		lowerStr, upperStr := aoc.SplitTwo(line, "-")
		lower := aoc.MustInt(lowerStr)
		upper := aoc.MustInt(upperStr)
		rang := aoc.NewInclusiveRange(lower, upper)
		freshIdRanges = append(freshIdRanges, rang)
	}

	for ii, rang := range freshIdRanges {
		if rang.IsEmpty() {
			continue
		}
		for jj := ii + 1; jj < len(freshIdRanges); jj++ {
			freshIdRanges[ii], freshIdRanges[jj] = mergeRanges(freshIdRanges[ii], freshIdRanges[jj])
		}
		// fmt.Println(freshIdRanges)
	}

	freshIds := 0
	for _, rang := range freshIdRanges {
		freshIds += rang.Len()
	}

	return strconv.Itoa(freshIds)
}
