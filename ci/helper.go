package docs_test

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/go-sql-driver/mysql"
	. "github.com/onsi/gomega"
)

var db *sql.DB
var AWS_S3_ACCESS_KEY = os.Getenv("AWS_S3_ACCESS_KEY")
var AWS_S3_SECRET_KEY = os.Getenv("AWS_S3_SECRET_KEY")
var GCS_SERVICE_ACCOUNT_EMAIL = os.Getenv("GCS_SERVICE_ACCOUNT_EMAIL")
var GCS_SERVICE_ACCOUNT_PRIVATE_KEY_ID = os.Getenv("GCS_SERVICE_ACCOUNT_PRIVATE_KEY_ID")
var GCS_SERVICE_ACCOUNT_PRIVATE_KEY = os.Getenv("GCS_SERVICE_ACCOUNT_PRIVATE_KEY")
var AZURE_ADLS2_SHARED_KEY = os.Getenv("AZURE_ADLS2_SHARED_KEY")

func GetDSNConnection() (*sql.DB, error) {
	SR_FE_HOST := os.Getenv("SR_FE_HOST")
	fmt.Println("SR HOST is " + SR_FE_HOST)
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 SR_FE_HOST + ":9030",
		AllowNativePasswords: true,
	}
	return sql.Open("mysql", cfg.FormatDSN())
}

func SQLFromFile(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return string(bytes)
}

func AddAWSCredentials(sql string) string {
	re := strings.NewReplacer(
		"AAAAAAAAAAAAAAAAAAAA", AWS_S3_ACCESS_KEY,
		"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB", AWS_S3_SECRET_KEY,
	)
	return re.Replace(sql)
}

func AddGCSCredentials(sql string) string {
	decodedKey, err := base64.StdEncoding.DecodeString(GCS_SERVICE_ACCOUNT_PRIVATE_KEY)
	Expect(err).ToNot(HaveOccurred())
	re := strings.NewReplacer(
		"sampledatareader@xxxxx-xxxxxx-000000.iam.gserviceaccount.com", GCS_SERVICE_ACCOUNT_EMAIL,
		"baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", GCS_SERVICE_ACCOUNT_PRIVATE_KEY_ID,
		"-----BEGIN PRIVATE KEY----- ----END PRIVATE KEY-----", string(decodedKey),
	)
	return re.Replace(sql)
}

func AddAzureCredentials(sql string) string {
	re := strings.NewReplacer(
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb+ccccccccc==", AZURE_ADLS2_SHARED_KEY,
	)
	return re.Replace(sql)
}

func LongRunningScript(filename string) {
	longJob := exec.Command(filename)
	err := longJob.Start()
	Expect(err).ToNot(HaveOccurred())
	err = longJob.Wait()
	Expect(err).ToNot(HaveOccurred())
}

func ResetSettings() {
}
