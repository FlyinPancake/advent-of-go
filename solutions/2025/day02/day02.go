package day02

import (
	"fmt"
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
	if idLength%2 != 0 {
		return true
	}
	half := idLength / 2

	return s[:half] != s[half:]
}

// Part1 solves part 1 of Day 02
func (s *Solution) Part1(input string) string {
	ranges := aoc.CommaSeparated(aoc.TrimInput(input))
	total := 0
	buf := make([]byte, 0, 20)
	for _, r := range ranges {
		rangeStartStr, rangeEndStr, _ := strings.Cut(r, "-")
		rangeStart := aoc.MustInt(rangeStartStr)
		rangeEnd := aoc.MustInt(rangeEndStr)
		for id := rangeStart; id <= rangeEnd; id++ {
			buf := strconv.AppendInt(buf[:0], int64(id), 10)
			if !validateIdRepeatsTwice(string(buf)) {
				// fmt.Printf("%d\n", id)
				total += id
			}
		}
	}

	return strconv.Itoa(total)
}

func isRepeatedPattern(s string, patternLen int) bool {
	for ii := patternLen; ii < len(s); ii += 1 {
		if s[ii] != s[ii%patternLen] {
			return false
		}
	}
	return true
}

// ID is invalid if it is a sequence repeated at least twice like 1212 is invalid since it's 12 repeated twice, but 101 is valid
func validateIdRepeatsAtLeastTwice(s string) bool {
	idLength := len(s)
	for possibleSeqLen := 1; possibleSeqLen <= (idLength / 2); possibleSeqLen++ {
		if idLength%possibleSeqLen != 0 {
			continue
		}

		if isRepeatedPattern(s, possibleSeqLen) {
			return false
		}
	}
	return true
}

// Part2 solves part 2 of Day 02
func (s *Solution) Part2(input string) string {
	ranges := aoc.CommaSeparated(aoc.TrimInput(input))
	var total int64
	buf := make([]byte, 0, 20)
	for _, r := range ranges {
		rangeStartStr, rangeEndStr, _ := strings.Cut(r, "-")
		rangeStart := aoc.MustInt64(rangeStartStr)
		rangeEnd := aoc.MustInt64(rangeEndStr)
		for id := rangeStart; id <= rangeEnd; id++ {
			buf = strconv.AppendInt(buf[:0], id, 10)
			if !validateIdRepeatsAtLeastTwice(string(buf)) {
				// fmt.Printf("%d\n", id)
				total += id
			}
		}
	}

	return fmt.Sprintf("%d", total)
}
