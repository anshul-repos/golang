package main

import "fmt"




type Stack []string

func (s *Stack) Push(str string ) {
	*s = append(*s, str)
}


func(s *Stack) Pop()(string){

	index := len(*s) - 1

	element := (*s)[index]

	*s = (*s)[:index]

return element

}


func main() {


	var stack Stack

	stack.Push("Eelemnt 1")
	stack.Push("Eelemnt 2")

	for len(stack) > 0 {

		x := stack.Pop()
		fmt.Println(x)

	}

}