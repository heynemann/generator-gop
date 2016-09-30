package metadata_test

import (
	"github.com/hashicorp/go-version"
	. "<%= importName %>/metadata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metadata", func() {
	Describe("With default version", func() {
		It("Should have proper version", func() {
			v := GetVersion()
			_, err := version.NewVersion(v)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("With custom version", func() {
		It("Should return beta version", func() {
			v := GetVersion(&Version{
				Major: 10,
				Minor: 12,
				Patch: 25,
				Beta:  true,
			})

			Expect(v).To(Equal("10.12.25-beta"))

			_, err := version.NewVersion(v)
			Expect(err).NotTo(HaveOccurred())
		})

		It("Should return rc version", func() {
			v := GetVersion(&Version{
				Major:            10,
				Minor:            12,
				Patch:            25,
				ReleaseCandidate: 3,
			})

			Expect(v).To(Equal("10.12.25-rc.3"))

			_, err := version.NewVersion(v)
			Expect(err).NotTo(HaveOccurred())
		})

		It("Should return beta version if both are set", func() {
			v := GetVersion(&Version{
				Major:            10,
				Minor:            12,
				Patch:            25,
				Beta:             true,
				ReleaseCandidate: 3,
			})

			Expect(v).To(Equal("10.12.25-beta"))

			_, err := version.NewVersion(v)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
