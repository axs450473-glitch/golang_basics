package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

// NOTE: we use value receiver when there is no change or modification involved to receiver instance

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// NOTE: we use pointer receiver when there is  change or modification involved to receiver instance
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

//methods can be associated with  any type
type MyInt int

func (m MyInt) IsPositive() bool {
	return m > 0
}
func (MyInt) welcomeMessage() string {
	return "Welcome Home"
}

type Shape struct {
	Rectangle
}

func main() {
	// rect is the receiver instance
	rect := Rectangle{
		length: 8,
		width:  7,
	}
	area := rect.Area()
	fmt.Println("BEFORE MODIFCATION", area)
	rect.Scale(2.5)
	area = rect.Area()
	fmt.Println("AFTER MODIFCATION", area)

	myInt := MyInt(-5)
	fmt.Println(myInt.IsPositive())
	fmt.Println(myInt.welcomeMessage())

	s := Shape{
		Rectangle: Rectangle{
			length: 10,
			width:  4,
		},
	}
	fmt.Println(s.Area())

}
