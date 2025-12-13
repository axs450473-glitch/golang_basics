package main

import (
	"errors"
	"fmt"
)

type customErrors struct {
	code    int
	message string
	err     error
}

// Error returns the error message implementing Error() method of error interface

func (e *customErrors) Error() string {
	return fmt.Sprintf("Error %d : %s %v", e.code, e.message, e.err)
}
func process() error {
	return &customErrors{code: 400, message: "Bad Request"}
}

// wrapped error
func doSomething() error {
	return errors.New("internal now")
}
func doSomethingElse() error {
	e := doSomething()
	if e != nil {
		return &customErrors{
			code:    500,
			message: "Internal server error",
			err:     e,
		}
	}
	return nil
}

func main() {
	// err := process()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	err1 := doSomethingElse()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
}
