package main

import "fmt"

func main() {

  n := 5

  for i:=1;i <= n ; i++ {
  	for j:=1; j <= n-i ; j++{

  		fmt.Print(" ")
	}
	k := 0
	for {
		fmt.Print("*")
		k++
		if k == 2*i-1 {
			break
		}
	}
	fmt.Println()
  }


	//for i := 1; i <= n; i++ { //cols
	//
	//	for j := 1; j <= n-i; j++ { //rows
	//		fmt.Print("  ")
	//	}
	//	k:=0
	//	for {
	//		fmt.Print("* ")
	//		k++
	//		if k == 2*i-1 {
	//			break
	//		}
	//	}
	//	fmt.Println()
	//}

}