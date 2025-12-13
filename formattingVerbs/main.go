package main

import "fmt"

func main() {

	i := 64_36.9
	str := "Hello AB"

	fmt.Printf("%v\n", i)  //default format
	fmt.Printf("%#v\n", i) // display result in go syntax format
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	fmt.Printf("%v\n", str)  //default format
	fmt.Printf("%#v\n", str) // display result in go syntax format
	fmt.Printf("%T\n", str)

	int := 18
	fmt.Printf("%b\n", int)  //binary output
	fmt.Printf("%d\n", int)  // base 10 format
	fmt.Printf("%+d\n", int) // positive or negative
	fmt.Printf("%o\n", int)  // base 8 -- octa decimal represtn
	fmt.Printf("%O\n", int)
	fmt.Printf("%x\n", int) // hexa decimal
	fmt.Printf("%X\n", int)
	fmt.Printf("%#x\n", int) // hexadecimal with leading 0x notation
	fmt.Printf("%4d\n", int)
	fmt.Printf("%-4d\n", int)
	fmt.Printf("%04d\n", int)

	text := "hello"
	fmt.Printf("%s\n", text)
	fmt.Printf("%q\n", text)
	fmt.Printf("%8s\n", text)
	fmt.Printf("%-8s\n", text)
	fmt.Printf("%x\n", text)
	fmt.Printf("% x\n", text)

	t := true
	f := false
	fmt.Printf("%t\n", t)
	fmt.Printf("%v\n", t)
	fmt.Printf("%t\n", f)

	flt := 9188.89
	fmt.Printf("%e\n", flt)
	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
	fmt.Printf("%6.2f\n", flt)
	fmt.Printf("%g\n", flt)

}
