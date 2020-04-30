/*
@Time : 2020/3/27 17:54
@Author : zxr
@File : 1.5
@Software: GoLand
*/
package linkNode

import "github.com/isdamir/gotype"

//找出单链表中倒数第K个元素
func FindLinkNodeByKey(head *gotype.LNode, k int) interface{} {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for fast != nil {
		fast = fast.Next
		if k <= 0 {
			slow = slow.Next
		}
		k--
	}
	return slow.Data
}

//方法二
func FindLinkNodeByKey2(head *gotype.LNode, k int) *gotype.LNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	//从0开始， 所以要小于k
	for i := 0; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
