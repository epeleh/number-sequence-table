package numsequences_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/epeleh/number-sequence-table/cli/numsequences"
)

var _ = Describe("Fibonacci", func() {
	knownFibonacciNumbers := []int{ // Source: https://r-knott.surrey.ac.uk/fibonacci/fibtable.html
		0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765,
		10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269, 2178309,
		3524578, 5702887, 9227465, 14930352, 24157817, 39088169, 63245986, 102334155, 165580141,
		267914296, 433494437, 701408733, 1134903170, 1836311903, 2971215073, 4807526976, 7778742049,
		12586269025, 20365011074, 32951280099, 53316291173, 86267571272, 139583862445, 225851433717,
	}

	for n, expectedResult := range knownFibonacciNumbers {
		When(fmt.Sprintf("n = %d", uint(n)), func() {
			It(fmt.Sprintf("returns %d", expectedResult), func() {
				Expect(numsequences.Fibonacci(uint(n))).To(Equal(expectedResult))
			})
		})
	}
})
