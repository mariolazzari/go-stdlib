package main

import (
	"fmt"
	"time"
)

func main() {
	// timer
	timer := time.NewTimer(5 * time.Second)
	<-timer.C

	// ticker
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	const max = 3
	for t := range ticker.C {
		fmt.Println("Tick at:", t)
		i++
		if i == max {
			break
		}
	}

	// sleep
	time.Sleep(5 * time.Second)

	// measuring time
	start := time.Now()
	time.Sleep(1 * time.Second)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed)

}
