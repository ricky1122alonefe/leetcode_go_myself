package main

import "fmt"

func main() {
	arr :=[]int{8,1,2,2,3}
	fmt.Println(smallerNumbersThanCurrent(arr))
}

func smallerNumbersThanCurrent(nums []int) []int {
	result:=make([]int,len(nums))
	for i:=0;i<len(nums);i++{
		var count int
		for j:=0;j<len(nums);j++{
			if i==j{
				continue
			}
			if nums[i]>nums[j]{
				count++
			}
		}
		result[i] = count
	}
	return result
}
