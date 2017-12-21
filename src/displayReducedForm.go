/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   displayReducedForm.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: ljoly <ljoly@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/12/21 14:19:23 by ljoly             #+#    #+#             */
/*   Updated: 2017/12/21 14:19:24 by ljoly            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func displayReduced(e Env) {
	if e.a == 0 && e.b == 0 {
		if e.c != 0 {
			error(falseExpression, e, "")
		} else if !e.hasX {
			error(isEquality, e, "")
		}
	}
	if e.degree == 2 {
		if e.signA == '-' {
			fmt.Printf("Reduced form: %c %g\033[1;32mX^2\033[0m %c %g\033[1;32mX\033[0m %c %g = 0\n", e.signA, e.a, e.signB, e.b, e.signC, e.c)
		} else {
			fmt.Printf("Reduced form: %g\033[1;32mX^2\033[0m %c %g\033[1;32mX\033[0m %c %g = 0\n", e.a, e.signB, e.b, e.signC, e.c)
		}
	} else if e.degree == 1 {
		if e.signB == '-' {
			fmt.Printf("Reduced form: %c %g\033[1;32mX\033[0m %c %g = 0\n", e.signB, e.b, e.signC, e.c)
		} else {
			fmt.Printf("Reduced form: %g\033[1;32mX\033[0m %c %g = 0\n", e.b, e.signC, e.c)
		}
	}
}
