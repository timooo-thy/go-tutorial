package main

import (
	"errors"
	"fmt"
)

func main() {
	val := "Hello World"
	printMe(val)

	numerator := 10
	denominator := 2
	quotient, remainder, err := intDivision(numerator, denominator)

	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v\n", quotient)
	} else {
		fmt.Printf("The result of the integer division is %v and the remainder is %v\n", quotient, remainder)
	}

	switch {
		case err != nil:
			fmt.Println(err.Error())
		case remainder == 0:
			fmt.Printf("The result of the integer division is %v\n", quotient)
		default:
			fmt.Printf("The result of the integer division is %v and the remainder is %v\n", quotient, remainder)
	}

	switch remainder {
		case 0:
			fmt.Printf("The result of the integer division is %v\n", quotient)
		default:
			fmt.Printf("The result of the integer division is %v and the remainder is %v\n", quotient, remainder)
	}

	// Comparision operators

	// fmt.Println(1 == 1)
	// fmt.Println(1 != 1)
	// fmt.Println(1 < 1)
	// fmt.Println(1 <= 1)
	// fmt.Println(1 > 1)
	// fmt.Println(1 >= 1)
	// if 1 == 1 && 2 == 2 {
	// 	fmt.Println("Both conditions are true")
	// }
	// if 1 == 1 || 2 == 2 {
	// 	fmt.Println("At least one condition is true")
	// }
}

func printMe(val string) {
	fmt.Println(val)
}

func intDivision(a int, b int) (int, int, error) {
	err := error(nil)

	if b == 0 {
		err = errors.New("division by zero is not allowed")
		return 0, 0, err
	}

	return a / b, a % b, err
}