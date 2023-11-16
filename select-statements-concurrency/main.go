package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// select statements
	// a select statement chooses which of a set of possible
	// send or receive operations will proceeed
	// it looks similar to a switch statement buyt with the cases
	// all reffering to communcation operations

	// concurrency channels
	ch1 := make(chan int)
	ch2 := make(chan int)

	// int63 n milisecond interval
	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	// go routines
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 41
	}()

	// select statement
	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}

}
