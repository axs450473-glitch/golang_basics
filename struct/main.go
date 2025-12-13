package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address // embedded struct
	Phone             // anonymous field
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// by design in go Structs & associated method should be declared at global scope
// i.e at pkg level
// methods are not declared inside Struct unlike classes

// pointer recievers allows methods to modify orginial struct instances
func (p *Person) incrementAge() {
	p.age++
}

type Address struct {
	city    string
	country string
}

type Phone struct {
	home string
	cell string
}

func main() {

	// Struct is a composite datatype
	// similar to classes in Java but doesnt follow inheritance property
	// ligtweight than class

	// initializing
	p := Person{
		firstName: "John",
		lastName:  "Wick",
		age:       100,
		address: Address{
			city:    "London",
			country: "UK",
		},
		Phone: Phone{
			home: "Nalbari",
			cell: "240150",
		},
	}
	fmt.Println(p)

	// anonymous struct
	user := struct {
		username string
		email    string
	}{
		username: "Aihsa",
		email:    "hsgdhsg@dgshd",
	}
	fmt.Println(user)

	fmt.Println(p.fullName())
	p.incrementAge()
	fmt.Println("AFTER INCREMENTING AGE BY 1", p)
}
