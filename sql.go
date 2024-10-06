package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Setting struct {
	id         int64
	org_id     int64
	key_name   string
	key_value  string
	created_at string
	updated_at string
}

const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "reallist"
)

func main() {
	fmt.Println("---> SQL <---")

	fmt.Println("Username..." + username)
	fmt.Println("Password..." + password)
	fmt.Println("Hostname..." + hostname)
	fmt.Println("DB Name..." + dbname)

	db, err := connectDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to the database")

	records, nil := db.Query("SELECT * FROM settings")
	if err != nil {
		fmt.Println("Error executing query", err)
		return
	}
	fmt.Println("Query executed")
	// fmt.Println("Total records count: ", len(records))
	for records.Next() {
		setting := Setting{}
		err := records.Scan(&setting.id, &setting.org_id, &setting.key_name, &setting.key_value, &setting.created_at, &setting.updated_at)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Setting: ", setting)
	}

	defer db.Close()
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname))
	if err != nil {
		return nil, err
	}

	return db, nil
}
