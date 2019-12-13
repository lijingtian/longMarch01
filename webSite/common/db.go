package common

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var mysqlConn *sql.DB

var dsn string = "root:123456@tcp(localhost:3306)/longmarch01_website"

func init(){
	if mysqlConn == nil{
		var err error 
		mysqlConn, err = sql.Open("mysql", dsn)
		if err != nil{
			fmt.Println(err)
			return
		}
		mysqlConn.SetMaxOpenConns(10)
		mysqlConn.SetMaxIdleConns(5)
	}
}

func GetMysqlConn() *sql.DB{
	return mysqlConn
}