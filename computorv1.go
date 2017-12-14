package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func reduce(e Env, s string) {
	var re = regexp.MustCompile(formatRegex)
	var degreeTwo = regexp.MustCompile(degreeRegex)
	// var x = regexp.MustCompile(`(?m)(?:X)`)
	c := ""

	for _, group := range re.FindAllStringSubmatch(s, -1) {
		if degreeTwo.FindAllStringSubmatch(group[0], -1) != nil {
			if group[2] == "" {
				c = "1"
			} else {
				c = group[2]
			}
			if _, err := strconv.ParseFloat(group[1]+c, 64); err != nil {
				error(coefNotInt, e, group[0])
			} else {
				coef, _ := strconv.ParseFloat(group[1]+c, 64)
				e.a += coef
			}
			// } else if x.FindAllStringSubmatch(s, -1) != nil {

			// } else {

		}
	}
}

func getPolynom(s string, e *Env) int64 {
	var re = regexp.MustCompile(formatRegex)
	var reNegativeDegree = regexp.MustCompile(negativeDegreeRegex)
	var exp = regexp.MustCompile(degreeRegex)
	var polynom int64
	len := utf8.RuneCountInString(s)
	count := 0

	for _, v := range s {
		if v == 'X' {
			polynom = 1
			break
		}
	}

	if reNegativeDegree.FindAllStringSubmatch(s, -1) != nil {
		error(degreeNegative, *e, s)
		return -1
	}

	for _, match := range re.FindAllStringSubmatch(s, -1) {
		count += utf8.RuneCountInString(match[0])
		if polynom > 0 {
			var p int64
			for _, v := range match[0] {
				if v == 'X' && exp.FindAllStringSubmatch(match[0], -1) != nil {
					for j := 7; j > 0; j-- {
						if match[j] != "" {
							if _, err := strconv.ParseInt(match[j], 10, 64); err != nil {
								error(degreeNotInt, *e, match[0])
							} else {
								p, _ = strconv.ParseInt(match[j], 10, 64)
								break
							}
						}
					}
				} else if v == 'X' {
					p = 1
				}
				if p > 0 {
					break
				}
			}

			if p > polynom {
				polynom = p
			}
		}
	}
	if count != len {
		e := Env{}
		error(wrongFormat, e, "")
	}
	return polynom
}

func main() {
	e := Env{}
	if len(os.Args) > 2 {
		error(tooManyArguments, e, "")
	} else if len(os.Args) < 2 {
		error(noArgument, e, "")
	}
	arg := removeWhitespaces(os.Args[1])
	fmt.Println("AAAAAAARRRRRGGGGG:", arg)
	str := strings.Split(arg, "=")
	if len(str) != 2 {
		error(wrongFormat, e, "")
	}
	e.left = addFirstSign(str[0])
	e.right = addFirstSign(str[1])
	pLeft := getPolynom(e.left, &e)
	pRight := getPolynom(e.right, &e)
	if pLeft < 0 || pRight < 0 {
		e.degree = -1
	} else {
		if pLeft > pRight {
			e.degree = pLeft
		} else {
			e.degree = pRight
		}
	}
	if e.degree > 2 {
		error(degreeTooHigh, e, "")
	} else if e.degree < 0 {
		error(degreeNegative, e, "")
	}
	e.left = removeUselessSigns(e.left)
	e.right = removeUselessSigns(e.right)
	fmt.Println("\nLEFT\n", e.left)
	fmt.Println("\nRIGHT\n", e.right)
	e.right = revertSigns(e.right)
	fmt.Println("\nREVERSE RIGHT\n", e.right)
	fmt.Println("RIGHT + LEFT", e.left+e.right)
	// reduce(e, e.right)
	// reduce(e, e.left)
	fmt.Printf("\nReduced form: %f + %f + %f\n", e.a, e.b, e.c)
	fmt.Println("Polynomial degree:", e.degree)
}
