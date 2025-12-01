package aoc

// Abs returns the absolute value of x
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// AbsDiff return the absolute difference between a and b
func AbsDiff(a, b int) int {
	diff := a - b
	return Abs(diff)
}

func PosMod(a, b int) int {
	return ((a % b) + b) % b
}

// Abs64 returns the absolute value of x for int64
func Abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinMax returns both min and max of two integers
func MinMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// Clamp constrains a value between min and max
func Clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// GCD returns the greatest common divisor of a and b
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return Abs(a)
}

// LCM returns the least common multiple of a and b
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return Abs(a*b) / GCD(a, b)
}

// LCMSlice returns the least common multiple of all integers in the slice
func LCMSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = LCM(result, nums[i])
	}
	return result
}

// GCDSlice returns the greatest common divisor of all integers in the slice
func GCDSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = GCD(result, nums[i])
	}
	return result
}

// Sign returns -1, 0, or 1 based on the sign of x
func Sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

// Pow returns a^b for integers (b must be non-negative)
func Pow(a, b int) int {
	if b < 0 {
		panic("Pow: negative exponent")
	}
	result := 1
	for b > 0 {
		if b&1 == 1 {
			result *= a
		}
		a *= a
		b >>= 1
	}
	return result
}

// ModPow returns (a^b) mod m using fast exponentiation
func ModPow(a, b, m int) int {
	if b < 0 {
		panic("ModPow: negative exponent")
	}
	result := 1
	a = a % m
	for b > 0 {
		if b&1 == 1 {
			result = (result * a) % m
		}
		a = (a * a) % m
		b >>= 1
	}
	return result
}

// Sum returns the sum of all integers in the slice
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Product returns the product of all integers in the slice
func Product(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 1
	for _, n := range nums {
		result *= n
	}
	return result
}

// MinSlice returns the minimum value in a slice
func MinSlice(nums []int) int {
	if len(nums) == 0 {
		panic("MinSlice: empty slice")
	}
	min := nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
	}
	return min
}

// MaxSlice returns the maximum value in a slice
func MaxSlice(nums []int) int {
	if len(nums) == 0 {
		panic("MaxSlice: empty slice")
	}
	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

// ManhattanDistance returns the Manhattan distance between two 2D points
func ManhattanDistance(x1, y1, x2, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}
