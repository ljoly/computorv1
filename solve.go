package main

import "fmt"

func solveDegreeTwo(e Env) {
	delta := e.b*e.b - 4*e.a*e.c
	fmt.Println("\nΔ =", delta)
	if delta < 0 {
		fmt.Println("Δ is negative. There is no real solution.")
	} else if delta == 0 {
		sol := -e.b / 2 * e.a
		fmt.Println("Δ is nul. The solution is:", sol)
	} else if delta > 0 {
		fmt.Println("Δ is positive. There are to real solutions:")
		fmt.Printf("x1 = (%g - √%g) / %g\n", -e.b, delta, 2*e.a)
		fmt.Printf("x2 = (%g + √%g) / %g\n", -e.b, delta, 2*e.a)
	}
}

func solveDegreeOne(e Env) {
	sol := -e.c / e.b
	fmt.Println("The solution is:", sol)
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
