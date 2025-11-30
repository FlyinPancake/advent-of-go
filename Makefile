.PHONY: run test bench new help clean build fmt lint

# Default year
YEAR ?= 2025

# Help message
help:
	@echo "Advent of Code - Go Solutions"
	@echo ""
	@echo "Usage:"
	@echo "  make run DAY=1                  Run solution for day 1 (current year)"
	@echo "  make run DAY=1 YEAR=2023        Run solution for day 1, year 2023"
	@echo "  make run DAY=1 PART=1           Run only part 1 of day 1"
	@echo "  make test DAY=1                 Run tests for day 1"
	@echo "  make test YEAR=2023             Run all tests for year 2023"
	@echo "  make test                       Run all tests"
	@echo "  make bench DAY=1                Run benchmarks for day 1"
	@echo "  make new DAY=2                  Create template for day 2"
	@echo "  make new DAY=1 YEAR=2023        Create template for day 1, year 2023"
	@echo "  make list                       List all registered solutions"
	@echo "  make clean                      Remove build artifacts"
	@echo ""
	@echo "Options:"
	@echo "  DAY=N       Day number (1-25)"
	@echo "  PART=N      Part number (1 or 2)"
	@echo "  YEAR=YYYY   Year (default: 2025)"

# Run a solution
run:
ifndef DAY
	@echo "Error: DAY is required. Usage: make run DAY=1 [YEAR=2025]"
	@exit 1
endif
ifdef PART
	go run main.go -day=$(DAY) -part=$(PART) -year=$(YEAR)
else
	go run main.go -day=$(DAY) -year=$(YEAR)
endif

# List all registered solutions
list:
	go run main.go -list

# Run tests
test:
ifdef DAY
	go test -v ./solutions/$(YEAR)/day$(shell printf '%02d' $(DAY))/...
else ifdef YEAR
	go test -v ./solutions/$(YEAR)/...
else
	go test -v ./...
endif

# Run benchmarks
bench:
ifdef DAY
	go test -bench=. -benchmem ./solutions/$(YEAR)/day$(shell printf '%02d' $(DAY))/...
else ifdef YEAR
	go test -bench=. -benchmem ./solutions/$(YEAR)/...
else
	go test -bench=. -benchmem ./...
endif

# Create a new day from template
new:
ifndef DAY
	@echo "Error: DAY is required. Usage: make new DAY=2 [YEAR=2024]"
	@exit 1
endif
	@DAY_PADDED=$$(printf '%02d' $(DAY)); \
	if [ -d "solutions/$(YEAR)/day$$DAY_PADDED" ]; then \
		echo "Error: solutions/$(YEAR)/day$$DAY_PADDED already exists"; \
		exit 1; \
	fi; \
	mkdir -p solutions/$(YEAR)/day$$DAY_PADDED; \
	echo 'package day'$$DAY_PADDED > solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo 'import (' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '	"github.com/flyinpancake/advent-of-go/aoc"' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo ')' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo 'func init() {' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '	aoc.Register($(YEAR), $(DAY), &Solution{})' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '}' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '// Solution implements aoc.Solution for Year $(YEAR) Day '$$DAY_PADDED >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo 'type Solution struct{}' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '// Part1 solves part 1 of Day '$$DAY_PADDED >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo 'func (s *Solution) Part1(input string) string {' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '	lines := aoc.Lines(input)' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '	_ = lines // Use the parsed input' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '	// TODO: Implement Part 1' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '	return "not implemented"' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '}' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '// Part2 solves part 2 of Day '$$DAY_PADDED >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo 'func (s *Solution) Part2(input string) string {' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '	lines := aoc.Lines(input)' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '	_ = lines // Use the parsed input' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '	// TODO: Implement Part 2' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '	return "not implemented"' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo '}' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED.go; \
	echo 'package day'$$DAY_PADDED > solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo 'import (' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	"testing"' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	"github.com/flyinpancake/advent-of-go/aoc"' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo ')' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo 'var exampleInput = `your example input here`' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo 'func TestPart1(t *testing.T) {' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	testCases := []aoc.TestCase{' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '		{' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '			Name:     "example",' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '			Input:    exampleInput,' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '			Expected: "expected_result",' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '		},' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	}' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	aoc.RunPart1Tests(t, &Solution{}, testCases)' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '}' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo 'func TestPart2(t *testing.T) {' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	testCases := []aoc.TestCase{' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '		{' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '			Name:     "example",' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '			Input:    exampleInput,' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '			Expected: "expected_result",' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '		},' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	}' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	aoc.RunPart2Tests(t, &Solution{}, testCases)' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '}' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo 'func BenchmarkPart1(b *testing.B) {' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	input, err := aoc.ReadInput($(YEAR), $(DAY))' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	if err != nil {' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '		b.Skip("input not available")' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	}' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	aoc.BenchmarkPart1(b, &Solution{}, input)' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '}' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo 'func BenchmarkPart2(b *testing.B) {' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	input, err := aoc.ReadInput($(YEAR), $(DAY))' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	if err != nil {' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '		b.Skip("input not available")' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	}' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '	aoc.BenchmarkPart2(b, &Solution{}, input)' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	echo '}' >> solutions/$(YEAR)/day$$DAY_PADDED/day$$DAY_PADDED\_test.go; \
	mkdir -p inputs/$(YEAR); \
	touch inputs/$(YEAR)/day$$DAY_PADDED.txt; \
	echo "Created solutions/$(YEAR)/day$$DAY_PADDED/"; \
	echo "Created inputs/$(YEAR)/day$$DAY_PADDED.txt"; \
	echo ""; \
	echo "Don't forget to add the import to main.go:"; \
	echo '  _ "github.com/flyinpancake/advent-of-go/solutions/$(YEAR)/day'$$DAY_PADDED'"'

# Clean build artifacts
clean:
	go clean
	rm -f advent-of-go

# Build the binary
build:
	go build -o advent-of-go main.go

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	go vet ./...
