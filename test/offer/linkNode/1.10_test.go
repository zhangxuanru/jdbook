/*
@Time : 2020/3/30 11:59
@Author : zxr
@File : 1.10_test
@Software: GoLand
*/
package linkNode

import (
	"fmt"
	"jdbook/offer/linkNode"
	"testing"

	"github.com/isdamir/gotype"
)

func TestRemoveNode(t *testing.T) {
	head := &gotype.LNode{}
	CreateT(head)
	linkNode.RemoveNode(head, 5)

	for head != nil {
		fmt.Println(head.Data, ",")
		head = head.Next
	}
}

func CreateT(node *gotype.LNode) {
	cur := node
	for i := 1; i < 8; i++ {
		cur.Next = &gotype.LNode{}
		cur.Data = i
		cur = cur.Next
	}
}
