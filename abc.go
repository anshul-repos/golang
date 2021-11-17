package main

import "fmt"

func main(){
	var n int
	//fmt.Scan(&n)

	n = 7

	for i:=0;i<n;i++{
		for j:=0;j<n-i-1;j++{
			fmt.Print(" ")
		}
		for j:=0;j<=i;j++{
			fmt.Print("#")
		}
		fmt.Println()
	}
}