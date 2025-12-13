package main

import "fmt"

func main() {

	// Printing Functions -> Print(), Printf(). Println()
	// Formatting Functions -> Sprint(),Sprintf(),Sprintln()
	// Scanning functions -> Scan(), Scanf(), Scanln()
	// Error Formatting Function -> Errorf()

	s := fmt.Sprint("HelloWorld", 123, 456)
	fmt.Print(s)
	y := fmt.Sprintln("Hello World", 123, 456) // with newline character
	fmt.Print(y)
	fmt.Print(y)
	age := 30
	x := fmt.Sprintf("Ram age is %d", age)
	fmt.Println(x)

	var name string
	var ag int
	fmt.Print("Enter your name & age :")
	//fmt.Scan(&name, &ag)
	//fmt.Scanln(&name, &ag)
	fmt.Scanf("%s %d", &name, &ag)
	fmt.Printf("The name is %s and age is %d", name, ag)
	err := checkAge(ag)
	if err != nil {
		fmt.Println("Error :", err)
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to derive", age)
	}
	return nil
}
