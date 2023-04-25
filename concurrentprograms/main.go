package main

import (
	"fmt"
	"math/rand"
	"time"
)

// You have two goroutines, one that generates random integers between 1 and 10,
// and another that calculates the square of the integer and prints it to the console.
// Implement the communication between the two goroutines using an unbuffered channel.

func main() {

	ch := make(chan int)

	go generateRand(ch)

	go squareInt(ch)

	time.Sleep(time.Second * 2)

}

func generateRand(ch chan int) {

	rand.Seed(time.Now().UnixNano())
	for {
		num := rand.Intn(10) + 1

		ch <- num
	}
}

func squareInt(ch chan int) {

	number := <-ch

	sqVal := number * number

	fmt.Println("sqVal", sqVal, "number", number)
}
