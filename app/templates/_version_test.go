package metadata_test

import (
	. "<%= url.replace('http://', '').replace('https://', '') %>/metadata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metadata", func() {
	It("Should have proper version", func() {
		Expect(VERSION).To(Equal("<%= version %>"))
	})
})
