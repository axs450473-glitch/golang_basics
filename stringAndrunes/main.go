package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// string is sequence of byte(uint8 value)
	// immutable - once created , its value cant be changed
	message := "Hello Golang developers"
	message1 := "Hello, \r Golang!"
	rawMsg := `Hello\nGolang\ndevelopers` // discard escaped characters
	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(rawMsg)

	// array of unicode characters or runes
	fmt.Println("Length of message variable is :", len(message)) //23
	fmt.Println("Length of message variable is :", len(rawMsg))  //25

	for _, char := range message {
		fmt.Print(char, "\t")    // gives ascii value of the each characters // prints ascii values
		fmt.Printf("%X\t", char) // prints hexadecimal values

	}
	fmt.Println()
	// lexicographical comparison
	str1 := "apple"
	str2 := "Apple"
	str3 := "app"
	fmt.Println(str1 > str2) // A has ASCII value of 65 & a has 97
	fmt.Println(str3 < str1)
	fmt.Println("Rune count : ", utf8.RuneCountInString("HELLO"))

	//immutability ~ cant append more runes in any existing string
	str4 := "Macbook"
	newStr := str1 + " " + str4
	fmt.Println(newStr)

	// rune is an int value (int32)
	var ch rune = 'a'
	jch := 'ら'
	fmt.Println(ch)
	fmt.Println(jch)
	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)
	// convert rune to string
	cstr := string(ch)
	fmt.Printf("%T is the type of cstr\n", cstr)

}
