package main

import "fmt"

func main() {
	// panic stops execution of a function
	// panic(interface{}) -> can use any value of any type as an argument
	process(-3)
}

func process(i int) {
	if i < 0 {
		// defer will execute even though our func is panicking
		//panic is executed after all defer are executed
		defer fmt.Println("inside panic condition 1")
		defer fmt.Println("inside panic condition 2")
		panic("input cannot be an negative number")

	}
}
