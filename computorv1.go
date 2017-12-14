package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

// Env : a structure to store members of the equation
type Env struct {
	left, right, reduced, a, b, c string
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

func reduceMember(s string) string {
	// len := len(s)
	ret := ""

	// fmt.Println(s)
	// re2 := regexp.MustCompile("X\\^2")
	// re1 := regexp.MustCompile("X\\^1")
	// for regexp.MustCompile("X\\^2") != nil {
	// 	fmt.Println(re.FindStringIndex(s))
	// loc := re.FindStringIndex(s)
	// fmt.Println(s[loc[0]:loc[1]])
	// }

	// reg := regexp.MustCompile("ab?")
	// fmt.Println(reg.FindStringIndex("tablett"))
	// fmt.Println(reg.FindStringIndex("foo") == nil)

	return ret
}

// func getReduced(e Env) {
// 	if utf8.RuneCountInString(e.right) != 1 && e.right[0] != '0' {
// 		e.right = reduceMember(e.right)
// 	}
// 	e.left = reduceMember(e.left)

// 	// e.reduced = e.reduced + e.left + e.right + "= 0"
// 	// strings.Contains(e.reduced, "^2")
// 	// fmt.Println("Reduced form:", e.reduced)
// }

func revertSigns(s string) string {
	var ret string

	if s[0] != '-' {
		ret += "-"
	}
	for _, v := range s {
		if v == '-' {
			ret += "+"
		} else if v == '+' {
			ret += "-"
		} else {
			ret += string(v)
		}
	}
	return ret
}

func calculate(e Env, s string) {

}

func removeUselessSigns(s string) string {
	var ret string
	len := utf8.RuneCountInString(s)

	i := 0
	if s[0] == '+' {
		i++
	}
	for ; i < len; i++ {
		a := s[i]
		if a == 'X' && s[i+1] == '^' && s[i+2] == '0' {
			i += 2
		} else if a == 'X' && s[i+1] == '^' && s[i+2] == '1' {
			ret += "X"
			i += 2
		} else if a == 'X' && s[i+1] == '^' && s[i+2] == '2' {
			ret += "X^2"
			i += 2
		} else if a != '*' {
			ret += string(a)
		}
	}
	return ret
}

func removeWhitespaces(s string) string {
	var ret string

	for _, v := range s {
		if v != ' ' && v != '\r' && v != '\n' && v != '\t' && v != '\f' && v != '\v' {
			ret += string(v)
		}
	}
	return ret
}

// (?m)([+-])(?:(?:(\d+(?:\.\d+)?)(?:(?:\*X(?:\^(\d+))?)|X(?:\^(\d+))?))|(\d+(?:\.\d+)?)|(?:(?:\*X(?:\^(\d+))?)|X(?:\^(\d+))?))
// SIZE FULL MATCH POUR PREMIER CHECK
// CHECK TEST.GO

func main() {
	if len(os.Args) > 2 {
		error(2)
	} else if len(os.Args) < 2 {
		error(3)
	}
	arg := removeWhitespaces(os.Args[1])
	str := strings.Split(arg, "=")
	if len(str) != 2 {
		fmt.Println("LEN")
		error(1)
	}
	if str[0][1] != '+' || str[0][1] != '-' { // ADD '+' if no sign at first occurence
		str[0] = "+" + str[0]
	}
	fmt.Println(str[0])
	fmt.Println(str[1])
	regex := `(?m)([+-])(?:(?:(\d+(?:\.\d+)?)(?:(?:\*X(?:\^(\d+))?)|X(?:\^(\d+))?))|(\d+(?:\.\d+)?)|(?:(?:\*X(?:\^(\d+))?)|X(?:\^(\d+))?))`
	matchLeft, _ := regexp.MatchString(regex, str[0])  // CHANGER FUNCTION
	matchRight, _ := regexp.MatchString(regex, str[1]) // SAME
	if !matchLeft || !matchRight {
		fmt.Println("MATCH")
		error(1)
	}
	e := Env{str[0], str[1], "", "", "", ""}
	e.left = removeUselessSigns(e.left)
	e.right = removeUselessSigns(e.right)
	fmt.Println("LEFT", e.left)
	fmt.Println("RIGHT", e.right)
	e.right = revertSigns(e.right)
	fmt.Println("REVERSE", e.right)
	calculate(e, e.right)
	calculate(e, e.left)
	e.reduced = e.a + e.b + e.c
	// fmt.Println("Reduced form:", e.reduced)
}
