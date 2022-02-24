package main

import (
	"fmt"
	"sync"
)


var wg sync.WaitGroup

func runner1()  {
	defer wg.Done()
	fmt.Println("I am first")
}
func runner2()  {
	defer wg.Done()
	fmt.Println("I am second")
}


func execute(){
	wg.Add(2)
	go runner1()
	go runner2()

	wg.Wait()
}


func main() {
	execute()
}
