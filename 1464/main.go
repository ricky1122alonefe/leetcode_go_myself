package main

import (
	"sort"
	"fmt"
)

func main(){
	arr:=[]int{3,4,5,2}
	fmt.Println(maxProduct(arr))
}


func maxProduct(nums []int) int {
	sort.Ints(nums)

	fmt.Println(nums)
	return (nums[len(nums)-1]-1)*(nums[len(nums)-2]-1)
}