package main

import (
	"strconv"
	"strings"
	"fmt"
)

func main(){
	fmt.Println(maximum69Number(9966))
}

func maximum69Number (num int) int {
	str:=strconv.Itoa(num)
	index:=strings.Index(str,"6")
	if index>-1{
		str = strings.Replace(str,"6","9",1)
	}

	result ,_ :=strconv.Atoi(str)
	return result
}