package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	// This is the same as above
	// p2 := new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	// Change the value of the pointer p points to
	*p = 42
	// This reassigns the address of p to the address of i
	p = &i
	*p = 27
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	// This is an array of 5 floats
	x := [5]float64{1, 2, 3, 4, 5}
	fmt.Println(square(&x))
	fmt.Println(x)
}

func square(x *[5]float64) [5]float64 {
	for i := range(x) {
		x[i] = x[i] * x[i]
	}
	return *x
}