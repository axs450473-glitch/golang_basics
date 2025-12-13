package main

import "fmt"

func main() {

	// pointers are memory address of a variable , itprevents memory leak
	// limited to : referencing (&) & dereferencing(*)
	// go doesn't support pointer arithmatic like C or C++

	// pointer declaration
	var ptr *int
	var a int
	a = 10
	ptr = &a          // referencing
	fmt.Println(*ptr) //10 ~ dereferencing
	fmt.Println(ptr)  // 0xc00000a0d8 ~ memory address of a in RAM

	// ** A pointer that doesn't reference to any memory address is called Nil Pointer

	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int) {
	*ptr++

}
