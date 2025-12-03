package day03

import (
	"strconv"

	"github.com/flyinpancake/advent-of-go/aoc"
)

func init() {
	aoc.Register(2025, 3, &Solution{})
}

// Solution implements aoc.Solution for Year 2024 Day 02
type Solution struct{}

// largestTwo returns the largest two ints in the string as a concatenated number
func largestTwo(s string) int {

	chars := aoc.Chars(s)
	len := len(chars)
	// nums is pos,val
	nums := []int{}
	for _, c := range chars {
		i, _ := strconv.Atoi(string(c))
		nums = append(nums, i)
	}

	largestFirst := 0
	largestSecond := 1

	for ii, val := range nums {
		if (ii < len-1) && (nums[largestFirst] < val) {
			largestFirst = ii
			largestSecond = ii + 1
		} else if ii != 0 && nums[largestSecond] < val {
			largestSecond = ii
		}
	}

	return nums[largestFirst]*10 + nums[largestSecond]
}

// Part1 solves part 1 of Day 02
func (s *Solution) Part1(input string) string {
	lines := aoc.Lines(input)

	total := 0
	for _, line := range lines {
		l2 := largestTwo(line)
		// fmt.Printf("%s: %d\n", line, l2)
		total += l2
	}

	return strconv.Itoa(total)
}

func largestN(s string, n int) int {
	chars := aoc.Chars(s)
	len := len(chars)
	if len < n {
		panic("n can't be smaller than len")
	}
	// nums is pos,val
	nums := []int{}
	for _, c := range chars {
		i, _ := strconv.Atoi(string(c))
		nums = append(nums, i)
	}

	largest := []int{}

	for ii := 0; ii < n; ii++ {
		largest = append(largest, ii)
	}

	for numIdx := 0; numIdx <= len-n; numIdx++ {
		for offsetFromFirst := 0; offsetFromFirst < n; offsetFromFirst++ {
			if nums[numIdx+offsetFromFirst] > nums[largest[offsetFromFirst]] {
				for offsetFromCurrentLargest := 0; offsetFromCurrentLargest < n-offsetFromFirst; offsetFromCurrentLargest++ {
					largest[offsetFromFirst+offsetFromCurrentLargest] = numIdx + offsetFromFirst + offsetFromCurrentLargest
				}
			}
		}
	}
	total := 0
	for ii, numIdx := range largest {
		total += aoc.Pow(10, n-ii-1) * nums[numIdx]
	}

	return total
}

// Part2 solves part 2 of Day 02
func (s *Solution) Part2(input string) string {
	lines := aoc.Lines(input)
	_ = lines // Use the parsed input

	total := 0
	for _, line := range lines {
		l2 := largestN(line, 12)
		// fmt.Printf("%s: %d\n", line, l2)
		total += l2
	}

	return strconv.Itoa(total)
}
