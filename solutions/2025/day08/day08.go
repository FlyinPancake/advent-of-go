package day08

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"

	"github.com/flyinpancake/advent-of-go/aoc"
)

func init() {
	aoc.Register(2025, 8, &Solution{})
}

// Solution implements aoc.Solution for Year 2024 Day 02
type Solution struct{}

const alreadyConnected float64 = -1
const shortestSize int = 10
const targetConnections int = 10

// Part1 solves part 1 of Day 02
func (s *Solution) Part1(input string) string {
	lines := aoc.Lines(input)
	junctionBoxes := []aoc.Point3{}
	for _, line := range lines {
		fields := aoc.CommaSeparatedInts(line)
		junctionBoxes = append(junctionBoxes, aoc.Point3{X: fields[0], Y: fields[1], Z: fields[2]})
	}

	junctionBoxCount := len(junctionBoxes)

	connections := make([][]float64, junctionBoxCount)
	for i := range connections {
		connections[i] = make([]float64, junctionBoxCount)
	}

	for fromJunctionId, fromJunction := range junctionBoxes {
		for toJunctionId, toJunction := range junctionBoxes[fromJunctionId:] {
			dist := fromJunction.Eucledian(toJunction)
			connections[fromJunctionId][fromJunctionId+toJunctionId] = dist
		}
	}

	// for _, r := range connections {
	// 	for _, c := range r {
	// 		fmt.Printf("%06.0f|", c)
	// 	}
	// 	fmt.Print("\n")
	// }

	type DistanceWithIds struct {
		from int
		to   int
		val  float64
	}

	circuits := []aoc.Set[int]{}

	connectionsMade := 0

	for connectionsMade < targetConnections {
		// find the next two junction boxes with minimal distance
		min := DistanceWithIds{val: math.MaxFloat64}

		for fromId, r := range connections {
			for toIdOffset, c := range r[fromId:] {
				if (c != float64(0) && c != alreadyConnected) && c < min.val {
					toId := fromId + toIdOffset
					min = DistanceWithIds{
						to:   toId,
						from: fromId,
						val:  c,
					}
				}
			}
		}

		connections[min.from][min.to] = alreadyConnected

		// check if this pair of junction boxes is already connected to a circuit
		// (transient connections count as well like 'to' is connected to 'a' and 'a' is connected to 'from')
		transientlyConnected := false
		for _, cir := range circuits {
			if cir.Contains(min.from) && cir.Contains(min.to) {
				transientlyConnected = true
				break
			}
		}

		if transientlyConnected {
			continue
		}

		connectionsMade++

		// -1 means not in any circuit
		fromCirId := -1
		toCirId := -1

		for cirId, cir := range circuits {
			if cir.Contains(min.from) {
				fromCirId = cirId
			}
			if cir.Contains(min.to) {
				toCirId = cirId
			}
		}

		if fromCirId >= 0 && toCirId >= 0 {
			toCir := aoc.NewSet(circuits[toCirId].ToSlice()...)
			for _, e := range toCir.ToSlice() {
				circuits[fromCirId].Add(e)
			}
			circuits = slices.Delete(circuits, toCirId, toCirId+1)
		} else if fromCirId >= 0 {
			circuits[fromCirId].Add(min.to)
		} else if toCirId >= 0 {
			circuits[toCirId].Add(min.from)
		} else {
			circuits = append(circuits, aoc.NewSet(min.from, min.to))
		}
	}

	fmt.Println("Circuits:", circuits)

	sizes := make([]int, len(circuits))
	for i, cir := range circuits {
		sizes[i] = cir.Len()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	product := 1

	for _, v := range sizes[:3] {
		product *= v
	}

	return strconv.Itoa(product)
}

// Part2 solves part 2 of Day 02
func (s *Solution) Part2(input string) string {
	lines := aoc.Lines(input)
	_ = lines // Use the parsed input

	// TODO: Implement Part 2
	return "not implemented"
}
