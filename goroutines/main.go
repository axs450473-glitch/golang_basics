package main

import (
	"fmt"
	"time"
)

// go routines are functions with go keyword
// its taken out from main thread & run in bg
// they come back to join the main() after completing its execution
// Goroutines donot stop the programme flow & are non blocking
func main() {
	fmt.Println("Before Goroutine exec.")
	go sayHello() // leaves the main thread
	fmt.Println("After Goroutine exec.")
	time.Sleep(2 * time.Second)
	/*
	   Before Goroutine exec.
	   After Goroutine exec.
	   Say Hello Goroutines

	*/

	// all the goroutines run concurrently
	// Goroutine scheduling in go
	/*
		     -> Managed by go runtime schedular
			 -> Uses M:N scheduling model 
			   - M goroutines are mapped to N OS threads .
			 -> Efficient Multiplexing
			   - means switching - schedular switches goroutines to available os threads
			   - many goroutines on limited number of os threads

	*/
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Say Hello Goroutines")
}


