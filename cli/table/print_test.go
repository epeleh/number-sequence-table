package table_test

import (
	"fmt"
	"io"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/epeleh/number-sequence-table/cli/common"
	"github.com/epeleh/number-sequence-table/cli/table"
)

var _ = Describe("Print", func() {
	type printArgs struct {
		width, height       uint
		numberSequence      common.NumberSequence
		arithmeticOperation common.ArithmeticOperation
	}

	testCases := map[printArgs]string{
		printArgs{3, 3, common.Prime, common.Multiplication}: strings.Join([]string{
			"  4   6  10 \n",
			"  6   9  15 \n",
			" 10  15  25 \n",
		}, ""),
		printArgs{5, 4, common.Fibonacci, common.Multiplication}: strings.Join([]string{
			"  1   1   2   3   5 \n",
			"  1   1   2   3   5 \n",
			"  2   2   4   6  10 \n",
			"  3   3   6   9  15 \n",
		}, ""),
		printArgs{11, 2, common.Prime, common.Addition}: strings.Join([]string{
			"  4   5   7   9  13  15  19  21  25  31  33 \n",
			"  5   6   8  10  14  16  20  22  26  32  34 \n",
		}, ""),
		printArgs{10, 10, common.Fibonacci, common.Addition}: strings.Join([]string{
			"  2   2   3   4   6   9  14  22  35  56 \n",
			"  2   2   3   4   6   9  14  22  35  56 \n",
			"  3   3   4   5   7  10  15  23  36  57 \n",
			"  4   4   5   6   8  11  16  24  37  58 \n",
			"  6   6   7   8  10  13  18  26  39  60 \n",
			"  9   9  10  11  13  16  21  29  42  63 \n",
			" 14  14  15  16  18  21  26  34  47  68 \n",
			" 22  22  23  24  26  29  34  42  55  76 \n",
			" 35  35  36  37  39  42  47  55  68  89 \n",
			" 56  56  57  58  60  63  68  76  89 110 \n",
		}, ""),
		printArgs{0, 3, common.Fibonacci, common.Multiplication}: "\n\n\n",
		printArgs{5, 0, common.Prime, common.Multiplication}:     "",
		printArgs{0, 0, common.Prime, common.Addition}:           "",
	}

	for args, expectedStdout := range testCases {
		When(fmt.Sprintf("%+v", args), func() {
			It("prints the table properly", func() {
				stdout := captureStdout(func() {
					table.Print(args.width, args.height, args.numberSequence, args.arithmeticOperation)
				})

				Expect(stdout).To(Equal(expectedStdout))
			})
		})
	}
})

func captureStdout(f func()) string {
	realStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	os.Stdout = realStdout
	w.Close()
	stdout, _ := io.ReadAll(r)
	return string(stdout)
}
