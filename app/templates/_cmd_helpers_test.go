package cmd_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/onsi/gomega"
)

var ResetStdout func()
var ReadStdout func() string

func MockStdout() {
	stdout := os.Stdout
	r, w, err := os.Pipe()
	Expect(err).NotTo(HaveOccurred())
	os.Stdout = w

	ReadStdout = func() string {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		Expect(err).NotTo(HaveOccurred())
		r.Close()
		return buf.String()
	}

	ResetStdout = func() {
		w.Close()
		os.Stdout = stdout
	}
}
