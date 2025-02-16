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
	It("works", func() {
		origArgs := os.Args[:]
		os.Args = os.Args[:1]
		defer func() { os.Args = origArgs[:] }()

		stdout := captureStdout(func() { cli.StartDialog() })
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
