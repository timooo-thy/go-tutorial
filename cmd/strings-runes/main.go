package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "rεsumε"
	// Unicode points number which represent the characters
	myRunes := []rune(myString)

	for i, v := range myString {
		fmt.Printf("The index is %v and the value is %v\n", i, string(v))
	}

	for i, v := range myRunes {
		fmt.Printf("The index is %v and the value is %v\n", i, string(v))
	}

	var strSlice = []string{"cat", "dog", "fish"}
	var catStr = ""

	// Concatenating strings but this is not the best way to do it
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Println(catStr)

	// Instead, use strings.builder
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	fmt.Println(strBuilder.String())
}