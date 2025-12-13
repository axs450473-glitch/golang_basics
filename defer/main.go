package main

import "fmt"

func main() {

	//defer is a mechanism that postpones execn of a function
	// until sorrounding function returns
	process(10)

}

func process(i int) {
	// follows LIFO - Last In First Out
	//** defer arguments will be evaluated as soon as it is encountered .. Doesnt mean executed last means evealuated last
	defer fmt.Println("Defer Value of i :", i)
	defer fmt.Println("Defer statement executed 1")
	defer fmt.Println("Defer statement executed 2")
	defer fmt.Println("Defer statement executed 3")
	i++
	fmt.Println("Actual value of i :", i)
	fmt.Println("Actual Statement")
}
