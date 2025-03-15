package main

import "fmt"

func main() {
	// Default value is 0 for all elements
	var intArr [3]int = [3]int{1, 2, 3}
	// intArr := [3]int{1, 2, 3}
	// intArr := [...]int{1, 2, 3} // The compiler will count the number of elements for you
	intArr[2] = 10
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:])

	// Slices are like dynamic arrays
	intSlice := []int{1, 2, 3}
	fmt.Printf("The length is %v and the capacity is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 4)
	fmt.Printf("The length is %v and the capacity is %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	// Appending arrays
	intSlice2 := []int{5, 6, 7}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	
	// Making slices with make, second argument is the length and the third is the capacity
	intSlice3 := make([]int, 3, 8)
	fmt.Println(intSlice3)

	// Maps
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	fmt.Println(myMap["one"])
	// If accessing a key that doesn't exist, it will return the default value of the value type
	fmt.Println(myMap["three"])

	age, ok := myMap["three"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("The key doesn't exist")
	}

	// Deleting a key
	// delete(myMap, "one")
	// fmt.Println(myMap)

	// Iterating over a map
	for key, val := range myMap {
		fmt.Printf("The key is %v and the value is %v\n", key, val)
	}

	// Iterating over a slice
	for index, val := range intSlice {
		fmt.Printf("The index is %v and the value is %v\n", index, val)
	}

	// For loops
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }
	
	// Modernised for loop
	for i := range 5 {
		fmt.Println(i)
	}

	// While loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	
}