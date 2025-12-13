package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Cant be a negative number")
	}
	return 1, nil
}

// Custom errors

type myError struct {
	message string
}

//The error interface is defined simply as: type error interface { Error() string }. (builtin pkg in go)
/*

  By implementing the Error() string method from builtin pkg,
  the myError type (specifically, a pointer to it, *myError) is now considered an error
  and can be returned from functions that specify an error return type.

*/
func (e *myError) Error() string {
	return fmt.Sprintf("Error : %s", e.message)
}

func eprocess() error {
	return &myError{"Custom Error"}
}

// fmt.Errof()

func readConfig() error {
	return errors.New("Config Error")
}
func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func main() {

	fmt.Println(sqrt(-8.9))

	// e1 := eprocess()
	// if e1 != nil {
	// 	fmt.Println(e1)
	// 	return
	// }

	if e2 := readData(); e2 != nil {
		fmt.Println(e2)
		return
	}
	fmt.Println("Config Read successfully")
}
