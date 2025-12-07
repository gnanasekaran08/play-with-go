package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() *sql.DB {
	fmt.Println("Connecting to the database...")
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = ""
	config.Net = "tcp"
	config.Addr = "127.0.0.1:3306"
	config.DBName = "reallist_one"

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Database ping failed:", err)
		return nil
	}

	fmt.Println("Database connection established successfully!")
	return db
}
