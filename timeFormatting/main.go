package main

import (
	"fmt"
	"time"
)

func main() {

	// Mon Jan 2 15:04:05 MST 2006 <- reference time go uses

	layout := "2006-01-02T15:04:05Z07:00"
	str := "2026-01-04T15:05:00Z"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)

	str1 := "Jan 04,2026 03:10 PM"
	layout1 := "Jan 02,2006 03:04 PM"
	t1, _ := time.Parse(layout1, str1)
	fmt.Println(t1)

}
