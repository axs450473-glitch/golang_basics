package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())

	// specific time
	specificTime := time.Date(2026, time.July, 3, 12, 0, 0, 0, time.UTC)
	fmt.Println(specificTime)

	//parse time
	parsedTime, _ := time.Parse("2006-01-02", "2026-01-03") //Mon Jan 2 2006 15:04:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02", "26-01-03")    //Mon Jan 2 2006 15:04:05 MST 2006
	parsedTime2, _ := time.Parse("6-1-2", "26-1-3")         //Mon Jan 2 2006 15:04:05 MST 2006  ->  layout didnt match
	// time.Parse("6-1-2", "06-1-3")  ** correct layout match
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)

	//formatting time
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 04:15"))

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	fmt.Println("Rounded Time : ", t.Round(time.Hour))

	// loc, _ := time.LoadLocation("Asia/Kolkata")
	// t = time.Date(2026, time.January, 3, 14, 34, 7, 9, time.UTC)

	// //Convert this ti specific time zone
	// tLocal := t.In(loc)

	// //Perform rounding
	// roundedTime := t.Round(time.Hour)
	// localRoundedTime := roundedTime.In(loc)

	// fmt.Println("Original Time (UTC): ", t)
	// fmt.Println("Original Time (Local): ", tLocal)
	// fmt.Println("Rounded time : ", roundedTime)
	// fmt.Println("Local Rounded Time : ", localRoundedTime)

	fmt.Println("Truncated Time : ", t.Truncate(time.Hour))

	loc, _ := time.LoadLocation("America/New_York")
	// convert time to location
	tInNy := time.Now().In(loc)
	fmt.Println(tInNy)

	t1 := time.Date(2026, time.January, 3, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2026, time.February, 3, 12, 0, 0, 0, time.UTC)
	fmt.Println(t2.Sub(t1))
	fmt.Println("t2 is after t1? :", t2.After(t1))

}
