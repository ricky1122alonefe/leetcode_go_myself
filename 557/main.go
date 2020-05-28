package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
func reverseWords(s string) string {
	var result string
	list:=strings.Split(s," ")
	list2:=make([]string,len(list))
	for _,v:=range list{
		r := []rune(v)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		list2 = append(list2,string(r))
	}
	result = strings.Join(list2," ")
	return strings.TrimLeft(result," ")
}
