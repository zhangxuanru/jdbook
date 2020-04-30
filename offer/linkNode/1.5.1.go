/*
@Time : 2020/3/27 19:24
@Author : zxr
@File : 1.5.1
@Software: GoLand
*/
package linkNode

import (
	"github.com/isdamir/gotype"
)

//下周继续

//将单链表向右旋转K个位置
func RouteK(head *gotype.LNode, k int) {
	if head == nil || head.Next == nil {
		return
	}
	fast := head.Next
	slow := head.Next
	for i := 0; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	tmp := slow
	slow = slow.Next
	tmp.Next = nil
	fast.Next = head.Next
	head.Next = slow
}
