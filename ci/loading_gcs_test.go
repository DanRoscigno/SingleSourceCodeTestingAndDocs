package docs_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Docs", func() {

	When("Loading from GCS docs/loading/gcs", Ordered, func() {

		BeforeAll(func() {
			By("Setting the replication number")
			SQL := SQLFromFile("SQL/loading/cloud/gcs/2-admin-set-repl-num.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("Creating a database")
			SQL = SQLFromFile("SQL/loading/cloud/gcs/3-ddl.sql")
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
		/*
		   It("Use the FILES table fxn", func() {

		       By("Selecting directly from GCS")
		       SQL := SQLFromFile("SQL/loading/cloud/gcs/1-select-from-files.sql")
		       SQLWithCreds := AddGCSCredentials(SQL)
		       _, err := db.Exec(SQLWithCreds)
		       Expect(err).ToNot(HaveOccurred())

		       By("Create Table As Select (CTAS)")
		       SQL = SQLFromFile("SQL/loading/cloud/gcs/4-ctas-inferred.sql")
		       SQLWithCreds = AddGCSCredentials(SQL)
		       _, err = db.Exec(SQLWithCreds)
		       Expect(err).ToNot(HaveOccurred())

		       By("Describing the table")
		       //SQL = SQLFromFile("SQL/loading/cloud/gcs/5-describe.sql")
		       SQL = `SELECT COLUMN_NAME, DATA_TYPE from INFORMATION_SCHEMA.COLUMNS where TABLE_NAME = 'user_behavior_inferred';`
		       rows, err := db.Query(SQL)
		       Expect(err).NotTo(HaveOccurred())
		       defer rows.Close()

		       fieldTypes := []string{}
		       for rows.Next() {
		           var COLUMN_NAME string
		           var DATA_TYPE string
		           err := rows.Scan(&COLUMN_NAME, &DATA_TYPE)
		           Expect(err).NotTo(HaveOccurred())
		           fieldTypes = append(fieldTypes, COLUMN_NAME + "-" + DATA_TYPE)
		       }
		       Expect(fieldTypes).To(ContainElement("UserID-bigint"))
		       Expect(fieldTypes).To(ContainElement("ItemID-bigint"))
		       Expect(fieldTypes).To(ContainElement("CategoryID-bigint"))
		       Expect(fieldTypes).To(ContainElement("BehaviorType-varchar"))
		       Expect(fieldTypes).To(ContainElement("Timestamp-varchar"))

		       By("Selecting from the table")
		       SQL = SQLFromFile("SQL/loading/cloud/gcs/6-select.sql")
		       _, err = db.Exec(SQL)
		       Expect(err).ToNot(HaveOccurred())

		   })

		   It("Create table and then load with FILES table fxn", func() {

		       By("Specifying the table schema")
		       SQL := SQLFromFile("SQL/loading/cloud/gcs/7-ddl-table.sql")
		       _, err := db.Exec(SQL)
		       Expect(err).ToNot(HaveOccurred())

		       By("Inserting")
		       SQL = SQLFromFile("SQL/loading/cloud/gcs/8-insert.sql")
		       SQLWithCreds := AddGCSCredentials(SQL)
		       _, err = db.Exec(SQLWithCreds)
		       Expect(err).ToNot(HaveOccurred())

		       By("Selecting")
		       SQL = SQLFromFile("SQL/loading/cloud/gcs/9-select.sql")
		       _, err = db.Exec(SQL)
		       Expect(err).ToNot(HaveOccurred())

		   })

		   It("Create table and then load with BROKER LOAD", func() {

		       By("Creating a table")
		       SQL := SQLFromFile("SQL/loading/cloud/gcs/11-ddl-table.sql")
		       _, err := db.Exec(SQL)
		       Expect(err).ToNot(HaveOccurred())

		       By("another load label")
		       SQL = SQLFromFile("SQL/loading/cloud/gcs/12-load-label.sql")
		       SQLWithCreds := AddGCSCredentials(SQL)
		       _, err = db.Exec(SQLWithCreds)
		       Expect(err).ToNot(HaveOccurred())

		       By("Selecting")
		       SQL = SQLFromFile("SQL/loading/cloud/gcs/14-select.sq")
		       _, err = db.Exec(SQL)
		       Expect(err).ToNot(HaveOccurred())

		   })
		*/
		It("Create table and then load with Pipe and FILES", func() {

			By("Pipe DDL")
			SQL := SQLFromFile("SQL/loading/cloud/gcs/15-ddl-pipe.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("Create a pipe")
			SQL = SQLFromFile("SQL/loading/cloud/gcs/16-create-pipe.sql")
			SQLWithCreds := AddGCSCredentials(SQL)
			_, err = db.Exec(SQLWithCreds)
			Expect(err).ToNot(HaveOccurred())

			// sleep to give the pipe time to connect to the provider
			// and register the names of files
			fmt.Println("Sleeping for 50 seconds while pipe gets primed")
			time.Sleep(50 * time.Second)

			By("Show pipes")
			SQL = SQLFromFile("SQL/loading/cloud/gcs/17-show-pipes.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("Describing the table for pipe")
			SQL = `select COLUMN_NAME, DATA_TYPE from INFORMATION_SCHEMA.COLUMNS where TABLE_NAME = 'user_behavior_from_pipe';`
			rows, err := db.Query(SQL)
			Expect(err).NotTo(HaveOccurred())
			defer rows.Close()

			fmt.Println("Checking schema created with Pipe")
			fmt.Println("COLUMN_NAME\tDATA_TYPE")
			fieldTypes := []string{}
			for rows.Next() {
				var COLUMN_NAME string
				var DATA_TYPE string
				err := rows.Scan(&COLUMN_NAME, &DATA_TYPE)
				Expect(err).NotTo(HaveOccurred())
				fieldTypes = append(fieldTypes, COLUMN_NAME+"-"+DATA_TYPE)
				fmt.Println(COLUMN_NAME + "\t" + DATA_TYPE)
			}
			Expect(fieldTypes).To(ContainElement("UserID-int"))
			Expect(fieldTypes).To(ContainElement("ItemID-int"))
			Expect(fieldTypes).To(ContainElement("CategoryID-int"))
			Expect(fieldTypes).To(ContainElement("BehaviorType-varchar"))
			Expect(fieldTypes).To(ContainElement("Timestamp-datetime"))

			By("Check pipe loading status")
			SQL = `SELECT FILE_NAME, LOAD_STATE from INFORMATION_SCHEMA.pipe_files where PIPE_NAME = 'user_behavior_pipe';`
			rows, err = db.Query(SQL)
			Expect(err).NotTo(HaveOccurred())
			defer rows.Close()

			fmt.Println("Checking schema created with Pipe")
			fmt.Println("Filename\tLoading state")
			fileState := []string{}
			for rows.Next() {
				var fileName string
				var loadState string
				err := rows.Scan(&fileName, &loadState)
				Expect(err).NotTo(HaveOccurred())
				fileState = append(fileState, fileName+"-"+loadState)
				fmt.Println(fileName + "\t" + loadState)
			}

		})
	})
})
