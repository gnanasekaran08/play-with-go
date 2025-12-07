package main

import (
	"fmt"
	"play-go/db"
)

func main() {
	fmt.Println("---> Play With Go <---")
	db.Connect()
}
