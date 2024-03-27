package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/Avengers?parseTime=true")
	if err != nil {
		fmt.Println("Failed Connect to Database: ", err)
		return
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Error Ping to Database: ", err)
		return
	}
	fmt.Println("Success Connect to Database!")
}
