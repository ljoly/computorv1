/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   error.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: ljoly <ljoly@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/12/21 14:19:31 by ljoly             #+#    #+#             */
/*   Updated: 2017/12/21 14:19:32 by ljoly            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
)

func error(e ErrorMsg, env Env, exp string) {
	if e == noArgument {
		fmt.Println("Error: No argument")
	} else if e == tooManyArguments {
		fmt.Println("Error: Too many arguments")
	} else if e == wrongFormat {
		fmt.Println("Error: Wrong format")
	} else if e == degreeTooHigh {
		fmt.Printf("Error: Polynomial degree: \033[1;31m%d\033[0m\n", env.degree)
		fmt.Println("Degree too high. Can't solve.")
	} else if e == degreeNegative {
		fmt.Printf("Error: in \033[1;31m%s\033[0m\n", exp)
		fmt.Println("Polynomial degree negative. Can't solve.")
	} else if e == degreeNotInt {
		fmt.Printf("Error: in \033[1;31m%s\033[0m\n", exp)
		fmt.Println("Degree is not an Int. Can't solve.")
	} else if e == falseExpression {
		fmt.Println("Error: false expression.")
	} else if e == isEquality {
		fmt.Println("Error: expression is a simple equality.")
	}
	os.Exit(0)
}
