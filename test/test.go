package main

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode/utf8"
)

func main() {
	var re = regexp.MustCompile(`(?m)([+-])(?:(?:(\d+(?:\.\d+)?)(?:(?:\*X(?:\^(?:[+-])?(\d+))?)|X(?:\^(?:[+-])?(\d+))?))|(\d+(?:\.\d+)?)|(?:(?:\*X(?:\^(?:[+-])?(\d+))?)|X(?:\^(?:[+-])?(\d+))?))`)
	var str = `+2*X^2+2-5+2-1*X^1`
	len := utf8.RuneCountInString(str)
	count := 0
	polynom := 0

	for _, v := range str {
		if v == 'X' {
			polynom = 1
		}
	}

	for i, match := range re.FindAllStringSubmatch(str, -1) {
		fmt.Println(match, "found at index", i)
		count += utf8.RuneCountInString(match[0])
		if polynom > 0 {
			p := 0
			for j := 7; j > 0; j-- {
				if match[j] != "" {
					p, _ = strconv.Atoi(match[j])
					break
				}
			}
			if p > polynom {
				polynom = p
			}
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
		fmt.Println("FORMAT OK")
	}
}
