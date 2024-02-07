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

var _ = BeforeSuite(func() {
	By("Connecting to StarRocks FE")
	db, _ = GetDSNConnection()
	db.SetMaxOpenConns(1)
	Expect(db.Ping()).Should(Succeed())
})
