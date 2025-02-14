package cli

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/epeleh/sample-number-table/cli/common"
	"github.com/epeleh/sample-number-table/cli/fibonaccitable"
	"github.com/epeleh/sample-number-table/cli/primetable"
)

func StartDialog() {
	width, height := askForTableDimensions()
	tableType := askForTableType()
	arithmeticOperation := askForArithmeticOperation()

	printTableFunc := map[common.TableType]func(uint, uint, common.ArithmeticOperation){
		common.Prime:     primetable.Print,
		common.Fibonacci: fibonaccitable.Print,
	}[tableType]

	fmt.Println()
	printTableFunc(width, height, arithmeticOperation)
	fmt.Println()
}

func askForTableDimensions() (width, height uint) {
	for scanner := bufio.NewScanner(os.Stdin); ; {
		fmt.Print("=> Please give table dimensions (<width>x<height>)\n-> ")
		if !scanner.Scan() {
			continue
		}

		dimensions := strings.ReplaceAll(strings.ToLower(scanner.Text()), " ", "")
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

func askForTableType() common.TableType {
	for scanner := bufio.NewScanner(os.Stdin); ; {
		fmt.Print("=> Should I use (P)rime numbers or (F)ibonacci numbers?\n-> ")
		if !scanner.Scan() {
			continue
		}

		switch strings.TrimSpace(strings.ToLower(scanner.Text())) {
		case "p", "prime":
			return common.Prime
		case "f", "fib", "fibonacci":
			return common.Fibonacci
		}
	}
}

func askForArithmeticOperation() common.ArithmeticOperation {
	for scanner := bufio.NewScanner(os.Stdin); ; {
		fmt.Print("=> Multiplication (*) or Addition (+)\n-> ")
		if !scanner.Scan() {
			continue
		}

		switch strings.TrimSpace(strings.ToLower(scanner.Text())) {
		case "m", "mul", "multiplication", "*":
			return common.Multiplication
		case "a", "add", "addition", "+":
			return common.Addition
		}
	}
}
