package main

import (
	"fmt"
	"math"
)

// want to be cnsidered a type of geometry type?
// you have to implement both the methods inside the interface
// supports polymorphism
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}
func (r rect) perim() float64 {
	return 2 * (r.height * r.width)
}

type circle struct {
	radius float64
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}
func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type is Integer")
	case string:
		fmt.Println("Type is String")
	default:
		fmt.Println("Invalid type selected...")
	}
}

func main() {
	r := rect{
		width:  3,
		height: 5,
	}
	c := circle{
		radius: 4,
	}
	measure(r)
	measure(c)
	myPrinter(10, 2, "dog", true)
	printType(9)
	printType("Abinash")
	printType(true)
}
