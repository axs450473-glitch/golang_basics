package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "Hello Go!"

	for i, v := range str {
		fmt.Printf("Value at index %d is %c", i, v)
		fmt.Println()
	}
	fmt.Println(str[:5]) //hello

	//string conversion
	num := 18
	str1 := strconv.Itoa(num)
	fmt.Println(len(str1)) //2

	//Split
	names := "Abinash, Simi"
	str2 := strings.Split(names, ",")
	fmt.Println(str2) // [Abinash  Simi]

	//Join : concatenate elements of a slice into a single string using separators
	fruits := []string{"apple", "orange", "banana"}
	str3 := strings.Join(fruits, "-")
	fmt.Println(str3) // apple-orange-banana

	// if string contains subset of characters
	fmt.Println(strings.Contains(str, "Go")) //true

	//Replace : 1 represent replacing 1 occurence of the substring
	str4 := strings.Replace(str, "Go", "Abinash", 1)
	fmt.Println(str4) //Hello Abinash!

	// TrimSpace : remove leading & tailing space
	str5 := " Hello Man "
	fmt.Println(strings.TrimSpace(str5)) //Hello Man

	// Repeat s string
	fmt.Println(strings.Repeat("Haka", 6))

	// count a substring inside a string
	fmt.Println(strings.Count("Hello", "l")) //2

	//check prefix & suffix in a string
	fmt.Println(strings.HasPrefix("Hello", "He"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	str6 := "Hello 123 Go 656"
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re.FindAllString(str6, -1)) //[123 656]   -1 means search for all the matches & store in a slice

	str7 := "名前はなんですか"
	fmt.Println(utf8.RuneCountInString(str7))

	// STRING BUILDER  := for memory efficiency to avoid unnecessary memory allocations

	var builder strings.Builder
	// Write some string
	// since string is immutable so we can keep on building the string
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world")

	// make the final string
	// convert builder to a string
	result := builder.String()
	fmt.Println(result) // Hello, world

	// using Writerune to add character
	builder.WriteRune(',')
	builder.WriteString("axs")
	result = builder.String()
	fmt.Println(result) // Hello, world,axs

	// to start with new string
	builder.Reset()
	builder.WriteString("No No No")
	result = builder.String()
	fmt.Println(result) //No No No

}
