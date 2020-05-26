package main

/**
	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
*/

func main(){

}

func lengthOfLongestSubstring(s string) int {
	// 传入字符串长度为0
	if len(s) == 0 {
		return 0
	}
	// 返回值声明
	result :=0

	left ,right :=0,-1
	var freq [256]int
	for left < len(s){
		if right+1 <len(s) && freq[s[right+1]-'a'] ==0{
			freq[s[right+1]-'a']++
			right++
		}else{
			freq[s[left]-'a']--
			left++
		}
	}
	return result
}