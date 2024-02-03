package docs_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"database/sql"
	"strings"
	"os"
	"fmt"
	//"time"
	"github.com/go-sql-driver/mysql"
)


var _ = Describe("Docs", func() {
	AWS_S3_ACCESS_KEY := os.Getenv("AWS_S3_ACCESS_KEY")
	AWS_S3_SECRET_KEY := os.Getenv("AWS_S3_SECRET_KEY")

	When("instance is running in staging", Ordered, func() {
		var db *sql.DB

		cfg := mysql.Config{
			User:   "root",
			Passwd: "",
			Net:    "tcp",
			Addr:   "fe:9030",
			AllowNativePasswords: true,
		}

		BeforeAll(func() {
			db, _ = sql.Open("mysql", cfg.FormatDSN())
			db.SetMaxOpenConns(1)
		})

		AfterAll(func() {
			var err error
			_, err = db.Exec(`DROP TABLE IF EXISTS DocsQA.user_behavior_inferred`)
			Expect(err).ToNot(HaveOccurred())
			_, err = db.Exec(`DROP DATABASE IF EXISTS DocsQA`)
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
			b, err := os.ReadFile("SQL/files_table_fxn.sql")
			if err != nil {
				fmt.Print(err)
			}
			SQL := string(b)
			re := strings.NewReplacer( "AAAAAAAAAAAAAAAAAAAA", AWS_S3_ACCESS_KEY, "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB", AWS_S3_SECRET_KEY)

			SQLWithKey := re.Replace(SQL)
			_, err = db.Exec(SQLWithKey)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
