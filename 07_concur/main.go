package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var waitGroup sync.WaitGroup
	// add 2 routines to wait group
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
	}()

	go func() {
		defer waitGroup.Done()
	}()

	// wait for both func
	waitGroup.Wait()

	var counter atomic.Int64
	counter.Add(1)
	val := counter.Load()
	counter.Store(100)

	fmt.Println(val)

}
