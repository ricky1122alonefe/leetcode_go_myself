package main

import (
	"fmt"
)

func main() {
	a := []int{1, 4, 3, 2, 2}
	fmt.Println(kidsWithCandies(a, 3))
}
func kidsWithCandies(candies []int, extraCandies int) []bool {

	// type candyKv struct {
	// 	Index int
	// 	Value int
	// }
	var result []bool
	// var current int
	// m :=make(map[int]int )
	// for i:=0;i<len(candies);i++{
	// 	current =candies[i]+extraCandies
	// 	m[i]  = current
	// }
	// var list []candyKv
	// for k,v:=range m{
	// 	list = append(list,candyKv{k,v})
	// }
	//
	// sort.Slice(list, func(i, j int) bool {
	// 	return list[i].Value > list[j].Value  // 降序
	// })
	//
	// // 取出最高的值
	// fmt.Println(list)
	// topvalue:=list[0].Value
	// for i:=0;i<len(candies);i++{
	// 	if m[i] < topvalue{
	// 		result = append(result,false)
	// 	}else{
	// 		result = append(result,true )
	// 	}
	// }

	for i := 0; i < len(candies); i++ {
		// 当前 增加结果
		current := candies[i] + extraCandies

	}
	return result
}
