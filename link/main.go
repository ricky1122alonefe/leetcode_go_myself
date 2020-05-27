package main


func main(){

}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	    Val int
	     Next *ListNode
	 }
func getKthFromEnd(head *ListNode, k int) *ListNode {
	i:=0
	former,latter :=head,head       //定义两个指针,一个former指针,latter指针
	for former!=nil{                //如果头指针不为nil
		if i>=k{                    //只有k>=i的时候再开始让第二个指针开始工作
			latter = latter.Next
		}
		former = former.Next        //第一个指针一直工作
		i++                         //用来计数
	}
	return latter

}