package main

import "fmt"

func main()  {

	a := []int32{5,6,7}

	b := []int32{3,6,10}

	fmt.Println(compareTriplets(a,b))
}


func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here

	var alice int32
	var bob int32
	alice = 0
	bob = 0

	if a[0] > b[0] {
		alice++
	} else if a[0] < b[0]{
		bob++
	}
	if a[1] > b[1] {
		alice++
	} else if a[1] < b[1]{
		bob++
	}
	if a[2] > b[2] {
		alice++
	} else if a[2] < b[2]{
		bob++
	}

	var res []int32

	res = append(res,alice)
	res = append(res,bob)

	return res
}
