package main

import (
	"strings"
	"fmt"
)

func main(){
	fmt.Println(toLowerCase("Hello"))
}
func toLowerCase(str string) string {
	return strings.ToLower(str)
}