package main

import "fmt"

func reverseArray(revArr []int) []int {

	for i , j := 0 , len(revArr) - 1 ; i < j ; i,j = i+1, j-1 {

		revArr[i] , revArr[j] = revArr[j], revArr[i]
	}

	return revArr

}

func reverseString(str string) string  {

	strByte := []rune(str)
	for i,j := 0, len(strByte) - 1 ; i<j; i,j=1+1,j-1{
		strByte[i],strByte[j] = strByte[j],strByte[i]
	}

	return string(strByte)
}

func main(){
	fmt.Println(reverseArray([]int{1,2,3,4,5,6,7,8,9}))
	fmt.Println(reverseString("ABCD"))

}
