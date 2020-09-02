package main

import "fmt"

func main() {
	for _,v:=range printNumbers(2){
		fmt.Println(v)
	}
}
func printNumbers(n int) []int {
	var res []int
	var max int
	for n>0{
		max =max*10+9
		n--
	}
	for i:=1;i<=max;i++{
		res = append(res,i)
	}
	return res
}