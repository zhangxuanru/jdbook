/*
@Time : 2020/3/30 11:35
@Author : zxr
@File : 1.9_test
@Software: GoLand
*/
package linkNode

import (
	"fmt"
	"jdbook/offer/linkNode"
	"testing"

	"github.com/isdamir/gotype"
)

//合并两个有序的链表
func TestMergeLinkNode(t *testing.T) {
	head1 := &gotype.LNode{}
	head2 := &gotype.LNode{}
	CreateNodeT(head1, 1)
	CreateNodeT(head2, 2)
	head := linkNode.MergeLinkNode(head1, head2)
	for head != nil {
		fmt.Println(head.Data, ",")
		head = head.Next
	}
}

func CreateNodeT(node *gotype.LNode, i int) {
	cur := node
	for ; i < 9; i += 2 {
		cur.Next = &gotype.LNode{}
		cur.Data = i
		cur = cur.Next
	}
}
