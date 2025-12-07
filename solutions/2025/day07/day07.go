package day07

import (
	"strconv"

	"github.com/flyinpancake/advent-of-go/aoc"
)

func init() {
	aoc.Register(2025, 7, &Solution{})
}

// Solution implements aoc.Solution for Year 2024 Day 02
type Solution struct{}

// Part1 solves part 1 of Day 02
func (s *Solution) Part1(input string) string {
	grid := aoc.Grid(input)

	start := aoc.Point{X: 0, Y: 0}

	for i := range grid[0] {
		if grid[0][i] == 'S' {
			start.X = i
			break
		}
	}

	q := make(aoc.Queue[aoc.Point], 0)
	visited := make(aoc.Set[aoc.Point])

	q.Enqueue(start)

	gridH := len(grid)
	gridW := len(grid[0])

	splits := 0

	for !q.IsEmpty() {
		element := q.Dequeue()
		if element.X >= gridW || element.X < 0 || element.Y >= gridH || visited.Contains(element) {
			continue
		}
		visited.Add(element)
		gridTile := grid[element.Y][element.X]

		switch gridTile {
		case '.', 'S':
			q.Enqueue(element.Add(aoc.Down))

		case '^':
			splits++
			q.Enqueue(element.Add(aoc.Down).Add(aoc.Left))
			q.Enqueue(element.Add(aoc.Down).Add(aoc.Right))
		}
	}

	return strconv.Itoa(splits)
}

// Part2 solves part 2 of Day 02
func (s *Solution) Part2(input string) string {
	grid := aoc.Grid(input)

	start := aoc.Point{X: 0, Y: 0}

	for i := range grid[0] {
		if grid[0][i] == 'S' {
			start.X = i
			break
		}
	}

	q := make(aoc.Queue[aoc.Point], 0)
	q.Enqueue(start)

	possibleRoutesTo := make(map[aoc.Point]int)
	possibleRoutesTo[start] = 1

	visited := make(aoc.Set[aoc.Point])

	gridH := len(grid)
	gridW := len(grid[0])

	for !q.IsEmpty() {
		element := q.Dequeue()
		if element.X >= gridW || element.X < 0 || element.Y >= gridH || visited.Contains(element) {
			continue
		}
		visited.Add(element)
		possibleRoutes := possibleRoutesTo[element]

		gridTile := grid[element.Y][element.X]

		switch gridTile {
		case '.', 'S':
			q.Enqueue(element.Add(aoc.Down))
			possibleRoutesTo[element.Add(aoc.Down)] += possibleRoutes

		case '^':
			q.Enqueue(element.Add(aoc.Down).Add(aoc.Left))
			possibleRoutesTo[element.Add(aoc.Down).Add(aoc.Left)] += possibleRoutes

			q.Enqueue(element.Add(aoc.Down).Add(aoc.Right))
			possibleRoutesTo[element.Add(aoc.Down).Add(aoc.Right)] += possibleRoutes

		}
	}

	timelines := 0
	for x := 0; x < gridW; x++ {
		timelines += possibleRoutesTo[aoc.Point{X: x, Y: gridH - 1}]
	}

	return strconv.Itoa(timelines)
}
