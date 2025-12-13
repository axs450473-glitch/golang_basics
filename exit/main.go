package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Deferred statement")
	// os.Exit(codes) executes without any deferred function & stops execution of the program ith
	// right status code
	// bypasses defer,panic,recover mechanism
	fmt.Println("Start main")
	os.Exit(1) // unsuccessfull exit
	// if we use os.Exit(0) // succesfull exit
	// in both the case this print statement will not get executed
	fmt.Println("This line will never get executed")
}
