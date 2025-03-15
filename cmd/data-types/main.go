package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
    var intNum int = 10
	fmt.Println(intNum)

	var floatNum float64 = 9
	fmt.Println(floatNum)

	// Casting int to float if the types are different
	var result float64 = float64(intNum) + floatNum
	fmt.Println(result)

	// This should return 1 because the result is an int
	var result2 int = intNum / int(floatNum)
	fmt.Println(result2)

	var myString string = "Hello World"
	var myString2 string = `This is a 
	multiline string`

	fmt.Println(myString + " " + myString2)

	// To get the length of a string, use utf8.RuneCountInString
	fmt.Println(utf8.RuneCountInString(myString))
	fmt.Println(utf8.RuneCountInString(myString2))

	// In go, there is rune type
	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = true
	fmt.Println(myBoolean)

	// In go, the default value of a variable is 0 for all int, float, rune, "" for string, and false for boolean.
	var myInt int 
	var myString3 string
	var myBoolean2 bool
	fmt.Println(myInt, myString3, myBoolean2)

	// In go, shorthand declaration is also possible
	myInt2 := 10
	fmt.Println(myInt2)

	// Multiple shorthand declaration is also possible
	myInt3, myString4 := 10, "Hello"
	fmt.Println(myInt3, myString4)

	// Constants in go
	const myConst int = 10
	fmt.Println(myConst)
	
	// This will throw an error because constants are immutable
	// myConst = 20 


}
