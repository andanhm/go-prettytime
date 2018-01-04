package main

import (
	"fmt"
	"time"

	prettyTime "go-prettytime"
)

var now = time.Now()
var timeFormat = time.RFC3339
var interval  = time.Duration(10)

func main() {
	fmt.Println(now.Format(timeFormat), "->", prettyTime.Format(now))

	Time := now.Add(time.Hour*time.Duration(0) + time.Minute*time.Duration(0) + time.Second*5)
	fmt.Println(Time.Format(timeFormat), "->", prettyTime.Format(Time))

	Time = now.Add(time.Hour*time.Duration(0) + time.Minute*time.Duration(3) + time.Second)
	fmt.Println(Time.Format(timeFormat), "->", prettyTime.Format(Time))

	Time = now.AddDate(0, 0, 1)
	fmt.Println(Time.Format(timeFormat), "->", prettyTime.Format(Time))

	Time = now.AddDate(0, 0, -1)
	fmt.Println(Time.Format(timeFormat), "->", prettyTime.Format(Time))

	Time = now.AddDate(0, 0, 7)
	fmt.Println(Time.Format(timeFormat), "->", prettyTime.Format(Time))

	Time = now.AddDate(0, 0, -7)
	fmt.Println(Time.Format(timeFormat), "->", prettyTime.Format(Time))

	Time = now.AddDate(0, 1, 0)
	fmt.Println(Time.Format(timeFormat), "->", prettyTime.Format(Time))

	Time = now.AddDate(0, -1, 0)
	fmt.Println(Time.Format(timeFormat), "->", prettyTime.Format(Time))

	go func() {
		for range time.Tick(time.Second * interval) {
			now.Add(time.Hour*time.Duration(0) + time.Minute*time.Duration(0) + time.Second*interval)
		}
	}()
	time.Sleep(time.Second * interval)
	for range time.Tick(time.Second * interval) {
		fmt.Println(time.Now().Format(timeFormat), "->",prettyTime.Format(now))
	}
}
