package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2025, time.April, 10, 15, 22, 10, 15, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	// Parse string to time
	formatter := "2006-01-02 15:04:05" // format type see docs
	value := "2025-05-10 13:34:10"
	valueTime, err := time.Parse(formatter, value)
	if err == nil {
		fmt.Println(valueTime)
		fmt.Println(valueTime.Hour())
		fmt.Println(valueTime.Year())
	} else {
		fmt.Println("Error", err.Error())
	}

	// Duration
	var duration1 time.Duration = time.Second * 100 // 100 seconds
	var duration2 time.Duration = time.Millisecond * 10 // 10 millis
	var duration3 time.Duration = duration1 - duration2
	fmt.Println(duration3)
	fmt.Printf("duration: %d\n", duration3) // in nano seconds
}