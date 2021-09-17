package main

import "fmt"

func main(){

	arr := []int{1, 3, 2, 4, 5}


	insertionsort(arr)

	fmt.Println("\n--- Sorted ---\n\n", arr, "\n")

}

func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}


