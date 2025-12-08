package main

import (
	"fmt"
	"log"
	"os/exec"
	"play-go/db"
	"time"
)

var x, y int = 10, 20

func main() {
	fmt.Println("---> Play With Go <---")
	fmt.Printf("x: %d, y: %d\n", x, y)

	cmd := exec.Command("bash", "-c", "echo 'Hello from Bash executed by Go!'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing command: %v\nOutput: %s", err, output)
	}
	fmt.Printf("Bash output: %s", output)

	_db := db.Connect()
	if _db != nil {
		defer _db.Close()
	}

	fmt.Println("Main function execution completed.")

	var users_count int
	err = _db.QueryRow("select count(*) from users").Scan(&users_count)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	fmt.Println("Total Users Count Query Executed:", users_count)
	// measure feedbacks count query
	start := time.Now()
	var entries_count int
	err = _db.QueryRow("select count(id) as entries_count from feedbacks where id > 0").Scan(&entries_count)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	fmt.Println("Total Feedbacks Count Query Executed:", entries_count)
	elapsed := time.Since(start)
	fmt.Printf("Feedbacks count query took %s\n", elapsed)

	roster_start := time.Now()
	var rosters_count int
	err = _db.QueryRow("select count(id) as rosters_count from rosters where id > 0").Scan(&rosters_count)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	fmt.Println("Total Rosters Count Query Executed:", rosters_count)
	roster_elapsed := time.Since(roster_start)
	fmt.Printf("Rosters count query took %s\n", roster_elapsed)

	db.Close()
}
