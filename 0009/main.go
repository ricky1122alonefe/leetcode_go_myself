package main

import (
	"strconv"
	"fmt"
)

func main(){
	fmt.Println(isPalindrome(123321))
}

func isPalindrome(x int) bool {

	// 将int 转换成字符串
	s:=strconv.Itoa(x)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}