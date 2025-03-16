package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
	owner
	// ownerInfo owner
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
	owner
}

type owner struct {
	name string
}

// This is an interface. It is a contract that says that any type that has a method called milesLeft() that returns a uint8 can be considered an engine
type engine interface {
	milesLeft() uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

// Here we can generalise the function to take any type that implements the engine interface
func canMakeit(e engine, miles uint8) {
	if miles < e.milesLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You can't make it!")
	}
}

func main() {
	var myCar gasEngine = gasEngine{20, 10, owner{"John Doe"}}
	myCar.gallons = 18
	fmt.Println(myCar.gallons, myCar.mpg, myCar.name)

	// This is an anonymous struct and it is not reusable
	var myCar2 = struct {
		mpg uint8
		gallons uint8
		owner
	}{20, 10, owner{"Jane Doe"}}
	fmt.Println(myCar2.gallons, myCar2.mpg, myCar2.name)
	fmt.Println(myCar.milesLeft())
	canMakeit(myCar, 2)
	canMakeit(myCar, 200)

	// This is an example of composition
	var myCar3 electricEngine = electricEngine{4, 10, owner{"Jane Doe"}}
	fmt.Println(myCar3.milesLeft())
	canMakeit(myCar3, 2)
	canMakeit(myCar3, 200)
}