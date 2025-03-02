package cli

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/epeleh/number-sequence-table/cli/common"
	"github.com/epeleh/number-sequence-table/cli/table"
)

var isArgsInput = len(os.Args) > 1
var isPipedInput bool

func init() {
	stat, err := os.Stdin.Stat()
	isPipedInput = err == nil && (stat.Mode()&os.ModeCharDevice) == 0
}

func StartDialog() {
	var scanner *bufio.Scanner
	if isArgsInput {
		scanner = bufio.NewScanner(strings.NewReader(strings.Join(os.Args[1:], "\n")))
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	width, height := askForTableDimensions(scanner)
	numberSequence := askForNumberSequence(scanner)
	arithmeticOperation := askForArithmeticOperation(scanner)

	fmt.Println()
	table.Print(width, height, numberSequence, arithmeticOperation)
	fmt.Println()
}

func askForTableDimensions(scanner *bufio.Scanner) (width, height uint) {
	for {
		fmt.Print("=> Please give table dimensions (<width>x<height>)\n-> ")
		if !scanner.Scan() {
			fmt.Println("3x3")
			return 3, 3 // use some default values if EOF is reached
		}

		input := scanner.Text()
		if isArgsInput || isPipedInput {
			fmt.Println(input)
		}

		dimensions := strings.ReplaceAll(strings.ToLower(input), " ", "")
		if !regexp.MustCompile(`\A\d+x\d+\z`).MatchString(dimensions) {
			continue
		}

		_, err := fmt.Sscanf(dimensions, "%dx%d", &width, &height)
		if err != nil {
			continue
		}

		return width, height
	}
}

func askForNumberSequence(scanner *bufio.Scanner) common.NumberSequence {
	for {
		fmt.Print("=> Should I use (P)rime numbers or (F)ibonacci numbers?\n-> ")
		if !scanner.Scan() {
			fmt.Println("p")
			return common.Prime // use some default value if EOF is reached
		}

		input := scanner.Text()
		if isArgsInput || isPipedInput {
			fmt.Println(input)
		}

		switch strings.TrimSpace(strings.ToLower(input)) {
		case "p", "prime":
			return common.Prime
		case "f", "fib", "fibonacci":
			return common.Fibonacci
		}
	}
}

func askForArithmeticOperation(scanner *bufio.Scanner) common.ArithmeticOperation {
	for {
		fmt.Print("=> Multiplication (*) or Addition (+)\n-> ")
		if !scanner.Scan() {
			fmt.Println("m")
			return common.Multiplication // use some default value if EOF is reached
		}

		input := scanner.Text()
		if isArgsInput || isPipedInput {
			fmt.Println(input)
		}

		switch strings.TrimSpace(strings.ToLower(input)) {
		case "m", "mul", "multiplication", "*":
			return common.Multiplication
		case "a", "add", "addition", "+":
			return common.Addition
		}
	}
}
