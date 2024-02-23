package docs_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("QuickstartIceberg", func() {

	When("Running the Iceberg Quick Start", Ordered, func() {

		// The database is already initialized, and a connection
		// is available with the variable `db` which is setup
		// in the helpers.go file.

		BeforeAll(func() {
		})

		AfterAll(func() {
		})

		It("DDL: External Catalog", func() {
			By("Create the external CATALOG")
			SQL := SQLFromFile("SQL/quickstart/iceberg/01-create-catalog.sql")
			_, err := db.Exec(SQL)

			By("SET CATALOG")
			SQL = SQLFromFile("SQL/quickstart/iceberg/02-set-catalog.sql")
			_, err = db.Exec(SQL)

			By("Basic SELECT")
			SQL = SQLFromFile("SQL/quickstart/iceberg/03-select.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())
		})
		It("SQL: SELECT FROM the Iceberg table", func() {

			By("Selecting from the table")
			//SQL := SQLFromFile("SQL/quickstart/iceberg/select.sql")
			// +-------+-------------+
			// | trips | hour_of_day |
			// +-------+-------------+
			// |  5381 |          18 |
			// |  5253 |          17 |
			// |  5091 |          16 |
			// |  4736 |          15 |
			// |  4393 |          14 |

			SQL := SQLFromFile("SQL/quickstart/iceberg/04-summarize-trips.sql")
			rows, err := db.Query(SQL)
			Expect(err).NotTo(HaveOccurred())
			defer rows.Close()

			records := []string{}
			for rows.Next() {
				var trips string
				var hour_of_day string
				err := rows.Scan(&trips, &hour_of_day)
				Expect(err).NotTo(HaveOccurred())
				records = append(records, trips + "-" + hour_of_day)
			}
			Expect(records).To(ContainElement("5381-18"))
			Expect(records).To(ContainElement("5253-17"))
			Expect(records).To(ContainElement("4736-15"))
			Expect(records).To(ContainElement("4393-14"))
		})
	})
})
