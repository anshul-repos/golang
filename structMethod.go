package main

import "fmt"



type Person struct {
	FirstName, LastName, Email string
	Age                        int
}

func (e Person) PersonInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	return "--------END--------------"
}

func main() {

	e := Person{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
	}

	fmt.Println(e.PersonInfo())
}