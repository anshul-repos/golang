package main

import (
	"fmt"
	"strings"
)



func repeated(s string) map[string]int {

	input := strings.Fields(s)

	l := len(input)

	fmt.Println(l)

	wc := make(map[string]int)

	for _,word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}

	}

	return wc
}

func main(){

	str1 := "this is this word is this"

	for index, value := range repeated(str1){
		fmt.Println(index, "=", value)
	}

}

