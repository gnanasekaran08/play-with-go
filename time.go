package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---> Play With Time <---")
	// location, _ := time.LoadLocation("Asia/Kolkata")
	// timeLayout := "2022-01-01 15:04:05"

	now := time.Now()
	fmt.Println("Current Time :", now)
	timer := "23:12"
	parseTime, _ := time.Parse("15:04", timer)

	scheduleTime := time.Date(now.Year(), now.Month(), now.Day(), parseTime.Hour(), parseTime.Minute(), 0, 0, time.Local)
	fmt.Println("Schedule Time :=", scheduleTime)

	sss := scheduleTime.Add(time.Hour * 4)
	fmt.Println("sss :=", sss)
}
