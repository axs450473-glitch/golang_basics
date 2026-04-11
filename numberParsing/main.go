package main

import (
	"fmt"
	"strconv"
)

func main() {

	// parsing refers to converting textual representation to respctive type

	numString := "12345"
	res, _ := strconv.Atoi(numString)
	fmt.Println(res)

	res1, _ := strconv.ParseInt(numString, 10, 32)
	fmt.Println(res1)

	floatStr := "8.54"
	res2, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Println(res2)

	binaryString := "10101"
	decimal, _ := strconv.ParseInt(binaryString, 2, 64)
	fmt.Println(decimal)
}
