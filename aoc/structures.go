package aoc

import "math"

// Point represents a 2D point with integer coordinates
type Point struct {
	X, Y int
}

// Add returns a new point that is the sum of p and other
func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

// Sub returns a new point that is the difference of p and other
func (p Point) Sub(other Point) Point {
	return Point{p.X - other.X, p.Y - other.Y}
}

// Scale returns a new point with coordinates scaled by factor
func (p Point) Scale(factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// Manhattan returns the Manhattan distance from p to other
func (p Point) Manhattan(other Point) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

// Neighbors4 returns the 4 cardinal neighbors (up, down, left, right)
func (p Point) Neighbors4() []Point {
	return []Point{
		p.Add(Up),
		p.Add(Left),
		p.Add(Right),
		p.Add(Left),
	}
}

// Neighbors8 returns all 8 neighbors (including diagonals)
func (p Point) Neighbors8() []Point {
	return []Point{
		p.Add(Up),
		p.Add(Up).Add(Right),
		p.Add(Right),
		p.Add(Right).Add(Down),
		p.Add(Down),
		p.Add(Down).Add(Left),
		p.Add(Left),
		p.Add(Left).Add(Up),
	}
}

// Common direction constants
var (
	Up    = Point{0, -1}
	Down  = Point{0, 1}
	Left  = Point{-1, 0}
	Right = Point{1, 0}

	// Cardinal directions in order: Up, Right, Down, Left
	Cardinals = []Point{Up, Right, Down, Left}

	// All 8 directions
	AllDirections = []Point{
		{0, -1}, {1, -1}, {1, 0}, {1, 1},
		{0, 1}, {-1, 1}, {-1, 0}, {-1, -1},
	}
)

// Point3 represents a 3D point with integer coordinates
type Point3 struct {
	X, Y, Z int
}

// Add returns a new point that is the sum of p and other
func (p Point3) Add(other Point3) Point3 {
	return Point3{p.X + other.X, p.Y + other.Y, p.Z + other.Z}
}

// Sub returns a new point that is the difference of p and other
func (p Point3) Sub(other Point3) Point3 {
	return Point3{p.X - other.X, p.Y - other.Y, p.Z - other.Z}
}

// Manhattan returns the Manhattan distance from p to other
func (p Point3) Manhattan(other Point3) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y) + Abs(p.Z-other.Z)
}

// Eucledian returns the eucledian distance from p to other
func (p Point3) Eucledian(other Point3) float64 {
	return math.Sqrt(
		math.Pow(float64(p.X-other.X), 2) +
			math.Pow(float64(p.Y-other.Y), 2) +
			math.Pow(float64(p.Z-other.Z), 2),
	)
}

// Neighbors6 returns the 6 face-adjacent neighbors in 3D
func (p Point3) Neighbors6() []Point3 {
	return []Point3{
		{p.X + 1, p.Y, p.Z}, {p.X - 1, p.Y, p.Z},
		{p.X, p.Y + 1, p.Z}, {p.X, p.Y - 1, p.Z},
		{p.X, p.Y, p.Z + 1}, {p.X, p.Y, p.Z - 1},
	}
}

// Set is a generic set implementation using a map
type Set[T comparable] map[T]struct{}

// NewSet creates a new set with optional initial elements
func NewSet[T comparable](elements ...T) Set[T] {
	s := make(Set[T])
	for _, e := range elements {
		s.Add(e)
	}
	return s
}

// Add adds an element to the set
func (s Set[T]) Add(element T) {
	s[element] = struct{}{}
}

// Remove removes an element from the set
func (s Set[T]) Remove(element T) {
	delete(s, element)
}

// Contains returns true if the element is in the set
func (s Set[T]) Contains(element T) bool {
	_, ok := s[element]
	return ok
}

// Len returns the number of elements in the set
func (s Set[T]) Len() int {
	return len(s)
}

// ToSlice returns all elements as a slice
func (s Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s))
	for e := range s {
		result = append(result, e)
	}
	return result
}

// Union returns a new set with elements from both sets
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := NewSet[T](s.ToSlice()...)
	for e := range other {
		result.Add(e)
	}
	return result
}

// Intersection returns a new set with elements in both sets
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := NewSet[T]()
	for e := range s {
		if other.Contains(e) {
			result.Add(e)
		}
	}
	return result
}

// Counter counts occurrences of elements
type Counter[T comparable] map[T]int

// NewCounter creates a new counter with optional initial elements
func NewCounter[T comparable](elements ...T) Counter[T] {
	c := make(Counter[T])
	for _, e := range elements {
		c.Add(e)
	}
	return c
}

// Add increments the count for an element
func (c Counter[T]) Add(element T) {
	c[element]++
}

// AddN adds n to the count for an element
func (c Counter[T]) AddN(element T, n int) {
	c[element] += n
}

// Get returns the count for an element (0 if not present)
func (c Counter[T]) Get(element T) int {
	return c[element]
}

// MostCommon returns the element with the highest count
func (c Counter[T]) MostCommon() (T, int) {
	var maxElement T
	maxCount := 0
	first := true
	for e, count := range c {
		if first || count > maxCount {
			maxElement = e
			maxCount = count
			first = false
		}
	}
	return maxElement, maxCount
}

// LeastCommon returns the element with the lowest count
func (c Counter[T]) LeastCommon() (T, int) {
	var minElement T
	minCount := 0
	first := true
	for e, count := range c {
		if first || count < minCount {
			minElement = e
			minCount = count
			first = false
		}
	}
	return minElement, minCount
}

// Stack is a generic LIFO stack
type Stack[T any] []T

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(element T) {
	*s = append(*s, element)
}

// Pop removes and returns the top element, panics if empty
func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		panic("pop from empty stack")
	}
	n := len(*s) - 1
	element := (*s)[n]
	*s = (*s)[:n]
	return element
}

// Peek returns the top element without removing it, panics if empty
func (s *Stack[T]) Peek() T {
	if len(*s) == 0 {
		panic("peek from empty stack")
	}
	return (*s)[len(*s)-1]
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Len returns the number of elements in the stack
func (s *Stack[T]) Len() int {
	return len(*s)
}

// Queue is a generic FIFO queue
type Queue[T any] []T

// Enqueue adds an element to the back of the queue
func (q *Queue[T]) Enqueue(element T) {
	*q = append(*q, element)
}

// Dequeue removes and returns the front element, panics if empty
func (q *Queue[T]) Dequeue() T {
	if len(*q) == 0 {
		panic("dequeue from empty queue")
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element
}

// Front returns the front element without removing it, panics if empty
func (q *Queue[T]) Front() T {
	if len(*q) == 0 {
		panic("front from empty queue")
	}
	return (*q)[0]
}

// IsEmpty returns true if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

// Len returns the number of elements in the queue
func (q *Queue[T]) Len() int {
	return len(*q)
}

// Range represents a range of numeric values [Start, End)
// By default, the range is exclusive of the end value
type Range struct {
	Start     int
	End       int
	Inclusive bool // If true, the range includes End
}

func NewRange(upper, lower int) Range {
	return Range{
		Start: lower,
		End:   upper,
	}
}

// AsInclusive returns the range as inclusive [Start, End].
func (r Range) AsInclusive() Range {
	r.Inclusive = true
	return r
}

// AsExclusive returns the range as exclusive [Start, End).
func (r Range) AsExclusive() Range {
	r.Inclusive = false
	return r
}

func NewInclusiveRange(start, end int) Range {
	return Range{
		Start: start,
		End:   end,

		Inclusive: true,
	}
}

// Contains checks if a value is within the range.
func (r Range) Contains(value int) bool {
	if r.Inclusive {
		return value >= r.Start && value <= r.End
	}
	return value >= r.Start && value < r.End
}

// IsEmpty returns true if the range contains no elements.
func (r Range) IsEmpty() bool {
	if r.Inclusive {
		return r.Start > r.End
	}
	return r.Start >= r.End
}

// Len returns the number of elements in the range.
func (r Range) Len() int {
	if r.IsEmpty() {
		return 0
	}

	count := r.End - r.Start

	if r.Inclusive {
		count++
	}

	return count
}

func EmptyRange() Range {
	return NewRange(0, 0)
}
