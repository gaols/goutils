# goutils

## Description

Some common used utilities to make your development more productive.

## utilities

### stringutils

if you come from Java, you should familiar with the those apis. It's the go implementation of
apache commons for string manipulation.

```go
import "github.com/gaols/goutils"

//Reverse a string
goutils.Reverse("fox") // out: xof
goutils.LeftPad("fox", 5, 'z') // out: zzfox
goutils.RightPad("fox", 5, 'z') // out: foxzz
goutils.IsBlank(" ") // out: true
goutils.IsBlank("a") // out: false
goutils.IsEmpty("")  // out: true
goutils.IsEmpty(" ") // out: false
```
