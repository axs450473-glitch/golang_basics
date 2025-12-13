package main

import "fmt"

func init() {
	fmt.Println("Initializing pkg 1")
}
func init() {
	fmt.Println("Initializing pkg 2")
}
func init() {
	fmt.Println("Initializing pkg 3")
}
func main() {

	// go execute init func automatically before initializing the package
	// before main func is executed
	// init funcs will be executed sequentially as the order its written
	fmt.Println("Inside Main")
}
