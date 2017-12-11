package main

import (
	"regexp"
	"fmt"
)

func main() {
	var re = regexp.MustCompile(`(?m)([+-])?(?:(?:(\d+(?:\.\d+)?)(?:(?:\*X(?:\^(\d+))?)|X(?:\^(\d+))?))|(\d+(?:\.\d+)?)|(?:(?:\*X(?:\^(\d+))?)|X(?:\^(\d+))?))`)
	var str = `-0.999999X+98479*X^2-5*X^0+4*X^1+8*X^1-1X^0+X^0+2X-5+2+2.5`

	for i, match := range re.FindAllStringSubmatch(str, -1) {
		fmt.Println(match, "found at index", i)
	}
}

