package common

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var dsn string = "root:@tcp(localhost:3306)/longmarch01_managesystem"

var dbConn *sql.DB

func init() {
	if dbConn == nil {
		var err error
		dbConn, err = sql.Open("mysql", dsn)
		if err != nil {
			fmt.Println(err)
		}
		dbConn.SetMaxOpenConns(10)
		dbConn.SetMaxIdleConns(5)
	}
}

func GetDbConn() *sql.DB {
	return dbConn
}
