package docs_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("QuickstartHudi", func() {

	When("Running the Hudi Quick Start", Ordered, func() {

		// The database is already initialized, and a connection
		// is available with the variable `db` which is setup
		// in the helpers.go file.

		BeforeAll(func() {
		})

		AfterAll(func() {
		})

		It("DDL: Setup quickstart DB", func() {
			By("creating a database")
			SQL := SQLFromFile("SQL/quickstart/hudi/quickstart_DB.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())
		})

	})
})
