package main

import "fmt"

func main() {
	// adder() is called only once
	// so i is initialized to 10 only once

	val := adder()

	fmt.Println(val()) //11
	fmt.Println(val()) //12
	fmt.Println(val()) //13
	fmt.Println(val()) //14

	val2 := adder()
	fmt.Println(val2())

	subtraction := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtraction(1)) //98
	fmt.Println(subtraction(1)) //97
	fmt.Println(subtraction(1)) //96
}
func adder() func() int {
	i := 10
	fmt.Println("Value of i before incrementing is :", i)
	return func() int {
		i++
		fmt.Println("Value of i inside anonymous func :", i)
		return i
	}
}
