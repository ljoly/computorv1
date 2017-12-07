package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Env : a structure to store members of the equation
type Env struct {
	left    string
	right   string
	reduced string
}

func error(e int) {
	if e == 3 {
		println("Error: No argument")
	} else if e == 2 {
		println("Error: Too many arguments")
	} else if e == 1 {
		println("Error: Wrong format")
	}
	os.Exit(-1)
}

func countChar(s string, c byte) int {
	len := len(s)
	m := 0
	for i := 0; i < len; i++ {
		if s[i] == c {
			m++
		}
	}
	return m
}

func reduceMember(s string) string {
	// len := len(s)
	ret := ""

	fmt.Println(s)
	re := regexp.MustCompile("\\^2")
	fmt.Println(re.FindStringIndex(s))
	loc := re.FindStringIndex(s)
	fmt.Println(s[loc[0]:loc[1]])

	// reg := regexp.MustCompile("ab?")
	// fmt.Println(reg.FindStringIndex("tablett"))
	// fmt.Println(reg.FindStringIndex("foo") == nil)

	return ret
}

func getReduced(e Env) {
	e.left = reduceMember(e.left)
	// e.right = reduceMember(e.right)

	// e.reduced = e.reduced + e.left + e.right + "= 0"
	// strings.Contains(e.reduced, "^2")
	fmt.Println("Reduced form:", e.reduced)
}

func removeWhitespaces(s string) string {
	var ret string
	len := len(s)
	for i := 0; i < len; i++ {
		a := s[i]
		if a == 'X' && i+2 < len && s[i+1] == '^' && s[i+2] == '0' {
			ret = ret + "X"
			i += 2
		} else if a != ' ' && a != '\r' && a != '\n' && a != '\t' && a != '\f' && a != '\v' {
			ret = ret + string(s[i])
		}
	}
	return ret
}

func main() {
	if len(os.Args) == 3 {
		error(3)
	}
	if len(os.Args) > 2 {
		error(2)
	}
	str := strings.Split(os.Args[1], "=")
	if len(str) != 2 {
		println("EQUALS")
		error(1)
	}
	regex := `^((\s*[+-]?\s*[1-9][0-9]{0,9}\s*\*\s*X\s*(|(\^\s*[0-2]))\s*)+)$`
	matchLeft, _ := regexp.MatchString(regex, str[0])
	matchRight, _ := regexp.MatchString(regex, str[1])
	if !matchLeft || !matchRight {
		println("FORMAT")
		error(1)
	}
	e := Env{removeWhitespaces(str[0]), removeWhitespaces(str[1]), ""}
	getReduced(e)
}
