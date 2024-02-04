package docs_test_test

import (
	"database/sql"
	//"log"
	"fmt"
  	"os"
	"os/exec"
	//"strings"

	"github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("QuickstartBasic", func() {

	//AWS_S3_ACCESS_KEY := os.Getenv("AWS_S3_ACCESS_KEY")
	//AWS_S3_SECRET_KEY := os.Getenv("AWS_S3_SECRET_KEY")

	When("Running the basic Quick Start", Ordered, func() {
		var db *sql.DB

		cfg := mysql.Config{
			User:                 "root",
			Passwd:               "",
			Net:                  "tcp",
			Addr:                 "fe:9030",
			AllowNativePasswords: true,
		}

		BeforeAll(func() {
			// download the crash data in /tmp/ dir
			// https://stackoverflow.com/questions/16703647/why-does-curl-return-error-23-failed-writing-body
			By("Downloading the NYPD Crash data")
			NYPDCurl := exec.Command("/usr/bin/curl","-s","-N","-o","/tmp/NYPD_Crash_Data.csv","https://raw.githubusercontent.com/StarRocks/demo/master/documentation-samples/quickstart/datasets/NYPD_Crash_Data.csv")
			err := NYPDCurl.Start()
			Expect(err).ToNot(HaveOccurred())
			err = NYPDCurl.Wait()
			Expect(err).ToNot(HaveOccurred())

			By("Downloading the NOAA weather data")
			WeatherCurl := exec.Command("/usr/bin/curl","-s","-N","-o","/tmp/72505394728.csv","https://raw.githubusercontent.com/StarRocks/demo/master/documentation-samples/quickstart/datasets/72505394728.csv")
			err = WeatherCurl.Start()
			Expect(err).ToNot(HaveOccurred())
			err = WeatherCurl.Wait()
			Expect(err).ToNot(HaveOccurred())

			// Connect to the database
			By("Connecting to StarRocks FE")
			db, err = sql.Open("mysql", cfg.FormatDSN())
			db.SetMaxOpenConns(1)
			Expect(err).ToNot(HaveOccurred())
			Expect(db.Ping()).Should(Succeed())

		})

		AfterAll(func() {
			By("dropping quickstart DB")
			_, err := db.Exec(`DROP DATABASE IF EXISTS quickstart`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("DDL: Setup quickstart DB", func() {
			By("creating a database")
			_, err := db.Exec(`CREATE DATABASE IF NOT EXISTS quickstart`)
			Expect(err).ToNot(HaveOccurred())

			By("choosing a database")
			_, err = db.Exec(`USE quickstart`)
			Expect(err).ToNot(HaveOccurred())

			By("setting the number of replicas")
			_, err = db.Exec(`ADMIN SET FRONTEND CONFIG ('default_replication_num' = "1");`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("DDL: Create quickstart tables", func() {
			By("creating the crash data table")
			b, err := os.ReadFile("SQL/quickstart/basic/NYPD_table.sql")
			if err != nil {
				fmt.Print(err)
			}
			SQL := string(b)
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("creating the weather data table")
			b, err = os.ReadFile("SQL/quickstart/basic/Weather_table.sql")
			if err != nil {
				fmt.Print(err)
			}
			SQL = string(b)
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should be able to load data via stream load", func() {
			By("uploading the NYPD crash data")
			NYPDStreamLoad := exec.Command("SHELL/quickstart/basic/NYPD_stream_load")
			err := NYPDStreamLoad.Start()
			Expect(err).ToNot(HaveOccurred())
			err = NYPDStreamLoad.Wait()
			Expect(err).ToNot(HaveOccurred())

			By("uploading the NOAA weather data")
			WeatherStreamLoad := exec.Command("SHELL/quickstart/basic/Weather_stream_load")
			err = WeatherStreamLoad.Start()
			Expect(err).ToNot(HaveOccurred())
			err = WeatherStreamLoad.Wait()
			Expect(err).ToNot(HaveOccurred())
		})

		It("should be able to query tables", func() {
			By("querying the crash data table")
			b, err := os.ReadFile("SQL/quickstart/basic/CrashesPerHour.sql")
			if err != nil {
				fmt.Print(err)
			}
			SQL := string(b)
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("querying the weather data table")
			b, err = os.ReadFile("SQL/quickstart/basic/AverageTemp.sql")
			if err != nil {
				fmt.Print(err)
			}
			SQL = string(b)
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("JOINing to see impact of low visibility")
			b, err = os.ReadFile("SQL/quickstart/basic/LowVisibility.sql")
			if err != nil {
				fmt.Print(err)
			}
			SQL = string(b)
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("JOINing to see impact of icy weather")
			b, err = os.ReadFile("SQL/quickstart/basic/Icy.sql")
			if err != nil {
				fmt.Print(err)
			}
			SQL = string(b)
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})