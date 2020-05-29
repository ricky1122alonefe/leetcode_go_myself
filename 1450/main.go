package main

import "fmt"

func main(){
	start:=[]int{1,2,3}
	end:=[]int{3,2,7}
	query:=4
	fmt.Println(busyStudent(start,end,query))
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var js int = 0;
	for k,v := range endTime{
		if v >= queryTime{
			if startTime[k] <= queryTime{
				js++
			}
		}
	}
	return js
	
}
