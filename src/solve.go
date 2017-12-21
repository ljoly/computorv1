/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   solve.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: ljoly <ljoly@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/12/21 14:19:40 by ljoly             #+#    #+#             */
/*   Updated: 2017/12/21 14:19:40 by ljoly            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func solveDegreeTwo(e Env) {
	delta := e.b*e.b - 4*e.a*e.c
	fmt.Println("\nΔ =", delta)
	if delta < 0 {
		fmt.Println("Δ is negative. There is no real solution, but two complex solutions:")
		if e.b != 0 {
			fmt.Printf("x1 = (%g - i√%g) / %g\n", -e.b, -delta, 2*e.a)
			fmt.Printf("x2 = (%g + i√%g) / %g\n", -e.b, -delta, 2*e.a)
		} else {
			fmt.Printf("x1 = -i√%g / %g\n", -delta, 2*e.a)
			fmt.Printf("x2 = i√%g / %g\n", -delta, 2*e.a)
		}
	} else if delta == 0 {
		sol := -e.b / (2 * e.a)
		fmt.Println("Δ is nul. The solution is:", sol)
	} else if delta > 0 {
		fmt.Println("Δ is positive. There are two real solutions:")
		if e.b != 0 {
			fmt.Printf("x1 = (%g - √%g) / %g\n", -e.b, delta, 2*e.a)
			fmt.Printf("x2 = (%g + √%g) / %g\n", -e.b, delta, 2*e.a)
		} else {
			fmt.Printf("x1 = -√%g / %g\n", delta, 2*e.a)
			fmt.Printf("x2 = √%g / %g\n", delta, 2*e.a)
		}
	}
}

func solveDegreeOne(e Env) {
	sol := -e.c / e.b
	fmt.Println("\nThe solution is:")
	fmt.Println("x =", sol)
}

func solve(e *Env) {
	if e.a == 0 && e.b == 0 && e.c == 0 {
		fmt.Println("Polynomial equality. All the real numbers are solutions.")
	} else {
		if e.signA == '-' {
			e.a *= -1
		}
		if e.signB == '-' {
			e.b *= -1
		}
		if e.signC == '-' {
			e.c *= -1
		}
		if e.degree == 1 {
			solveDegreeOne(*e)
		} else if e.degree == 2 {
			solveDegreeTwo(*e)
		}
	}
}
