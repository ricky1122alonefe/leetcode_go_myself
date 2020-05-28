package main

import (
	"strings"
	"fmt"
)

func main(){
	fmt.Println(freqAlphabets("10#11#12"))
}


func freqAlphabets(s string) string {
	var str string
	mp :=make(map[string]string)
	mp["1"] = "a"
	mp["2"] = "b"
	mp["3"] = "c"
	mp["4"] = "d"
	mp["5"] = "e"
	mp["6"] = "f"
	mp["7"] = "g"
	mp["8"] = "h"
	mp["9"] = "i"
	mp["10"] = "j"
	mp["11"] = "k"
	mp["12"] = "l"
	mp["13"] = "m"
	mp["14"] = "n"
	mp["15"] = "o"
	mp["16"] = "p"
	mp["17"] = "q"
	mp["18"] = "r"
	mp["19"] = "s"
	mp["20"] = "t"
	mp["21"] = "u"
	mp["22"] = "v"
	mp["23"] = "w"
	mp["24"] = "x"
	mp["25"] = "y"
	mp["26"] = "z"
	success:=strings.IndexAny(s,"#")

	switch success{
	// 此情况只有a-i
	case -1:
		for i:=0;i<len(s);i++{
			str+=mp[string(s[i])]
		}
	default:
		list:=strings.Split(s,"#")
		if strings.LastIndex(s,"#") == len(s)-1{
			list =list[0:len(list)-1]
		}

		for _,v:=range list{
			l :=len(v)
			switch l {
			case 2:
				str+= mp[v]
			default:
				for i:=0;i<len(v)-2;i++{
					str+=mp[string(v[i])]
				}
				str+=mp[string(v[len(v)-2:])]
			}
		}
	}
	return str
}