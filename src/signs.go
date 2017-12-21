/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   signs.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: ljoly <ljoly@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/12/21 14:19:35 by ljoly             #+#    #+#             */
/*   Updated: 2017/12/21 14:19:36 by ljoly            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"unicode/utf8"
)

func storeSign(e *Env) {
	if e.a < 0 {
		e.a *= -1
		e.signA = '-'
	} else {
		e.signA = '+'
	}

	if e.b < 0 {
		e.b *= -1
		e.signB = '-'
	} else {
		e.signB = '+'
	}

	if e.c < 0 {
		e.c *= -1
		e.signC = '-'
	} else {
		e.signC = '+'
	}
}

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

	for i := 0; i < len; i++ {
		if s[i] == 'X' && i+1 < len && s[i+1] == '^' && s[i+2] == '0' {
			if s[i-1] == '+' || s[i-1] == '-' {
				ret += "1"
			}
			i += 2
		} else if s[i] == 'X' && i+1 < len && s[i+1] == '^' && s[i+2] == '1' {
			ret += "X"
			i += 2
		} else if s[i] == 'X' && i+1 < len && s[i+1] == '^' && s[i+2] == '2' {
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
