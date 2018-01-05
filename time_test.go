package prettytime

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	var now = time.Now()

	t.Log(now.Format(time.RFC3339), "->", Format(now))

	Time := now.AddDate(-2, 0, 0)
	t.Log(Time.Format(time.RFC3339), "->", Format(Time))

	Time = now.Add(time.Hour*time.Duration(0) +
		time.Minute*time.Duration(59) +
		time.Second*time.Duration(59))
	t.Log(Time.Format(time.RFC3339), "->", Format(Time))

	Time = now.AddDate(50, 0, 0)
	t.Log(Time.Format(time.RFC3339), "->", Format(Time))
}

func ExampleFormat() {
	var now = time.Now()
	var timeFormat = time.RFC3339
	var interval = time.Duration(10)

	// 2018-01-04T17:07:46+05:30 -> just now
	fmt.Println(now.Format(timeFormat), "->", Format(now))

	// 2018-01-04T17:07:51+05:30 -> 5 seconds from now
	Time := now.Add(time.Hour*time.Duration(0) + time.Minute*time.Duration(0) + time.Second*5)
	fmt.Println(Time.Format(timeFormat), "->", Format(Time))

	// 2018-01-04T17:10:47+05:30 -> 3 minutes from now
	Time = now.Add(time.Hour*time.Duration(0) + time.Minute*time.Duration(3) + time.Second)
	fmt.Println(Time.Format(timeFormat), "->", Format(Time))

	// 2018-01-05T17:07:46+05:30 -> tomorrow
	Time = now.AddDate(0, 0, 1)
	fmt.Println(Time.Format(timeFormat), "->", Format(Time))

	// 2018-01-03T17:07:46+05:30 -> yesterday
	Time = now.AddDate(0, 0, -1)
	fmt.Println(Time.Format(timeFormat), "->", Format(Time))

	// 2018-01-11T17:07:46+05:30 -> 1 week from now
	Time = now.AddDate(0, 0, 7)
	fmt.Println(Time.Format(timeFormat), "->", Format(Time))

	// 2017-12-28T17:07:46+05:30 -> 1 week ago
	Time = now.AddDate(0, 0, -7)
	fmt.Println(Time.Format(timeFormat), "->", Format(Time))

	// 2018-02-04T17:07:46+05:30 -> 1 month from now
	Time = now.AddDate(0, 1, 0)
	fmt.Println(Time.Format(timeFormat), "->", Format(Time))

	// 2017-12-04T17:07:46+05:30 -> 1 month ago
	Time = now.AddDate(0, -1, 0)
	fmt.Println(Time.Format(timeFormat), "->", Format(Time))

	// Updates date every 5 sec
	go func() {
		for range time.Tick(time.Second * interval) {
			now.Add(time.Hour*time.Duration(0) + time.Minute*time.Duration(0) + time.Second*interval)
		}
	}()
	time.Sleep(time.Second * interval)
	for range time.Tick(time.Second * interval) {
		fmt.Println(now.Format(timeFormat), "->", Format(now))
	}
}
