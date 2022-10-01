package main

import (
	"math/rand"
	"time"
)


const numPol = 10
func RandomNumbers(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)

	return value

}

func CalculateValue(intChan chan int)  {
	randomNumber := RandomNumbers(numPol)

	intChan <- randomNumber
	
}

func main()  {

	intChan := make(chan int)

	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan

	println(num)
}