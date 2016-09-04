package cmd_test

import (
	"bytes"
	"fmt"
	"io"
	"os"

	. "<%= importName %>/cmd"
	. "<%= importName %>/metadata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Version Command", func() {
	var resetStdout func()
	var readStdout func() string

	mockStdOut := func() (func() string, func()) {
		stdout := os.Stdout
		r, w, err := os.Pipe()
		Expect(err).NotTo(HaveOccurred())
		os.Stdout = w

		return func() string {
				var buf bytes.Buffer
				_, err := io.Copy(&buf, r)
				Expect(err).NotTo(HaveOccurred())
				r.Close()
				return buf.String()
			}, func() {
				w.Close()
				os.Stdout = stdout
			}
	}

	BeforeEach(func() {
		readStdout, resetStdout = mockStdOut()
	})
	AfterEach(func() {
		resetStdout()
		resetStdout = nil
		readStdout = nil
	})

	It("Should write proper version", func() {
		VersionCmd.Run(VersionCmd, []string{})
		resetStdout()
		result := readStdout()
		Expect(result).To(BeEquivalentTo(fmt.Sprintf("<%= name %> v%s\n", VERSION)))
	})
})
