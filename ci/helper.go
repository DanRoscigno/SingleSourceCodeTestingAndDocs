
package docs_test

import (
    "database/sql"
    "os"
    "strings"
    "fmt"
    "github.com/go-sql-driver/mysql"
    "os/exec"
    . "github.com/onsi/gomega"
)
var db *sql.DB
var AWS_S3_ACCESS_KEY = os.Getenv("AWS_S3_ACCESS_KEY")
var AWS_S3_SECRET_KEY = os.Getenv("AWS_S3_SECRET_KEY")

func GetDSNConnection() (*sql.DB, error) {
    SR_FE_HOST := os.Getenv("SR_FE_HOST")
    fmt.Print("SR HOST is " + SR_FE_HOST)
    cfg := mysql.Config{
	User:   "root",
	Passwd: "",
	Net:    "tcp",
	Addr:   SR_FE_HOST + ":9030",
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

func LongRunningScript(filename string) {
    longJob := exec.Command(filename)
    err := longJob.Start()
    Expect(err).ToNot(HaveOccurred())
    err = longJob.Wait()
    Expect(err).ToNot(HaveOccurred())
}

func ResetSettings() {
}
