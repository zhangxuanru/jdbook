/*
@Time : 2020/3/27 14:25
@Author : zxr
@File : 1.3
@Software: GoLand
*/
package linkNode

import (
	"github.com/isdamir/gotype"
)

//如何计算两个单链表的数之和

/**
给定两个单链表，链表的每个结点代表一位数，计算两个数的和。
例如：输入链表(3->1->5)和链表(5->9->2)，输出：8->0->8，即513+295=808，注意个位数在链表头。
*/

func AddTwoNumbers(head1 *gotype.LNode, head2 *gotype.LNode) *gotype.LNode {
	if head1 == nil && head2 == nil {
		return nil
	}
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	head := &gotype.LNode{
		Data: -1,
		Next: nil,
	}
	tail := head
	flag := 0

	for head1 != nil || head2 != nil {
		val1, val2 := 0, 0
		if head1 != nil {
			if head1.Data != nil {
				val1 = head1.Data.(int)
			}
			head1 = head1.Next
		}
		if head2 != nil {
			if head2.Data != nil {
				val2 = head2.Data.(int)
			}
			head2 = head2.Next
		}
		sum := val1 + val2 + flag
		if sum >= 10 {
			flag = 1
			sum -= 10
		} else {
			flag = 0
		}
		tail = listAdd(tail, sum)
	}
	if flag == 1 {
		tail = listAdd(tail, 1)
	}
	return head.Next
}

func listAdd(tail *gotype.LNode, num int) *gotype.LNode {
	numNode := &gotype.LNode{
		num, nil,
	}
	tail.Next = numNode
	tail = numNode
	return tail
}
