# Advent of Code - Go Solutions

A Go harness for solving Advent of Code puzzles with testing, benchmarking, and utilities. Supports multiple years!

## Project Structure

```
advent-of-go/
├── inputs/              # Input files organized by year
│   ├── 2024/
│   │   ├── day01.txt
│   │   └── ...
│   └── 2023/
│       └── ...
└── solutions/           # Solutions organized by year
    ├── 2024/
    │   ├── day01/
    │   │   ├── day01.go
    │   │   └── day01_test.go
    │   └── day02/
    │       └── ...
    └── 2023/
        └── ...
```

## Quick Start

### 1. Create a New Day

```bash
# Create day 1 for 2025 (default year)
make new DAY=1

# Create day 1 for a different year
make new DAY=1 YEAR=2024
```

### 2. Add Your Input

Place your puzzle input at `inputs/2024/day01.txt` (or the appropriate year/day).

### 3. Implement Your Solution

Edit `solutions/2024/day01/day01.go`:

```go
func (s *Solution) Part1(input string) string {
    lines := aoc.Lines(input)
    // Your solution here
    return fmt.Sprintf("%d", result)
}
```

### 4. Register Your Solution

In `main.go`, add the import for your day:

```go
import (
    // ...
    // 2024
    _ "github.com/flyinpancake/advent-of-go/solutions/2024/day01"

    // 2023
    _ "github.com/flyinpancake/advent-of-go/solutions/2023/day01"
)
```

### 5. Run Your Solution

```bash
# Run a specific day (uses default year 2025)
make run DAY=1

# Run a specific year and day
make run DAY=1 YEAR=2023

# Run a specific part
make run DAY=1 PART=1

# List all registered solutions
make list

# Show help
go run main.go
```

## Commands

| Command                            | Description                   |
| ---------------------------------- | ----------------------------- |
| `make run DAY=N [YEAR=Y] [PART=P]` | Run solution for day N        |
| `make test DAY=N [YEAR=Y]`         | Run tests for day N           |
| `make test YEAR=Y`                 | Run all tests for year Y      |
| `make test`                        | Run all tests                 |
| `make bench DAY=N [YEAR=Y]`        | Run benchmarks for day N      |
| `make new DAY=N [YEAR=Y]`          | Create template for day N     |
| `make list`                        | List all registered solutions |
| `make clean`                       | Remove build artifacts        |

## Testing

```bash
# Run tests for a specific day
make test DAY=1

# Run tests for a specific year
make test YEAR=2024

# Run all tests
make test

# Run benchmarks
make bench DAY=1
```

## Available Utilities

### Parsing (`aoc` package)

| Function                | Description                            |
| ----------------------- | -------------------------------------- |
| `Lines(input)`          | Split input into lines                 |
| `Paragraphs(input)`     | Split by blank lines                   |
| `Ints(line)`            | Extract all integers from a string     |
| `IntFields(line)`       | Split by whitespace, parse as ints     |
| `Grid(input)`           | Parse into 2D rune grid                |
| `IntGrid(input)`        | Parse into 2D int grid (single digits) |
| `MustInt(s)`            | Parse string to int (panics on error)  |
| `CommaSeparatedInts(s)` | Parse comma-separated integers         |
| `SplitTwo(s, sep)`      | Split into exactly two parts           |

### Math (`aoc` package)

| Function                        | Description                                     |
| ------------------------------- | ----------------------------------------------- |
| `Abs(x)`                        | Absolute value                                  |
| `Min(a, b)` / `Max(a, b)`       | Min/max of two values                           |
| `GCD(a, b)` / `LCM(a, b)`       | Greatest common divisor / Least common multiple |
| `Sum(slice)` / `Product(slice)` | Sum/product of slice                            |
| `Pow(a, b)` / `ModPow(a, b, m)` | Integer power / modular exponentiation          |

### Data Structures (`aoc` package)

| Type              | Description                                 |
| ----------------- | ------------------------------------------- |
| `Point{X, Y}`     | 2D point with neighbors, Manhattan distance |
| `Point3{X, Y, Z}` | 3D point                                    |
| `Set[T]`          | Generic set with union/intersection         |
| `Counter[T]`      | Count occurrences                           |
| `Stack[T]`        | LIFO stack                                  |
| `Queue[T]`        | FIFO queue                                  |

### Directions

```go
aoc.Up, aoc.Down, aoc.Left, aoc.Right  // Direction vectors
aoc.Cardinals                           // []Point{Up, Right, Down, Left}
aoc.AllDirections                       // All 8 directions including diagonals
point.Neighbors4()                      // 4 cardinal neighbors
point.Neighbors8()                      // 8 neighbors including diagonals
```

## Example Solution

```go
package day01

import (
    "fmt"
    "github.com/flyinpancake/advent-of-go/aoc"
)

func init() {
    aoc.Register(2024, 1, &Solution{})
}

type Solution struct{}

func (s *Solution) Part1(input string) string {
    lines := aoc.Lines(input)
    total := 0
    for _, line := range lines {
        nums := aoc.Ints(line)
        total += aoc.Sum(nums)
    }
    return fmt.Sprintf("%d", total)
}

func (s *Solution) Part2(input string) string {
    // Part 2 solution
    return "not implemented"
}
```

## Working with Multiple Years

The harness is designed to easily switch between years:

```bash
# Work on 2024
make new DAY=1 YEAR=2024
make run DAY=1 YEAR=2024

# Work on 2023
make new DAY=1 YEAR=2023
make run DAY=1 YEAR=2023

# Run all tests across all years
make test
```

Each year's solutions are completely isolated in their own directory, making it easy to maintain and compare solutions across years.

## Tips

- Use `aoc.Lines()` for line-by-line parsing
- Use `aoc.Ints()` to extract all numbers from messy input
- Use `aoc.Grid()` for 2D puzzles
- Use `Point` with `Neighbors4()` for grid traversal
- Use `Set[Point]` for visited tracking
- Put example inputs in `inputs/2024/day01_example.txt`
- Write tests with the example input first before running on real input
- The default year is 2025; override with `YEAR=XXXX`
