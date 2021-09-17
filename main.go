package main

import (
	"fmt"
)

func main(){



	ar1 := []int{12,16,45,90,82}
	ar2 := []int{12,16,90,82}



	x := missingElement(ar1,ar2)

	fmt.Println(x)


}


func missingElement(a, b []int) ([]int){

	var x,y []int

	for j,_ := range a {
		for k,_ := range b {
			if b[k] != a[k] {
				x = append(x, a[j])
			} else {
				y = append(y, b[k])
			}
		}
	}
	return x
}
