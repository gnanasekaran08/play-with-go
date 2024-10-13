package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("---> Play With Time <---")
	// location, _ := time.LoadLocation("Asia/Kolkata")
	// timeLayout := "2022-01-01 15:04:05"
	// read only date and time

	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)

	log.Println("Current Time :", time.Now().Local().Format("2006-01-02 15:04:05"))
	now := time.Now()
	fmt.Println("Current Time :", now)
	timer := "23:12"
	parseTime, _ := time.Parse("15:04", timer)

	scheduleTime := time.Date(now.Year(), now.Month(), now.Day(), parseTime.Hour(), parseTime.Minute(), 0, 0, time.Local)
	fmt.Println("Schedule Time :=", scheduleTime)

	sss := scheduleTime.Add(time.Hour * 4)
	fmt.Println("sss :=", sss)
}
