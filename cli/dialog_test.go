package cli_test

import (
	"io"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/epeleh/number-sequence-table/cli"
)

var _ = Describe("StartDialog", func() {
	Context("when user's stdin is empty", func() {
		var origArgs, commandArguments []string
		var stdout string

		JustBeforeEach(func() {
			origArgs = os.Args[:]
			os.Args = os.Args[:1]
			os.Args = append(os.Args, commandArguments...)
			stdout = captureStdout(func() { cli.StartDialog() })
		})
		JustAfterEach(func() { os.Args = origArgs[:] })

		Context("when the user provided some command arguments", func() {
			BeforeEach(func() { commandArguments = []string{"5x4", "F", "M"} })

			It("uses the command arguments as answers", func() {
				Expect(stdout).To(Equal(strings.Join([]string{
					"=> Please give table dimensions (<width>x<height>)\n",
					"-> 5x4\n",
					"=> Should I use (P)rime numbers or (F)ibonacci numbers?\n",
					"-> F\n",
					"=> Multiplication (*) or Addition (+)\n",
					"-> M\n",
					"\n",
					"  1   1   2   3   5 \n",
					"  1   1   2   3   5 \n",
					"  2   2   4   6  10 \n",
					"  3   3   6   9  15 \n",
					"\n",
				}, "")))
			})
		})

		Context("when the user didn't provide any command arguments", func() {
			BeforeEach(func() { commandArguments = []string{} })

			It("uses the default answers", func() {
				Expect(stdout).To(Equal(strings.Join([]string{
					"=> Please give table dimensions (<width>x<height>)\n",
					"-> 3x3\n",
					"=> Should I use (P)rime numbers or (F)ibonacci numbers?\n",
					"-> p\n",
					"=> Multiplication (*) or Addition (+)\n",
					"-> m\n",
					"\n",
					"  4   6  10 \n",
					"  6   9  15 \n",
					" 10  15  25 \n",
					"\n",
				}, "")))
			})
		})
	})
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
