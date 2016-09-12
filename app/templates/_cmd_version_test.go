package cmd_test

import (
	"fmt"

	. "<%= importName %>/cmd"
	. "<%= importName %>/metadata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Version Command", func() {
	BeforeEach(func() {
		MockStdout()
	})
	AfterEach(func() {
		ResetStdout()
	})

	It("Should write proper version", func() {
		VersionCmd.Run(VersionCmd, []string{})
		ResetStdout()
		result := ReadStdout()
		Expect(result).To(BeEquivalentTo(fmt.Sprintf("<%= name %> v%s\n", GetVersion())))
	})
})
