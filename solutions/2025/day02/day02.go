package day02

import (
	"strconv"
	"strings"

	"github.com/flyinpancake/advent-of-go/aoc"
)

func init() {
	aoc.Register(2025, 2, &Solution{})
}

// Solution implements aoc.Solution for Year 2024 Day 02
type Solution struct{}

// ID is invalid if it is a sequence repeated at least twice like 1212 is invalid since it's 12 repeated twice, but 101 is valid
func validateIdRepeatsTwice(s string) bool {
	idLength := len(s)
	possibleRepeats := 2

	for possibleSeqLen := 1; possibleSeqLen <= (idLength / 2); possibleSeqLen++ {
		if possibleSeqLen*possibleRepeats != idLength {
			continue
		}
		// check remainder after possible seq
		possibleSeq := s[0:possibleSeqLen]
		if strings.Repeat(possibleSeq, possibleRepeats) == s[:possibleRepeats*possibleSeqLen] {
			return false
		}
	}
	return true
}

// Part1 solves part 1 of Day 02
func (s *Solution) Part1(input string) string {
	ranges := aoc.CommaSeparated(aoc.TrimInput(input))
	total := 0
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		rangeStart := aoc.MustInt(parts[0])
		rangeEnd := aoc.MustInt(parts[1])
		for id := rangeStart; id <= rangeEnd; id++ {
			if !validateIdRepeatsTwice(strconv.Itoa(id)) {
				// fmt.Printf("%d\n", id)
				total += id
			}
		}
	}

	return strconv.Itoa(total)
}

// ID is invalid if it is a sequence repeated at least twice like 1212 is invalid since it's 12 repeated twice, but 101 is valid
func validateIdRepeatsAtLeastTwice(s string) bool {
	idLength := len(s)
	for possibleSeqLen := 1; possibleSeqLen <= (idLength / 2); possibleSeqLen++ {
		possibleRepeats := idLength / possibleSeqLen
		if possibleSeqLen*possibleRepeats != idLength {
			continue
		}

		// check remainder after possible seq
		possibleSeq := s[0:possibleSeqLen]
		if strings.Repeat(possibleSeq, possibleRepeats) == s[:possibleRepeats*possibleSeqLen] {
			return false
		}
	}
	return true
}

// Part2 solves part 2 of Day 02
func (s *Solution) Part2(input string) string {
	ranges := aoc.CommaSeparated(aoc.TrimInput(input))
	total := 0
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		rangeStart := aoc.MustInt(parts[0])
		rangeEnd := aoc.MustInt(parts[1])
		for id := rangeStart; id <= rangeEnd; id++ {
			if !validateIdRepeatsAtLeastTwice(strconv.Itoa(id)) {
				// fmt.Printf("%d\n", id)
				total += id
			}
		}
	}

	return strconv.Itoa(total)
}
