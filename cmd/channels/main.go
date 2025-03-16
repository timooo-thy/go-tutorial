package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 4
var MAX_TOFU_PRICE float32 = 2


func main() {
	// Channel
	// var c = make(chan int)
	// Buffer channel
	var c = make(chan int, 5)

	// c1 := make(chan int)
	// go process(c)
	// fmt.Println(<-c)

	// Loops with channel
	go process(c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	chickenChannel := make(chan string)
	tofuChannel := make(chan string)
	websites := []string{"walmart.com", "costco.com", "wholefoods.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}

	sendMessage(chickenChannel, tofuChannel)
}

func process(c chan int) {
	// Tells other processes using this channel that the process is done
	// Defer means calling this right before the function exits
	defer close(c)
	for i := range 5 {
		c <- i
	}
	fmt.Println("Exiting Process...")
}

func checkChickenPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second+1)
		chicken_price := rand.Float32() * 10
		if chicken_price < MAX_CHICKEN_PRICE {
			c <- website
			break
		}
	}
}

func checkTofuPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second+1)
		tofu_price := rand.Float32() * 20
		if tofu_price < MAX_TOFU_PRICE {
			c <- website
			break
		}
	}
}

func sendMessage(chickenC chan string, tofuC chan string) {
	select {
		case website := <- chickenC:
			fmt.Printf("\nText Sent: Found deal on chicken at %v", website)
		case website := <- tofuC:
			fmt.Printf("\nEmail Sent: Found deal on tofu at %v", website)
	}
}