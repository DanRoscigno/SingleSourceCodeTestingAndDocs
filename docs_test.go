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
		})

		AfterAll(func() {
			_, err := db.Exec(`DROP DATABASE IF EXISTS DocsQA`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should be able to run SQL commands", func() {
			By("creating a database")
			_, err := db.Exec(`CREATE DATABASE IF NOT EXISTS DocsQA`)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
