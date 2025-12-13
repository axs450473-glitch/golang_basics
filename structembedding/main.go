package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	empId  string
	salary float64
}

func (p Person) introduce() {
	fmt.Printf("Hi I am %s and my age is %d", p.name, p.age)
}

// method overriding
func (p Employee) introduce() {
	fmt.Printf("Hi I am %s and my age is %d", p.name, p.age)
}

func main() {

	emp := Employee{
		Person: Person{
			name: "Abinash",
			age:  30,
		},
		empId:  "e001",
		salary: 52.00,
	}
	fmt.Println(emp.name) // directly accessing name field
	fmt.Println(emp.age)
	fmt.Println(emp.empId)
	fmt.Println(emp.salary)
	emp.introduce()
}
