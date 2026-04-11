package main

import "fmt"

func main() {

	num := 76
	fmt.Printf("%05d\n", num) //00076

	message := "Hello"
	fmt.Printf("|%10s|\n", message)  // right aligned :|     Hello|
	fmt.Printf("|%-10s|\n", message) // left aligned : |Hello     |

	// String interpolation using backtick ``
	// used to represent regular expression
	// also used to represent SQL queries
	m1 := "Hello \nWorld!"
	m2 := `Hello \nWorld!`
	fmt.Println(m1) //
	fmt.Println(m2) //Hello \nWorld!

}
