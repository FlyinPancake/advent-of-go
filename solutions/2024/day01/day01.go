package day01

import (
	"slices"
	"strconv"

	"github.com/flyinpancake/advent-of-go/aoc"
)

func init() {
	aoc.Register(2024, 1, &Solution{})
}

// Solution implements aoc.Solution for Year 2023 Day 01
type Solution struct{}

// Part1 solves part 1 of Day 01
func (s *Solution) Part1(input string) string {
	lines := aoc.Lines(input)

	leftList := []int{}
	rightList := []int{}

	for _, line := range lines {
		vals := aoc.Ints(line)
		leftList = append(leftList, vals[0])
		rightList = append(rightList, vals[1])
	}

	// Order the lists ascending
	slices.Sort(leftList)
	slices.Sort(rightList)

	totalDist := 0
	for i, _ := range leftList {
		totalDist += aoc.AbsDiff(leftList[i], rightList[i])
	}
	return strconv.Itoa(totalDist)
}

// Part2 solves part 2 of Day 01
func (s *Solution) Part2(input string) string {
	lines := aoc.Lines(input)
	leftList := []int{}
	rightList := []int{}

	for _, line := range lines {
		vals := aoc.Ints(line)
		leftList = append(leftList, vals[0])
		rightList = append(rightList, vals[1])
	}

	occurancesRight := make(map[int]int)
	for _, val := range rightList {
		occurancesRight[val]++
	}

	similarityScore := 0
	for _, val := range leftList {
		similarityScore += val * occurancesRight[val]
	}
	return strconv.Itoa(similarityScore)
}
