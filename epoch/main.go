package main

import (
	"fmt"
	"time"
)

func main() {

	// 00:00:00 UTC on Jan 1 , 1970
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println(unixTime) //epoch time

	t := time.Unix(unixTime, 0)
	fmt.Println(t) // human readable time format
	fmt.Println(t.Format("2006-01-02"))
}
