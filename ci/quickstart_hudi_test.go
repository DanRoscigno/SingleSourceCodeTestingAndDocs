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

		It("SQL: Work with Hudi", func() {
			By("creating an external catalog")
			SQL := SQLFromFile("SQL/quickstart/hudi/create_catalog.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("creating an external catalog")
			SQL := SQLFromFile("SQL/quickstart/hudi/set_catalog.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("creating an external catalog")
			SQL := SQLFromFile("SQL/quickstart/hudi/show_databases.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("creating an external catalog")
			SQL := SQLFromFile("SQL/quickstart/hudi/use_default.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("creating an external catalog")
			SQL := SQLFromFile("SQL/quickstart/hudi/show_tables.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("creating an external catalog")
			SQL := SQLFromFile("SQL/quickstart/hudi/select.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
