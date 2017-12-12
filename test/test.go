package main

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode/utf8"
)

func main() {
	var re = regexp.MustCompile(`(?m)([+-])(?:(?:(\d+(?:\.\d+)?)(?:(?:\*X(?:\^(\d+))?)|X(?:\^(\d+))?))|(\d+(?:\.\d+)?)|(?:(?:\*X(?:\^(\d+))?)|X(?:\^(\d+))?))`)
	var str = `+2*X^2+2-5+2-1*X^1`
	len := utf8.RuneCountInString(str)
	count := 0
	polynom := 0

	for i := 0; i < len; i++ {
		if str[i] == 'X' {
			polynom = 1
		}
	}

	for i, match := range re.FindAllStringSubmatch(str, -1) {
		fmt.Println(match, "found at index", i)
		// s := match[1] + match[2]
		// fmt.Println("STR=", s)
		count += utf8.RuneCountInString(match[0])
		p := 0
		for j := 7; j > 0; j-- {
			if match[j] != "" {
				p, _ = strconv.Atoi(match[j])
				break
			}
		}
		fmt.Println("P = ", p)
		if polynom > 0 && p > polynom {
			polynom = p
		}
	}
	msg := "Polynomial degree:"
	if polynom > 2 {
		msg = "Cant solve. " + msg
	}
	fmt.Println(msg, polynom)
	if count != len {
		fmt.Println("COUNT=", count, "len=", len)
	} else {
		fmt.Println("OK")
	}
}
