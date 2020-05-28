package main

import (
	"sort"
	"fmt"
)

func main() {
	list:=[]int{-4,-1,0,3,10}
	fmt.Println(sortedSquares(list))
}

func sortedSquares(A []int) []int {
	for k,v:=range A{
		A[k] *=v
	}
	sort.Ints(A)
	return A
}
