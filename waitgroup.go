package main

import (
	"fmt"
	"sync"
)


var wg sync.WaitGroup

func responseFunc(str string){

	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()

	fmt.Println("Step1 : ", str)
	fmt.Println("Step2 : ", str )
	fmt.Println("Step3 : ", str )

}

func main() {

	str1 := "Hello"
	str2 := "Hello there"
	str3 := "Hello End"

	// Add a count of three, one for each goroutine.
	wg.Add(3)

	go responseFunc(str1)
	go responseFunc(str2)
	go responseFunc(str3)


	// Wait for the goroutines to finish.
	wg.Wait()
	//time.Sleep(time.Second*10)

}
