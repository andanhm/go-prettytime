# Go Pretty Time

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/andanhm/go-prettytime)
[![Build Status](https://travis-ci.org/andanhm/go-prettytime.svg?branch=master)](https://travis-ci.org/andanhm/go-prettytime)
[![GoDoc](https://camo.githubusercontent.com/3de3bba30c9355c0d919804e7b31e6b504af74e2/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f6e617468616e792f6c6f6f7065723f7374617475732e737667)](https://godoc.org/github.com/andanhm/go-prettytime)
[![CodeCov](https://codecov.io/gh/icza/minquery/branch/master/graph/badge.svg)](https://codecov.io/gh/andanhm/go-prettytime)

Format Go date time in a *`pretty`* way. ex : just now, a minute ago, 2 hours ago , 3 minutes ago

Inspired by the `John Resig` Pretty Date [plug-in] for JQuery

## Install

```bash
go get github.com/andanhm/go-prettytime
```

### Example

```go
package main

import (
    "fmt"
    "time"

    prettyTime "github.com/andanhm/go-prettytime"
)

func main() {
    timeSlots := []struct {
        name string
        t    time.Time
    }{
		{name: "Just now", t: time.Now()},
		{name: "Second", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(1)),
		},
		{name: "Second Ago", t: time.Now().Add(
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
		{name: "Week Ago", t: time.Now().AddDate(0, 0, -7)},
		{name: "Month", t: time.Now().AddDate(0, 1, 0)},
		{name: "Month Ago", t: time.Now().AddDate(0, -1, 0)},
		{name: "Year", t: time.Now().AddDate(2, 0, 0)},
		{name: "Year Ago", t: time.Now().AddDate(-2, 0, 0)},
	}

	for _, timeSlot := range timeSlots {
        fmt.Printf("%s = %v\n", timeSlot.name, prettyTime.Format(timeSlot.t))
	}
}

```

### Contributions

Feel free to fork and add features, fix bugs and your pull request is more than welcome ❤

[dep]: <https://github.com/golang/dep>
[gvt]: <https://github.com/FiloSottile/gvt>
[go-prettytime]: <https://godoc.org/github.com/andanhm/go-prettytime>
[plug-in]: <http://ejohn.org/blog/javascript-pretty-date/>