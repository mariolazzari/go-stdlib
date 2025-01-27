package main

import (
	"fmt"
	"time"
)

func main() {
	start, err := time.Parse("15:04", "08:00")
	if err != nil {
		panic(err)
	}

	end, err := time.Parse("15:04", "09:30")
	if err != nil {
		panic(err)
	}

	duration := end.Sub(start)
	fmt.Println("Meeting duration:", duration)

}
