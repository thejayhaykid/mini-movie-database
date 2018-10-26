package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB connects to the mySQL database and returns that connection.
//  Connection string format is user:password@tcp(localhost:5555)/dbname?charset=utf8
func ConnectDB(conStr string) *sql.DB {
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		fmt.Println("ERROR connecting to database.")
		panic(err)
	}
	return db
}
