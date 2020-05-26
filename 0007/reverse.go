package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(reverse(12))
}


func reverse(x int) int {
	res := 0
	for x != 0 {
		carry := x % 10
		res = res*10 + carry
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return res
}