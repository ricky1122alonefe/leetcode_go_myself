package main

func main(){

}

func reverseString(s []byte)  {
	n := len(s) //[l,n)
	mid := n / 2
	for i := 0; i < mid; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}
