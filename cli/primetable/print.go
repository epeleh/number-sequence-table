package primetable

import (
	"fmt"
	"math"

	"github.com/epeleh/sample-number-table/cli/common"
)

// Print generates a table of numbers based on the prime numbers sequence and writes it to stdout
func Print(width uint, height uint, arithmeticOperation common.ArithmeticOperation) {
	for row := uint(1); row <= height; row++ {
		for column := uint(1); column <= width; column++ {
			y := primeNumber(row)
			x := primeNumber(column)

			var cellValue int
			switch arithmeticOperation {
			case common.Addition:
				cellValue = x + y
			case common.Multiplication:
				cellValue = x * y
			default:
				panic("unreachable")
			}

			fmt.Printf("%3d ", cellValue)
		}

		fmt.Println()
	}
}

// primeNumber returns the nth prime number
func primeNumber(n uint) (number int) {
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
