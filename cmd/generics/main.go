package main

import "fmt"

type gasEngine struct {
	gallons float32
	mpg float32
}

type electricEngine struct {
	kwh float32
	mpkwh float32
}

type car [T gasEngine | electricEngine] struct {
	carMake string
	carModel string
	engine T
}

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice(intSlice))
	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice(float32Slice))

	gasCar := car[gasEngine]{
		carMake: "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 20,
			mpg: 25,
		},
	}

	electricCar := car[electricEngine]{
		carMake: "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh: 50,
			mpkwh: 25,
		},
	}

	fmt.Printf("Electric car mileage is: %v\n", electricCar.engine.kwh * electricCar.engine.mpkwh)
	fmt.Printf("Gas car mileage is: %v\n", gasCar.engine.gallons * gasCar.engine.mpg)

}

func sumSlice[T int | float32 |float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}

	return sum
}
