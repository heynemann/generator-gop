package metadata_test

import (
	"github.com/hashicorp/go-version"
	. "<%= importName %>/metadata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metadata", func() {
	It("Should have proper version", func() {
		_, err := version.NewVersion(VERSION)
		Expect(err).NotTo(HaveOccurred())
	})
})
