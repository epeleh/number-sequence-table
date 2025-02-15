package fibonaccitable

import (
	"fmt"

	"github.com/epeleh/sample-number-table/cli/common"
)

// Print generates a table of numbers based on the Fibonacci sequence and writes it to stdout
func Print(width uint, height uint, arithmeticOperation common.ArithmeticOperation) {
	for row := uint(1); row <= height; row++ {
		for column := uint(1); column <= width; column++ {
			y := fibonacciNumber(row)
			x := fibonacciNumber(column)

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

// fibonacciNumber recursively finds the nth fibonacci number using memoization
func fibonacciNumber(n uint) int {
	// Check if the result is already computed
	if result, exists := fibonacciNumbersCache[n]; exists {
		return result
	}

	// Compute and store the result into the cache
	fibonacciNumbersCache[n] = fibonacciNumber(n-1) + fibonacciNumber(n-2)
	return fibonacciNumbersCache[n]
}

var fibonacciNumbersCache = map[uint]int{0: 0, 1: 1}
