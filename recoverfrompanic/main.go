package main

import "fmt"

func main() {

	// recover func is used to manage behaviour of panicking goroutines
	// to avoid abrupt panicking
	// recover allows the code to keep on executing after panicking
	// recover func is called inside defer func
	process(-6)

}
func process(i int) {
	//anonymous function
	// as soon as there is a panic , recover() returns the argument passed to the panic function(recover message)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered :", r)
		}
	}()
	if i < 0 {
		panic("Semething Went Wrong")
	}
	fmt.Println("START PROCESS")

}
