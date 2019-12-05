package common

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var dsn string = "root:123456@tcp(localhost:3306)/longmarch01_managesystem"

var dbConn *sql.DB

func init(){
	if dbConn == nil{
		var err error
		dbConn, err = sql.Open("mysql", dsn)
		if err != nil{
			fmt.Println(err)
		}
		dbConn.SetMaxOpenConns(10)
		dbConn.SetMaxIdleConns(5)
	}
}

func GetDbConn() *sql.DB{
	return dbConn
}

