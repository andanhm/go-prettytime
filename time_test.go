package prettytime_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/andanhm/go-prettytime"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{name: "Just now", t: time.Now(), want: "just now"},
		{name: "Second", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(1)),
			want: "1 second from now",
		},
		{name: "SecondAgo", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(-1)),
			want: "1 second ago"},
		{name: "Minutes", t: time.Now().Add(time.Hour*time.Duration(0) +
			time.Minute*time.Duration(59) +
			time.Second*time.Duration(59)), want: "60 minutes from now"},
		{name: "Tomorrow", t: time.Now().AddDate(0, 0, 1), want: "tomorrow"},
		{name: "Yesterday", t: time.Now().AddDate(0, 0, -1), want: "yesterday"},
		{name: "Week", t: time.Now().AddDate(0, 0, 7), want: "1 week from now"},
		{name: "WeekAgo", t: time.Now().AddDate(0, 0, -7), want: "1 week ago"},
		{name: "Month", t: time.Now().AddDate(0, 1, 0), want: "1 month from now"},
		{name: "MonthAgo", t: time.Now().AddDate(0, -1, 0), want: "1 month ago"},
		{name: "Year", t: time.Now().AddDate(50, 0, 0), want: "50 years from now"},
		{name: "YearAgo", t: time.Now().AddDate(-2, 0, 0), want: "2 years ago"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTimeSince := Format(tt.t); gotTimeSince != tt.want {
				t.Errorf("Format() = %v, want %v", gotTimeSince, tt.want)
			}
		})
	}
}

func ExampleFormat() {
	timeSlotes := []struct {
		name string
		t    time.Time
	}{
		{name: "Just now", t: time.Now()},
		{name: "Second", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(1)),
		},
		{name: "SecondAgo", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(-1)),
		},
		{name: "Minutes", t: time.Now().Add(time.Hour*time.Duration(0) +
			time.Minute*time.Duration(59) +
			time.Second*time.Duration(59))},
		{name: "Tomorrow", t: time.Now().AddDate(0, 0, 1)},
		{name: "Yesterday", t: time.Now().AddDate(0, 0, -1)},
		{name: "Week", t: time.Now().AddDate(0, 0, 7)},
		{name: "WeekAgo", t: time.Now().AddDate(0, 0, -7)},
		{name: "Month", t: time.Now().AddDate(0, 1, 0)},
		{name: "MonthAgo", t: time.Now().AddDate(0, -1, 0)},
		{name: "Year", t: time.Now().AddDate(2, 0, 0)},
		{name: "YearAgo", t: time.Now().AddDate(-2, 0, 0)},
	}

	for _, timeSlote := range timeSlotes {
		fmt.Printf("%s = %v\n", timeSlote.name, Format(timeSlote.t))
	}
}
