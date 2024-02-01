package docs_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"database/sql"
	//"fmt"
	//"time"
	"github.com/go-sql-driver/mysql"
)


var _ = Describe("Docs", func() {

	When("instance is running in staging", Ordered, func() {
		var db *sql.DB

		cfg := mysql.Config{
			User:   "root",
			Passwd: "",
			Net:    "tcp",
			Addr:   "127.0.0.1:9030",
			AllowNativePasswords: true,
		}

		BeforeEach(func() {
			db, _ = sql.Open("mysql", cfg.FormatDSN())
			db.SetMaxOpenConns(1)
		})

		AfterAll(func() {
			_, err := db.Exec(`DROP DATABASE IF EXISTS DocsQA2`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should be able to run SQL commands", func() {
			By("creating a database")
			_, err := db.Exec(`CREATE DATABASE IF NOT EXISTS DocsQA`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should be able to run SQL commands", func() {
			By("choosing a database")
			_, err := db.Exec(`USE DocsQA`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should be able to run SQL commands", func() {
			By("setting the number of replicas")
			_, err := db.Exec(`ADMIN SET FRONTEND CONFIG ('default_replication_num' = "1");`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should be able to run SQL commands", func() {
			By("creating and populating a table")
			_, err := db.Exec(`CREATE TABLE DocsQA.user_behavior_inferred AS SELECT * FROM FILES ( "path" = "s3://starrocks-examples/user_behavior_ten_million_rows.parquet", "format" = "parquet", "aws.s3.region" = "us-east-1", "aws.s3.access_key" = "AKIAZWYRUJIJ3S43RPOA", "aws.s3.secret_key" = "JL+uV5hb6zm5p8m4yWy9iTD5qIn8YGbtv0qYr9qJ");`)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
