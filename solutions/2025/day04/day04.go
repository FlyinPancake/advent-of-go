package day04

import (
	"strconv"

	"github.com/flyinpancake/advent-of-go/aoc"
)

func init() {
	aoc.Register(2025, 4, &Solution{})
}

// Solution implements aoc.Solution for Year 2024 Day 02
type Solution struct{}

// Part1 solves part 1 of Day 02
func (s *Solution) Part1(input string) string {
	grid := aoc.Grid(aoc.TrimInput(input))

	forkliftCertified := 0

	height := len(grid)
	width := len(grid[0])
	for y, line := range grid {
		for x, val := range line {
			p := aoc.Point{X: x, Y: y}
			neighbors := p.Neighbors8()
			if val != '@' {
				continue
			}
			takenNeighborSpots := 0
			for _, neighbor := range neighbors {
				if neighbor.Y >= 0 && neighbor.Y < height && neighbor.X >= 0 && neighbor.X < width && grid[neighbor.Y][neighbor.X] == '@' {
					takenNeighborSpots++
				}
			}
			if takenNeighborSpots < 4 {
				forkliftCertified++
			}
		}
	}

	return strconv.Itoa(forkliftCertified)
}

// Part2 solves part 2 of Day 02
func (s *Solution) Part2(input string) string {
	grid := aoc.Grid(aoc.TrimInput(input))

	removed := 0

	for {
		forkliftCertified := []aoc.Point{}

		height := len(grid)
		width := len(grid[0])
		for y, line := range grid {
			for x, val := range line {
				p := aoc.Point{X: x, Y: y}
				neighbors := p.Neighbors8()
				if val != '@' {
					continue
				}
				takenNeighborSpots := 0
				for _, neighbor := range neighbors {
					if neighbor.Y >= 0 && neighbor.Y < height && neighbor.X >= 0 && neighbor.X < width && grid[neighbor.Y][neighbor.X] == '@' {
						takenNeighborSpots++
					}
				}
				if takenNeighborSpots < 4 {
					forkliftCertified = append(forkliftCertified, p)
				}
			}
		}
		if len(forkliftCertified) == 0 {
			break
		}

		for _, p := range forkliftCertified {
			grid[p.Y][p.X] = '.'
			removed++
		}
	}
	return strconv.Itoa(removed)
}
