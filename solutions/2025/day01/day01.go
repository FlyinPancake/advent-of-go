package day01

import (
	"fmt"
	"strconv"

	"github.com/flyinpancake/advent-of-go/aoc"
)

func init() {
	aoc.Register(2025, 1, &Solution{})
}

// Solution implements aoc.Solution for Year 2025 Day 01
type Solution struct{}

// Part1 solves part 1 of Day 01
func (s *Solution) Part1(input string) string {
	lines := aoc.Lines(input)
	rotations := append([]int{}, 50)
	for i, line := range lines {
		var data int
		if line[0] == 'L' {
			data = -1 * aoc.MustInt(line[1:])
		} else {
			data = aoc.MustInt(line[1:])
		}
		newRot := rotations[i] + data
		rotations = append(rotations, newRot)
	}

	n := 0
	for _, rot := range rotations {
		if rot%100 == 0 {
			n++
		}
	}

	return strconv.Itoa(n)
}

// Rotate dial from `current` to `dir` by `amount` returning (`newCurrent`, `passedZeroTimes`)
func rotateDial(current int, dir byte, amount int) (int, int) {
	sign := 1
	if dir == 'L' {
		sign = -1
	}
	changed := current + (sign * amount)
	if changed < 0 {
		overflows := changed / 100
		rem := aoc.PosMod(changed, 100)
		var correction int
		if current == 0 {
			correction = 0
		} else {
			correction = 1
		}
		return rem, aoc.Abs(overflows) + correction
	}
	if changed >= 100 {
		overflows := changed / 100
		rem := aoc.PosMod(changed, 100)
		return rem, overflows
	}
	if changed == 0 {
		return 0, 1
	}
	return changed, 0
}

// Part2 solves part 2 of Day 01
func (s *Solution) Part2(input string) string {
	lines := aoc.Lines(input)
	counter := 0
	dial := 50
	for _, line := range lines {
		dir := line[0]
		data := aoc.MustInt(line[1:])
		var count int
		dial, count = rotateDial(dial, dir, data)
		fmt.Println(line, ": ", dial, count)
		counter += count
	}

	// TODO: Implement Part 2
	return strconv.Itoa(counter)
}
