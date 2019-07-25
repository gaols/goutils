# goutils

## Description

Some common utilities to make your development more productive.
It's self contained, no other dependencies is required.

## Install

```shell
go get github.com/gaols/goutils
```

## utilities

### stringutils

if you come from Java, you should familiar with the those apis. It's the go implementation of
apache commons for string manipulation.

```go
import "github.com/gaols/goutils"

goutils.Reverse("fox")                // out: xof
goutils.LeftPad("fox", 5, 'z')        // out: zzfox
goutils.RightPad("fox", 5, 'z')       // out: foxzz
goutils.IsBlank(" ")                  // out: true
goutils.IsBlank("a")                  // out: false
goutils.IsEmpty("")                   // out: true
goutils.IsEmpty(" ")                  // out: false
goutils.DefaultIfBlank(" ", "hello")  // out: hello
goutils.DefaultIfBlank("c", "hello")  // out: c
goutils.DefaultIfEmpty("", "hello")   // out: hello
goutils.DefaultIfEmpty("c", "hello")  // out: c
goutils.Substring("hello", 0, 1)      // out: h
goutils.IsAnyBlank("hello", "")       // out: true
goutils.IsAnyBlank("hello")           // out: false
```

### dateutils

date utils has a lot of date related sugar methods.

```go
goutils.Today()
goutils.Yesterday()
goutils.BeginningOfThisWeek()
goutils.BeginningOfThisMonth()
goutils.BeginningOfThisYear()
goutils.LastWeek()
goutils.LastMonth()
goutils.LastYear()
// ... and many others
```

### timeutils

```go
// assumes time now is: 2017-12-27 21:15:15
goutils.FmtTime(time.Now(), "-datetime")      // out: 2017-12-27 21:15:15
goutils.FmtTime(time.Now(), "-datetime-")     // out: 2017-12-27 21:15
goutils.FmtTime(time.Now(), "-datetime--")    // out: 2017-12-27 21
goutils.FmtTime(time.Now(), "-date")          // out: 2017-12-27
goutils.FmtTime(time.Now(), "time")           // out: 21:15:15
goutils.FmtTime(time.Now(), "time-")          // out: 21:15
goutils.FmtTime(time.Now(), "time--")         // out: 21

goutils.FmtTime(time.Now(), "/datetime")      // out: 2017/12/27 21:15:15
goutils.FmtTime(time.Now(), "/datetime-")     // out: 2017/12/27 21:15
goutils.FmtTime(time.Now(), "/datetime--")    // out: 2017/12/27 21
goutils.FmtTime(time.Now(), "/date")          // out: 2017/12/27

goutils.FmtTime(time.Now(), ".datetime")      // out: 2017.12.27 21:15:15
goutils.FmtTime(time.Now(), ".datetime-")     // out: 2017.12.27 21:15
goutils.FmtTime(time.Now(), ".datetime--")    // out: 2017.12.27 21
goutils.FmtTime(time.Now(), ".date")          // out: 2017.12.27

goutils.FmtTime(time.Now(), "xzzz")           // error
goutils.MustFmtTime(time.Now(), "xzzz")       // panic
goutils.MustParseTime("2017-12-27", "-date")  // time object@2017-12-27
```
