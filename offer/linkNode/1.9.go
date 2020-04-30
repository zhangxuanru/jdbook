/*
@Time : 2020/3/30 11:35
@Author : zxr
@File : 1.9
@Software: GoLand
*/
package linkNode

import (
	"github.com/isdamir/gotype"
)

//合并两个有序的链表
func MergeLinkNode(head1 *gotype.LNode, head2 *gotype.LNode) *gotype.LNode {
	if head1 == nil && head2 == nil {
		return nil
	}
	cur := &gotype.LNode{}
	head := cur
	for {
		if head1 == nil && head2 == nil {
			break
		}
		if head1 == nil || head1.Data == nil {
			head.Next = head2
			break
		}
		if head2 == nil || head2.Data == nil {
			head.Next = head1
			break
		}
		if head1.Data.(int) > head2.Data.(int) {
			head.Next = head2
			head2 = head2.Next
			head = head.Next
		}
		if head2.Data.(int) > head1.Data.(int) {
			head.Next = head1
			head1 = head1.Next
			head = head.Next
		}
	}
	return cur.Next
}
