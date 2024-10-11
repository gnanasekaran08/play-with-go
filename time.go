package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---> Play With Time <---")
	// timeLayout := "2022-01-01 15:04:05"
	currentTime := time.Now()
	fmt.Println("Current Time :", currentTime)

	timer := "18:00"

	fmt.Println("Timer :", timer)

	fmt.Println("In String Format :=", currentTime.String())

	// fmt.Println("Scheduled Time :=", scheduledTime)

}
