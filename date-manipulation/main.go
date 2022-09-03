package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Just add the method you want to test here.
}

func sumAndSubtractTime() {
	toAdd := 1 * time.Hour
	fmt.Println("1:", toAdd)

	toAdd += 1 * time.Minute
	fmt.Println("2:", toAdd)

	toAdd += 1 * time.Second
	fmt.Println("3:", toAdd)

	toAdd -= 1*time.Minute + 1*time.Second
	fmt.Println("4:", toAdd)

	d1 := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
	fmt.Println("The time is", d1)

	add := 24 * time.Hour
	fmt.Println("Adding", add)

	nd1 := d1.Add(add)
	fmt.Println("The new time is", nd1)
}

func compareDates() {
	d1 := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
	fmt.Println("First time: ", d1)

	d2 := time.Date(2021, 12, 25, 16, 30, 55, 100, time.UTC)
	fmt.Println("Second time: ", d2)

	fmt.Println("The first time before second time? ", d1.Before(d2))
	fmt.Println("The first time after second time? ", d1.After(d2))

	fmt.Println("Second time before first time? ", d2.Before(d1))
	fmt.Println("Second time after first time? ", d2.After(d1))

	fmt.Println("Duration between first and second time is ", d1.Sub(d2))
	fmt.Println("Duration between second and first time is ", d2.Sub(d1))
}

func parseStringToDate() {
	ts := "2022-09-03vd 07:42:00"
	t, err := time.Parse("2006-01-02 15:04:05", ts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The time is: ", t)
	fmt.Println("RFC3339 Format: ", t.Format(time.RFC3339))
}

func formattingDate() {
	t := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	fmt.Println("The t is", t)
	fmt.Println(t.Format("2006-01-02'T'15:04:05 pm"))
	fmt.Println(t.Format(time.RFC3339))
}

func printNow() {
	currentTime := time.Now()
	printWithPrintf(currentTime)
	printEachFieldByTime(currentTime)
}

func printWithPrintf(time time.Time) {
	fmt.Printf("%d-%d-%d %d:%d:%d\n", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
}

func printEachFieldByTime(time time.Time) {
	fmt.Println("The time is", time)
	fmt.Println("The month is", time.Month())
	fmt.Println("The day is", time.Day())
	fmt.Println("The hour is", time.Hour())
	fmt.Println("The minute is", time.Minute())
	fmt.Println("The second is", time.Second())
}
