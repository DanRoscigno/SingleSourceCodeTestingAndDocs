package docs_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Docs", func() {

	When("Loading from MinIO docs/loading/minio", Ordered, func() {

		BeforeAll(func() {
			By("Setting the replication number")
			SQL := SQLFromFile("SQL/loading/cloud/minio/2-admin-set-repl-num.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("Creating a database")
			SQL = SQLFromFile("SQL/loading/cloud/minio/3-ddl.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())
		})

		AfterAll(func() {
			var err error

			By("DROP table user_behavior_inferred")
			_, err = db.Exec(`DROP TABLE IF EXISTS user_behavior_inferred`)
			Expect(err).ToNot(HaveOccurred())

			By("DROP DB DocsQA")
			_, err = db.Exec(`DROP DATABASE IF EXISTS mydatabase`)
			Expect(err).ToNot(HaveOccurred())

			By("Reset settings")
			_, err = db.Exec(`ADMIN SET FRONTEND CONFIG ('default_replication_num' = "3");`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("Use the FILES table fxn", func() {

			By("Selecting directly from MinIO")
			SQL := SQLFromFile("SQL/loading/cloud/minio/1-select-from-files.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("Create Table As Seect (CTAS)")
			SQL = SQLFromFile("SQL/loading/cloud/minio/4-ctas-inferred.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("Describing the table")
			//SQL = SQLFromFile("SQL/loading/cloud/minio/5-describe.sql")
			SQL = `select COLUMN_NAME, DATA_TYPE from INFORMATION_SCHEMA.COLUMNS where TABLE_NAME = 'user_behavior_inferred';`
			rows, err := db.Query(SQL)
			Expect(err).NotTo(HaveOccurred())
			defer rows.Close()

			fieldTypes := []string{}
			for rows.Next() {
				var COLUMN_NAME string
				var DATA_TYPE string
				err := rows.Scan(&COLUMN_NAME, &DATA_TYPE)
				Expect(err).NotTo(HaveOccurred())
				fieldTypes = append(fieldTypes, COLUMN_NAME+"-"+DATA_TYPE)
			}
			Expect(fieldTypes).To(ContainElement("UserID-bigint"))
			Expect(fieldTypes).To(ContainElement("ItemID-bigint"))
			Expect(fieldTypes).To(ContainElement("CategoryID-bigint"))
			Expect(fieldTypes).To(ContainElement("BehaviorType-varbinary"))
			Expect(fieldTypes).To(ContainElement("Timestamp-varbinary"))

			By("Selecting from the table")
			SQL = SQLFromFile("SQL/loading/cloud/minio/6-select.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

		})

		It("Create table and then load with FILES table fxn", func() {

			By("Specifying the table schema")
			SQL := SQLFromFile("SQL/loading/cloud/minio/7-ddl-table.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("Inserting")
			SQL = SQLFromFile("SQL/loading/cloud/minio/8-insert.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("Selecting")
			SQL = SQLFromFile("SQL/loading/cloud/minio/9-select.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

		})

		It("Create table and then load with BROKER LOAD", func() {

			By("Creating a table")
			SQL := SQLFromFile("SQL/loading/cloud/minio/11-ddl-table.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("another load label")
			SQL = SQLFromFile("SQL/loading/cloud/minio/12-load-label.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("Selecting")
			SQL = SQLFromFile("SQL/loading/cloud/minio/14-select.sq")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

		})

		It("Create table and then load with Pipe and FILES", func() {

			By("Pipe DDL")
			SQL := SQLFromFile("SQL/loading/cloud/minio/15-ddl-pipe.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("Create a pipe")
			SQL = SQLFromFile("SQL/loading/cloud/minio/16-create-pipe.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("Show pipes")
			SQL = SQLFromFile("SQL/loading/cloud/minio/17-show-pipes.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

		})
	})
})
