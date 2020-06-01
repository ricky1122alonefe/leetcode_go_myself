package main

import (
	"sort"
	"reflect"
)

func main(){

}

func canBeEqual(target []int, arr []int) bool {
	sort.Ints(arr)
	sort.Ints(target)
	return reflect.DeepEqual(target,arr)
}
