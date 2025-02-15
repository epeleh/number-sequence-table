package numsequences

import "math"

// Prime returns the nth prime number
func Prime(n uint) (number int) {
	for i := uint(0); i < n; {
		number++
		if isPrime(number) {
			i++
		}
	}

	return number
}

// isPrime checks if a number is prime using memoization
func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	// Check if the result is already computed
	if result, exists := primeNumbersCache[number]; exists {
		return result
	}

	// Check divisibility from 2 to sqrt(number)
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			primeNumbersCache[number] = false
			return false
		}
	}

	primeNumbersCache[number] = true
	return true
}

var primeNumbersCache = map[int]bool{}
