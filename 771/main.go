package main

import "fmt"

func main(){
	fmt.Println(numJewelsInStones("aA","aAAbbbb"))
}

func numJewelsInStones(J string, S string) int {
	var result int
	myMap :=make(map[string]int)
	for _,v:=range S{
		if _,ok:=myMap[string(v)];ok{
			myMap[string(v)] +=1
			continue
		}

		myMap[string(v)] = 1
	}
	for _,v:=range J{
		result+=myMap[string(v)]
	}
	return result
}
