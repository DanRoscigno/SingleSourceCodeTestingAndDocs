package docs_test_test

import (
	"database/sql"
	"log"
	//"fmt"
//	"os"
	"os/exec"
	//"strings"

	"github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("QuickstartBasic", func() {

	//AWS_S3_ACCESS_KEY := os.Getenv("AWS_S3_ACCESS_KEY")
	//AWS_S3_SECRET_KEY := os.Getenv("AWS_S3_SECRET_KEY")

	When("Running a single FE and BE via Docker compose", Ordered, func() {
		var db *sql.DB

		cfg := mysql.Config{
			User:                 "root",
			Passwd:               "",
			Net:                  "tcp",
			Addr:                 "127.0.0.1:9030",
			AllowNativePasswords: true,
		}

		BeforeAll(func() {
			// download the crash data in /tmp/ dir
			// https://stackoverflow.com/questions/16703647/why-does-curl-return-error-23-failed-writing-body
			cmd := exec.Command("/usr/bin/curl","-s","-N","-o","/tmp/NYPD_Crash_Data.csv","https://raw.githubusercontent.com/StarRocks/demo/master/documentation-samples/quickstart/datasets/NYPD_Crash_Data.csv")
			err := cmd.Start()
			// I think I should use this instead of log.Fatal
			//Expect(err).ToNot(HaveOccurred())
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Waiting for curls to straighten...")
			err = cmd.Wait()
			if err != nil {
				log.Fatal(err)
			} else {
				log.Printf("Curls are now straightened")
			}

			// Connect to the database
			db, _ = sql.Open("mysql", cfg.FormatDSN())
			db.SetMaxOpenConns(1)

		})

		AfterAll(func() {
			_, err := db.Exec(`DROP DATABASE IF EXISTS DocsQA`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should be able to run SQL commands", func() {
			By("creating a database")
			_, err := db.Exec(`CREATE DATABASE IF NOT EXISTS quickstart`)
			Expect(err).ToNot(HaveOccurred())
		})

	})
})
