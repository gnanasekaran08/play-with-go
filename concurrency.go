package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---> START <---")
	go startOne()
	time.Sleep(1 * time.Second)
	go startTwo()
	time.Sleep(1 * time.Second)
	fmt.Println("Application started...")
}

func startOne() {
	fmt.Println("concurrency one...")
}

func startTwo() {
	fmt.Println("concurrency two...")
}
