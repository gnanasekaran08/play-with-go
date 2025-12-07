package main

import (
	"fmt"
	"play-go/db"
	"time"
)

func main() {
	fmt.Println("---> Play With Go <---")
	_db := db.Connect()
	if _db != nil {
		defer _db.Close()
	}

	fmt.Println("Main function execution completed.")

	var users_count int
	err := _db.QueryRow("select count(*) from users").Scan(&users_count)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	fmt.Println("Total Users Count Query Executed:", users_count)
	// measure feedbacks count query
	start := time.Now()
	var entries_count int
	err = _db.QueryRow("select count(id) from feedbacks").Scan(&entries_count)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	fmt.Println("Total Feedbacks Count Query Executed:", entries_count)
	elapsed := time.Since(start)
	fmt.Printf("Feedbacks count query took %s\n", elapsed)
}
