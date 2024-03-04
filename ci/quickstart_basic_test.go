package docs_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("QuickstartBasic", func() {

	When("Running the basic Quick Start", Ordered, func() {

		// The database is already initialized, and a connection
		// is available with the variable `db` which is setup
		// in the helpers.go file.

		BeforeAll(func() {
			// download the crash data in /tmp/ dir
			// https://stackoverflow.com/questions/16703647/why-does-curl-return-error-23-failed-writing-body
			By("Downloading the NYPD Crash data")
			LongRunningScript("SHELL/quickstart/basic/NYPD_download")

			By("Downloading the NOAA weather data")
			LongRunningScript("SHELL/quickstart/basic/Weather_download")
		})

		AfterAll(func() {
			By("dropping quickstart DB")
			_, err := db.Exec(`DROP DATABASE IF EXISTS quickstart`)
			Expect(err).ToNot(HaveOccurred())

			By("Reset settings")
			_, err = db.Exec(`ADMIN SET FRONTEND CONFIG ('default_replication_num' = "3");`)
			Expect(err).ToNot(HaveOccurred())
		})

		It("DDL: Setup quickstart DB", func() {
			By("creating a database")
			SQL := SQLFromFile("SQL/quickstart/basic/quickstart_DB.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())
		})

		It("DDL: Create quickstart tables", func() {
			By("creating the crash data table")
			SQL := SQLFromFile("SQL/quickstart/basic/NYPD_table.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("creating the weather data table")
			SQL = SQLFromFile("SQL/quickstart/basic/Weather_table.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should be able to load data via stream load", func() {
			By("uploading the NYPD crash data")
			LongRunningScript("SHELL/quickstart/basic/NYPD_stream_load")

			By("uploading the NOAA weather data")
			LongRunningScript("SHELL/quickstart/basic/Weather_stream_load")
		})

		It("should be able to query tables", func() {
			By("querying the crash data table")
			SQL := SQLFromFile("SQL/quickstart/basic/CrashesPerHour.sql")
			_, err := db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("querying the weather data table")
			SQL = SQLFromFile("SQL/quickstart/basic/AverageTemp.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("JOINing to see impact of low visibility")
			SQL = SQLFromFile("SQL/quickstart/basic/LowVisibility.sql")
			_, err = db.Exec(SQL)
			Expect(err).ToNot(HaveOccurred())

			By("JOINing to see impact of icy weather")
			SQL = SQLFromFile("SQL/quickstart/basic/Icy.sql")
			rows, err := db.Query(SQL)
			Expect(err).NotTo(HaveOccurred())
			defer rows.Close()

			records := []string{}
			fmt.Println("Checking JOIN of Crashes and Weather")
			fmt.Println("Crashes\tTemp_F\tPrecipitation\tHour")
			for rows.Next() {
				var Crashes string
				var Temp_F string
				var Precipitation string
				var Hour string
				err := rows.Scan(&Crashes, &Temp_F, &Precipitation, &Hour)
				Expect(err).NotTo(HaveOccurred())
				records = append(records, Crashes+"-"+Temp_F+"-"+Precipitation+"-"+Hour)
				fmt.Println(Crashes + "\t" + Temp_F + "\t" + Precipitation + "\t" + Hour)
			}
			Expect(records).To(ContainElement("192-34-0.09-18 Jan 2015 08:00"))
			Expect(records).To(ContainElement("138-33.5-0.02-18 Jan 2015 07:00"))
		})
	})
})
