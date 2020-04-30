/*
@Time : 2020/3/26 14:36
@Author : zxr
@File : 1.1
@Software: GoLand
*/
package linkNode

import (
	"github.com/isdamir/gotype"
)

//如何实现链表逆序

//方法一
func Reverse1(head *gotype.LNode) *gotype.LNode {
	var per *gotype.LNode
	if head == nil {
		return per
	}
	for head != nil {
		next := head.Next
		head.Next = per
		per = head
		head = next
	}
	return per
}

//方法二
func Reverse2(head *gotype.LNode) {
	if head == nil || head.Next == nil {
		return
	}
	var pre *gotype.LNode
	var cur *gotype.LNode
	next := head.Next
	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	head.Next = pre
}

//方法三 （递归）
func Reverse3(head *gotype.LNode) {
	if head == nil {
		return
	}
	child := Reverse3Child(head.Next)
	head.Next = child
}

func Reverse3Child(node *gotype.LNode) *gotype.LNode {
	if node == nil || node.Next == nil {
		return node
	}
	child := Reverse3Child(node.Next)
	node.Next.Next = node
	node.Next = nil
	return child
}

//方法四
func Reverse4(head *gotype.LNode) {
	isTail := false
	perNode := &gotype.LNode{}
	for head != nil {
		if isTail == false {
			perNode.Data = head.Data
			perNode.Next = nil
			isTail = true
		} else {
			tmp := &gotype.LNode{head.Data, perNode}
			perNode = tmp
		}
		head = head.Next
	}
	gotype.PrintNode("perNode:", perNode)
}
