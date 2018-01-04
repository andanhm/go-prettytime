## GoLang Pretty Time

> [go-prettytime] Format GoLang date time in a *`pretty`* way. ex : just now, a minute ago, 2 hours ago , 3 minutes ago

### Install

```bash
    go get github.com/andanhm/go-prettydate
```

### Using dependencies

#### [dep] GoLang Official Go dependency management tool
dep is a prototype dependency management tool for Go. It requires Go 1.8 or newer to compile.

```bash
    dep init
    dep ensure -add github.com/andanhm/go-prettytime
    #    Have any issue in vendoring try following command once 
    dep ensure -vendor-only
```

#### [gvt]
gvt is the go vendoring tool for the GO15VENDOREXPERIMENT, based on gb-vendor

```bash
    gvt fetch github.com/andanhm/go-prettytime
```


### Example

```go
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

```

#### Output

```
2018-01-04T17:07:46+05:30 -> just now
2018-01-04T17:07:51+05:30 -> 5 seconds from now
2018-01-04T17:10:47+05:30 -> 3 minutes from now
2018-01-05T17:07:46+05:30 -> tomorrow
2018-01-03T17:07:46+05:30 -> yesterday
2018-01-11T17:07:46+05:30 -> 1 week from now
2017-12-28T17:07:46+05:30 -> 1 week ago
2018-02-04T17:07:46+05:30 -> 1 month from now
2017-12-04T17:07:46+05:30 -> 1 month ago
2018-01-04T17:03:19+05:30 -> 20 seconds ago
2018-01-04T17:03:29+05:30 -> 30 seconds ago
2018-01-04T17:03:39+05:30 -> 40 seconds ago
2018-01-04T17:03:49+05:30 -> 50 seconds ago
2018-01-04T17:03:59+05:30 -> 1 minute ago
2018-01-04T17:04:09+05:30 -> 1 minute ago
2018-01-04T17:04:19+05:30 -> 1 minute ago
2018-01-04T17:04:29+05:30 -> 2 minute ago
```

> Inspired by the `John Resig` Pretty Date plugin for jQuery (http://ejohn.org/blog/javascript-pretty-date/)

### Contributions

Feel free to fork and add features, fix bugs and your pull request is more than welcome ❤

[dep]: <https://github.com/golang/dep>
[gvt]: <https://github.com/FiloSottile/gvt>
[go-prettytime]: <https://godoc.org/github.com/andanhm/go-prettytime>