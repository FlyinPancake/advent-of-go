package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/flyinpancake/advent-of-go/aoc"
	// Import all day solutions here, organized by year
	// 2025
	_ "github.com/flyinpancake/advent-of-go/solutions/2025/day01"
	_ "github.com/flyinpancake/advent-of-go/solutions/2025/day02"
	_ "github.com/flyinpancake/advent-of-go/solutions/2025/day03"
	_ "github.com/flyinpancake/advent-of-go/solutions/2025/day04"
	_ "github.com/flyinpancake/advent-of-go/solutions/2025/day05"
	_ "github.com/flyinpancake/advent-of-go/solutions/2025/day06"
	_ "github.com/flyinpancake/advent-of-go/solutions/2025/day07"

	// 2024
	_ "github.com/flyinpancake/advent-of-go/solutions/2024/day01"
	// _ "github.com/flyinpancake/advent-of-go/solutions/2024/day02"
	// 2023
	// _ "github.com/flyinpancake/advent-of-go/solutions/2023/day01"
)

func main() {
	day := flag.Int("day", 0, "Day to run (1-25)")
	part := flag.Int("part", 0, "Part to run (1 or 2, 0 for both)")
	year := flag.Int("year", 2025, "Year of Advent of Code")
	benchmark := flag.Bool("bench", false, "Run benchmarks")
	iterations := flag.Int("iterations", 1000, "Number of benchmark iterations")
	listAll := flag.Bool("list", false, "List all registered solutions")
	new := flag.Bool("new", false, "Scaffold a new day's solution (requires -day and -year)")
	flag.Parse()

	if *listAll {
		printRegisteredSolutions()
		os.Exit(0)
	}

	if *day == 0 {
		flag.Usage()
		os.Exit(0)
	}

	if *day < 1 || *day > 25 {
		fmt.Fprintf(os.Stderr, "Error: day must be between 1 and 25\n")
		os.Exit(1)
	}

	if *part < 0 || *part > 2 {
		fmt.Fprintf(os.Stderr, "Error: part must be 0, 1, or 2\n")
		os.Exit(1)
	}

	if *new {
		err := scaffoldNewDay(*year, *day)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scaffolding new day: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Scaffolded new solution for Year %d Day %02d\n", *year, *day)
		fmt.Printf("Don't forget to import the new day's package in main.go!\n")
		fmt.Printf("      _ \"github.com/flyinpancake/advent-of-go/solutions/%d/day%02d\"\n", *year, *day)
		os.Exit(0)
	}

	solution, ok := aoc.GetSolution(*year, *day)
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: no solution registered for year %d day %d\n", *year, *day)
		fmt.Fprintf(os.Stderr, "Hint: make sure to import the day's package in main.go\n")
		fmt.Fprintf(os.Stderr, "      _ \"github.com/flyinpancake/advent-of-go/solutions/%d/day%02d\"\n", *year, *day)
		os.Exit(1)
	}

	input, err := aoc.ReadInput(*year, *day)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		fmt.Fprintf(os.Stderr, "Hint: place your input at inputs/%d/day%02d.txt\n", *year, *day)
		os.Exit(1)
	}

	fmt.Printf("=== Year %d Day %02d ===\n", *year, *day)

	if *benchmark {
		runBenchmark(solution, input, *part, *iterations)
	} else {
		runSolution(solution, input, *part)
	}
}

func printRegisteredSolutions() {
	fmt.Println("Registered solutions:")
	years := aoc.RegisteredYears()
	if len(years) == 0 {
		fmt.Println("  (none registered yet)")
		return
	}
	for _, year := range years {
		days := aoc.RegisteredDays(year)
		fmt.Printf("  %d: ", year)
		for i, day := range days {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("Day %02d", day)
		}
		fmt.Println()
	}
}

func runSolution(solution aoc.Solution, input string, part int) {
	if part == 0 || part == 1 {
		start := time.Now()
		result := solution.Part1(input)
		elapsed := time.Since(start)
		fmt.Printf("Part 1: %s (took %v)\n", result, elapsed)
	}

	if part == 0 || part == 2 {
		start := time.Now()
		result := solution.Part2(input)
		elapsed := time.Since(start)
		fmt.Printf("Part 2: %s (took %v)\n", result, elapsed)
	}
}

func runBenchmark(solution aoc.Solution, input string, part int, iterations int) {
	fmt.Printf("Running %d iterations...\n", iterations)

	if part == 0 || part == 1 {
		start := time.Now()
		for i := 0; i < iterations; i++ {
			solution.Part1(input)
		}
		elapsed := time.Since(start)
		avg := elapsed / time.Duration(iterations)
		fmt.Printf("Part 1: %v avg (%v total for %d iterations)\n", avg, elapsed, iterations)
	}

	if part == 0 || part == 2 {
		start := time.Now()
		for i := 0; i < iterations; i++ {
			solution.Part2(input)
		}
		elapsed := time.Since(start)
		avg := elapsed / time.Duration(iterations)
		fmt.Printf("Part 2: %v avg (%v total for %d iterations)\n", avg, elapsed, iterations)
	}
}

func scaffoldNewDay(year, day int) error {
	return aoc.CreateNewSolutionScaffold(year, day)
}
