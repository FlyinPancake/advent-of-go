package day01

import (
	"strconv"

	"github.com/flyinpancake/advent-of-go/aoc"
)

func init() {
	aoc.Register(2025, 1, &Solution{})
}

// Solution implements aoc.Solution for Year 2025 Day 01
type Solution struct{}

const (
	dialSize     = 100
	dialStartRot = 50
)

// Part1 solves part 1 of Day 01
func (*Solution) Part1(input string) string {
	lines := aoc.Lines(input)
	currentRot := dialStartRot
	landedOnZero := 0
	for _, line := range lines {
		data := aoc.MustInt(line[1:])
		if line[0] == 'L' {
			data = -data
		}
		currentRot += data
		if currentRot%dialSize == 0 {
			landedOnZero++
		}
	}

	return strconv.Itoa(landedOnZero)
}

// rotateDial simulates rotating the dial from the current position
// in dir ('L' or 'R') by amount, returning new value (0<=value<100)
// and the number of times the 0 mark was passed
func rotateDial(current int, dir byte, amount int) (newRot, passedZero int) {
	sign := 1
	if dir == 'L' {
		sign = -1
	}
	changed := current + (sign * amount)
	if changed < 0 {
		overflows := changed / dialSize
		newRot = aoc.PosMod(changed, dialSize)
		passedZero = aoc.Abs(overflows) + 1
		// Correct for starting the rotation from 0
		if current == 0 {
			passedZero--
		}
		return
	}
	if changed >= dialSize {
		passedZero = changed / dialSize
		newRot = aoc.PosMod(changed, dialSize)
		return
	}
	if changed == 0 {
		return 0, 1
	}
	return changed, 0
}

// Part2 solves part 2 of Day 01
func (*Solution) Part2(input string) string {
	lines := aoc.Lines(input)
	counter := 0
	dial := dialStartRot
	for _, line := range lines {
		dir := line[0]
		data := aoc.MustInt(line[1:])
		var count int
		dial, count = rotateDial(dial, dir, data)
		counter += count
	}

	return strconv.Itoa(counter)
}
