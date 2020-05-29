package main

import (
	"strings"
	"reflect"
	"sort"
)

func main(){

}


func CheckPermutation(s1 string, s2 string) bool {
	if len(s1)!=len(s2){
		return false
	}
	list1:=strings.Split(s1,"")
	list2:= strings.Split(s2,"")
	sort.Strings(list1)
	sort.Strings(list2)
	return reflect.DeepEqual(list1,list2)
}