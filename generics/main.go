package main

import "fmt"

/*
func swap[T any](a, b T) (T, T) {
	return b, a
}
*/

// code reusability, type safety, performance
type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {

	s.elements = append(s.elements, element)

}

func (s *Stack[T]) pop() (T, bool) {

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s Stack[T]) printAll() {
	fmt.Println(s.elements)
}
func main() {

	// x, y := 1, 2
	// fmt.Println(swap(x, y))
	// x1, y1 := "Abinash", "Sarma"
	// fmt.Println(swap(x1, y1))

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(4)
	intStack.push(5)
	intStack.push(8)
	intStack.push(7)
	intStack.push(4)
	intStack.pop()
	intStack.printAll()

	intStack1 := Stack[string]{}
	intStack1.push("tr")
	intStack1.push("uy")
	intStack1.push("gg")
	intStack1.pop()
	intStack1.printAll()
}
