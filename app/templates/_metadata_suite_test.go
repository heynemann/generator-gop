package metadata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestChannel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Metadata Suite")
}
