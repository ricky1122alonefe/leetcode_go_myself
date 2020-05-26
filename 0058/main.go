package main

import "fmt"

func main(){
	fmt.Println(reverseLeftWords("abcdefg",2))
}
func reverseLeftWords(s string, n int) string {
	// 题目方面做了限制
	if len(s) ==0 || len(s)>10000 || n ==0{
		return s
	}
	strTail:=s[n:]
	strHead :=s[0:n]
	return strTail+strHead
}