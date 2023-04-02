package main

import (
	"basics/channels/helpers"
	"log"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)

	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)

	// defer = Execute what comes after it once the function is finished
	defer close(intChan)

	// go = Starting a new routine with whatever comes after it (separate thread)
	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
