package main

import (
	"fmt"
	"math"
)

func gradingStudents(grades []int32) []int32 {
	// Write your code here

	var ng []int32

	for _, i := range grades {

		ii := float64(i)
		res := math.Ceil((ii+5/2)/5) * 5
		r := int32(res)
		dif := r-i

		if r < 38 {
			ng = append(ng,i)
		} else {
			if r > 38 && dif < 3 {
				ng = append(ng,r)
			} else {
				ng = append(ng,i)
			}

		}
	}
	return ng
}


func main(){

	arr := []int32{75,67,38,33}

	gradingStudents(arr)

	fmt.Println(gradingStudents(arr))

}