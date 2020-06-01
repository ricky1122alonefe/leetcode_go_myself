package main

func main() {

}


func destCity(paths [][]string) string {

	var result string
	var k_v = make(map[string]string)
	for i:=0;i<len(paths);i++{
		k_v[paths[i][0]] = paths[i][1]
	}

	for _,v :=range k_v{
		if  _,ok := k_v[v]; !ok{
			result = v
		}
	}
	return result

}