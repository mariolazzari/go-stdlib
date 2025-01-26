package main

import (
	"fmt"
	"time"
)

func main() {
	// get current time
	now := time.Now()
	fmt.Println("Now:", now)

	// get location
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
	fmt.Println("Location loaded:", location)

	pstTime := now.In(location)
	fmt.Printf("Now in %v: %v\n", location, pstTime)

	// add time
	duration := time.Hour + 30*time.Minute
	fmt.Println("Duration:", duration)
	futureTime := now.Add(duration)
	fmt.Println("Future time:", futureTime)

	// difference
	difference := futureTime.Sub(now)
	fmt.Println("Difference", difference)

	// dates
	year, month, day := now.Date()
	fmt.Println(year, month, day)

	// weekday
	weekday := now.Weekday()
	fmt.Println("Week day:", weekday)

	// add date
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Println("Tomorrow:", tomorrow)
}
