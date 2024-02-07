package docs_test

import (
	"fmt"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Docs", func() {
	AWS_S3_ACCESS_KEY := os.Getenv("AWS_S3_ACCESS_KEY")
	AWS_S3_SECRET_KEY := os.Getenv("AWS_S3_SECRET_KEY")

	When("Loading from S3 docs/loading/s3", Ordered, func() {

		AfterAll(func() {
			var err error
			By("DROP table user_behavior_inferred")
			_, err = db.Exec(`DROP TABLE IF EXISTS DocsQA.user_behavior_inferred`)
			Expect(err).ToNot(HaveOccurred())

			By("DROP DB DocsQA")
			_, err = db.Exec(`DROP DATABASE IF EXISTS DocsQA`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("DDL: setup DocsQA DB", func() {
			By("creating a database")
			_, err := db.Exec(`CREATE DATABASE IF NOT EXISTS DocsQA`)
			Expect(err).ToNot(HaveOccurred())

			By("choosing a database")
			_, err = db.Exec(`USE DocsQA`)
			Expect(err).ToNot(HaveOccurred())

			By("setting the number of replicas")
			_, err = db.Exec(`ADMIN SET FRONTEND CONFIG ('default_replication_num' = "1");`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("Use the FILES table fxn", func() {
			By("creating and populating a table")
			b, err := os.ReadFile("SQL/files_table_fxn.sql")
			if err != nil {
				fmt.Print(err)
			}
			SQL := string(b)
			re := strings.NewReplacer(
				"AAAAAAAAAAAAAAAAAAAA", AWS_S3_ACCESS_KEY,
				"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB", AWS_S3_SECRET_KEY,
			)

			SQLWithKey := re.Replace(SQL)
			_, err = db.Exec(SQLWithKey)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
