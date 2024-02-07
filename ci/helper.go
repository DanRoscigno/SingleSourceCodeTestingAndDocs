
package docs_test

import (
    "database/sql"
	"github.com/go-sql-driver/mysql"
)
var db *sql.DB

func GetDSNConnection() (*sql.DB, error) {
    cfg := mysql.Config{
	User:   "root",
	Passwd: "",
	Net:    "tcp",
	Addr:   "fe:9030",
	AllowNativePasswords: true,
    }
    return sql.Open("mysql", cfg.FormatDSN())
}

