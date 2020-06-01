package main

import (
	"strconv"
	"fmt"
)

func main(){
	fmt.Println(subtractProductAndSum(234))
}

func subtractProductAndSum(n int) int {
	num_s := strconv.Itoa(n)
	var sub int = 1;
	var sum int = 0;
	for _,v := range num_s{
		tmp := int(v - '0')
		sub *= tmp
		sum += tmp
	}
	return sub - sum

}
