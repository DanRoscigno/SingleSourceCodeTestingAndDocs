package docs_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDocs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Docs Suite")
}
