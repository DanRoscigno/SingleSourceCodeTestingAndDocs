
package docs_test

import (
    "database/sql"
    "os"
    "fmt"
    "github.com/go-sql-driver/mysql"
    "os/exec"
    . "github.com/onsi/gomega"
)
var db *sql.DB

func GetDSNConnection() (*sql.DB, error) {
    SR_FE_HOST_PORT := os.Getenv("SR_FE_HOST_PORT")
    cfg := mysql.Config{
	User:   "root",
	Passwd: "",
	Net:    "tcp",
	Addr:   SR_FE_HOST_PORT,
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

func LongRunningScript(filename string) {
    longJob := exec.Command(filename)
    err := longJob.Start()
    Expect(err).ToNot(HaveOccurred())
    err = longJob.Wait()
    Expect(err).ToNot(HaveOccurred())
}

func ResetSettings() {
}
