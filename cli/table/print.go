package table

import (
	"fmt"

	"github.com/epeleh/sample-number-table/cli/common"
	"github.com/epeleh/sample-number-table/cli/numsequences"
)

// Print generates a table of numbers and writes it to stdout
func Print(
	width, height uint,
	numberSequence common.NumberSequence,
	arithmeticOperation common.ArithmeticOperation,
) {
	numberSequenceFunc := map[common.NumberSequence]func(uint) int{
		common.Prime:     numsequences.Prime,
		common.Fibonacci: numsequences.Fibonacci,
	}[numberSequence]

	for row := uint(1); row <= height; row++ {
		for column := uint(1); column <= width; column++ {
			y := numberSequenceFunc(row)
			x := numberSequenceFunc(column)

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
