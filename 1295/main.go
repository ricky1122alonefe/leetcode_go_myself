package main

import "strconv"

func main(){

}

func findNumbers(nums []int) int {
	var count int
	for _,v:=range nums{
		if len(strconv.Itoa(v))%2==0{
			count ++
		}
	}
	return count
}