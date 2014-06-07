Go Regexp Common Patterns
========

A collection of common regex patters for Go

### Usage

```
import (
	"fmt"
	"github.com/jsimnz/regexcom"
)

func main() {
	url := "http://github.com"
	if regexcom.Url.Match(url) {
		fmt.Println("Corret URL!")
	} else {
		fmt.Println("Sorry, invalid URL!")
	}
}

// The overall pattern for the regex common package
// is to call Match(s string) on any given pattern
// Ex
//    regexcom.<Pattern>.Match(s string) bool
//
// Look below for available patterns

```

### Patterns

```
Numeric 				- 	Numeric values.
Alpha 					-	Characters and letters (upper and lower case).
AlphaNumeric 			- 	Numbers and letters (upper and lower case).
AlphaCapsOnly			-	Upper case characters only.
AlphaNumericCapsOnly	-	Numeric values and upper case only characters.
Url 					-	Match a url patter. Can match http, https, ftp, ftps, no-scheme, no www. (See test).
Email 					-	Match an email pattern, user@domain.tld.
HashtagHex				-	Match hexidecimal prefixed with '#' (IE #afe), commonly used for colors.
ZeroXHex				-	Match C syntax 0x hexidecimal typing (IE 0xAF, only upper characters or only lower).
IPv4					-	Match an IPv4 address, (IE. 192.168.1.1)
IPv6					-	Match an IPv6 address, 

More to come...
```

### Create your own pattern

Each pattern is defined on a simple Pattern struct, and uses the standard syntax regexcom.<Pattern>.Match(string) bool
```
YouPattern := regexcom.NewPattern("regex-pattern")
// Then you can use it as normal
YouPattern.Match(string)
```

### Contribute
- Fork It
- Add your own Pattern in top 'var' block
- Create pull request
- Be happy! :smile:

### Author
- (jsimnz)[https://github.com/jsimnz]

### LICENCE
The MIT License (MIT)

Copyright (c) 2014 John-Alan Simmons

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
