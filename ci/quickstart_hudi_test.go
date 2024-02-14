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

		It("DDL: External Catalog", func() {
			By("creating an external catalog")
			SQL := SQLFromFile("SQL/quickstart/hudi/create_catalog.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("Setting the catalog")
			SQL = SQLFromFile("SQL/quickstart/hudi/set_catalog.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("USEing the database")
			SQL = SQLFromFile("SQL/quickstart/hudi/use_database.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("SHOW TABLES")
			SQL = SQLFromFile("SQL/quickstart/hudi/show_tables.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())
		})
		It("SQL: SELECT FROM the Hudi table", func() {

			By("Selecting from the table")
			//SQL := SQLFromFile("SQL/quickstart/hudi/select.sql")
			SQL = `SELECT language, users from hudi_coders_hive;`
			rows, err := db.Query(SQL)
			Expect(err).NotTo(HaveOccurred())
			defer rows.Close()

			records := []string{}
			for rows.Next() {
				var language string
				var users string
				err := rows.Scan(&language, &users)
				Expect(err).NotTo(HaveOccurred())
				records = append(records, language + "-" + users)
			}
			Expect(records).To(ContainElement("Pythin-100000"))
			Expect(records).To(ContainElement("Java-20000"))
			Expect(records).To(ContainElement("Scala-3000"))
		})
	})
})
