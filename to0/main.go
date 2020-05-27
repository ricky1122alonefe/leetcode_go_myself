package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14))
}

func numberOfSteps(num int) int {
	result := 0
	for num > 0 {
		if num%2 == 1 {
			num = num - 1
		} else {
			num = num / 2
		}
		result = result + 1
	}
	return result
}
