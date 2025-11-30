package aoc

import (
	"fmt"
	"sort"
	"sync"
)

// Solution defines the interface that each day's solution must implement.
type Solution interface {
	// Part1 solves part 1 of the puzzle and returns the answer as a string.
	Part1(input string) string
	// Part2 solves part 2 of the puzzle and returns the answer as a string.
	Part2(input string) string
}

// SolutionKey uniquely identifies a solution by year and day
type SolutionKey struct {
	Year int
	Day  int
}

func (k SolutionKey) String() string {
	return fmt.Sprintf("%d/day%02d", k.Year, k.Day)
}

var (
	registry = make(map[SolutionKey]Solution)
	mu       sync.RWMutex
)

// Register adds a solution to the registry for a specific year and day.
// This is typically called from an init() function in each day's package.
func Register(year, day int, solution Solution) {
	mu.Lock()
	defer mu.Unlock()
	key := SolutionKey{Year: year, Day: day}
	if _, exists := registry[key]; exists {
		panic(fmt.Sprintf("solution already registered for %s", key))
	}
	registry[key] = solution
}

// GetSolution retrieves a solution for a specific year and day.
// Returns the solution and true if found, nil and false otherwise.
func GetSolution(year, day int) (Solution, bool) {
	mu.RLock()
	defer mu.RUnlock()
	key := SolutionKey{Year: year, Day: day}
	solution, ok := registry[key]
	return solution, ok
}

// RegisteredSolutions returns a sorted slice of all registered solution keys.
func RegisteredSolutions() []SolutionKey {
	mu.RLock()
	defer mu.RUnlock()
	keys := make([]SolutionKey, 0, len(registry))
	for key := range registry {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].Year != keys[j].Year {
			return keys[i].Year < keys[j].Year
		}
		return keys[i].Day < keys[j].Day
	})
	return keys
}

// RegisteredYears returns a sorted slice of all years that have registered solutions.
func RegisteredYears() []int {
	mu.RLock()
	defer mu.RUnlock()
	yearSet := make(map[int]struct{})
	for key := range registry {
		yearSet[key.Year] = struct{}{}
	}
	years := make([]int, 0, len(yearSet))
	for year := range yearSet {
		years = append(years, year)
	}
	sort.Ints(years)
	return years
}

// RegisteredDays returns a sorted slice of all registered day numbers for a specific year.
func RegisteredDays(year int) []int {
	mu.RLock()
	defer mu.RUnlock()
	days := make([]int, 0)
	for key := range registry {
		if key.Year == year {
			days = append(days, key.Day)
		}
	}
	sort.Ints(days)
	return days
}
