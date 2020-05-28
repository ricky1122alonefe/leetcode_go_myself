package main

import "fmt"

func main()  {
	fmt.Println(maximum(1,2))
}

// 注意位运算

func maximum(a int, b int) int {
	v := uint64(a-b)>>63

	return int(v)*b + int(1^v)*a
}
