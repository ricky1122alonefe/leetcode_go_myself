package main

import "fmt"

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main(){
	nums :=[]int{2,7,11,15}
	fmt.Println(twoSums(nums,17))
}

func twoSums(num []int,target int)[]int{
	// 创建map 确保每个数字只能用一次
	numMap :=make(map[int]int)
	for i:=0;i<len(num);i++{
		// 这里做减法
		another := target-num[i]
		if _,ok:=numMap[another];ok{
			return []int{numMap[another],i}
		}
		numMap[num[i]] = i
	}
	return nil
}
