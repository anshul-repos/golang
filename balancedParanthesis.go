package main

import "fmt"

func main()  {

	x := isValid("{[()}")

	fmt.Println(x)

}


func isValid(s string) bool {

   var stack []rune

   for _, brac := range s {

		n := len(stack) - 1

		if brac == '}'{
			current := stack[n]
			stack = stack[:n]
			if current != '{'{
				return false
			}
		} else if brac == ')'{
			current := stack[n]
			stack = stack[:n]
			if current != '('{
				return false
			}
		} else if brac == ']'{
			 current := stack[n]
			 stack = stack[:n]
			if current != '[' {
				return false
			}
		} else {
			stack = append(stack, brac)
		}
   }
   if len(stack) == 0{
   	return true
   }
	return true
}