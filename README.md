# Go Pretty Time

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/andanhm/go-prettytime)
[![Build Status](https://travis-ci.org/andanhm/go-prettytime.svg?branch=master)](https://travis-ci.org/andanhm/go-prettytime)
[![GoDoc](https://camo.githubusercontent.com/3de3bba30c9355c0d919804e7b31e6b504af74e2/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f6e617468616e792f6c6f6f7065723f7374617475732e737667)](https://godoc.org/github.com/andanhm/go-prettytime)
[![CodeCov](https://codecov.io/gh/icza/minquery/branch/master/graph/badge.svg)](https://codecov.io/gh/andanhm/go-prettytime)

Format Go date time in a _`pretty`_ way. ex : just now, a minute ago, 2 hours ago , 3 minutes ago

Inspired by the `John Resig` Pretty Date [plug-in] for JQuery

```
prettytime.Format("2008-01-28T20:24:17Z") // => "2 hours ago"
prettytime.Format("2008-01-27T22:24:17Z") // => "Yesterday"
prettytime.Format("2008-01-26T22:24:17Z") // => "2 days ago"
prettytime.Format("2008-01-14T22:24:17Z") // => "2 weeks ago"
```

## Install

```bash
go get github.com/andanhm/go-prettytime
```

### Example

```go
package main

import (
	"log"
	"time"

	"github.com/andanhm/go-prettytime"
)

const (
	layout = "2006-01-02T15:04:05Z"
)

func main() {
	t, err := time.Parse(layout, "2008-01-28T20:24:17Z")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s \n", prettytime.Format(t))
}
// Output: 13 years ago
```

### i18n support

`go-prettytime` uses [go-i18n](https://github.com/nicksnyder/go-i18n) to provide i18n capabilities.

#### Translating a new language
Create an empty message file for the language that you want to add (e.g. translate.es.toml).

Run `goi18n merge en.toml translate.es.toml` to populate translate.es.toml with the messages to be translated.

After `translate.es.toml` has been translated, rename it to `es.toml`

Load `es.toml` into the `time.go` file via

```go
bundle.MustLoadMessageFile("es.toml")
```


### Contributions

Feel free to fork and add features, fix bugs and your pull request is more than welcome ❤

[go-prettytime]: https://pkg.go.dev/github.com/andanhm/go-prettytime
[plug-in]: http://ejohn.org/blog/javascript-pretty-date/
