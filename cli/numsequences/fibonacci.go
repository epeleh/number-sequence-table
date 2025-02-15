package numsequences

// Fibonacci recursively finds the nth fibonacci number using memoization
func Fibonacci(n uint) int {
	// Check if the result is already computed
	if result, exists := fibonacciNumbersCache[n]; exists {
		return result
	}

	// Compute and store the result into the cache
	fibonacciNumbersCache[n] = Fibonacci(n-1) + Fibonacci(n-2)
	return fibonacciNumbersCache[n]
}

var fibonacciNumbersCache = map[uint]int{0: 0, 1: 1}
