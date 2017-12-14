package main

import "unicode/utf8"

func revertSigns(s string) string {
	var ret string

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

func removeUselessSigns(s string) string {
	var ret string
	len := utf8.RuneCountInString(s)

	i := 0
	if s[0] == '+' {
		i++
	}
	for ; i < len; i++ {
		if s[i] == 'X' && s[i+1] == '^' && s[i+2] == '0' {
			if s[i-1] == '+' || s[i-1] == '-' {
				ret += "1"
			}
			i += 2
		} else if s[i] == 'X' && s[i+1] == '^' && s[i+2] == '1' {
			ret += "X"
			i += 2
		} else if s[i] == 'X' && s[i+1] == '^' && s[i+2] == '2' {
			ret += "X^2"
			i += 2
		} else if s[i] != '*' {
			ret += string(s[i])
		}
	}
	return ret
}

func addFirstSign(s string) string {
	ret := s

	if s[0] != '+' && s[0] != '-' {
		ret = "+" + s
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
