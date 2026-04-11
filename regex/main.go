package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("He said, \"I am good\"") //He said, "I am good"
	fmt.Println(`He said, "I am good"`)   // He said, "I am good"

	//compile a regex to match email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`) // \. is writtent together outside square braces
	// + refers to one or more occurences
	// {2,} refers to 2 or more occurences
	email1 := "user@email.com"
	email2 := "invalid_email"

	fmt.Println("Email1 : ", re.MatchString(email1))
	fmt.Println("Email2 : ", re.MatchString(email2))

	// Capturing Groups
	// Compile regex to capture date components
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	date := "2026-01-03"
	//find all submatches
	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	// Target String
	str := "Hello World"
	re = regexp.MustCompile(`[aeiou]`)
	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	// i = case sensitive
	// m = multiline model
	// s = dot match all

	re = regexp.MustCompile(`(?i)go`) // always strt ? with flags
	test1 := "Hello Golang"
	fmt.Println(re.MatchString(test1)) // true
}
