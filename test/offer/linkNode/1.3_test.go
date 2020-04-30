/*
@Time : 2020/3/27 14:28
@Author : zxr
@File : 1.3_test
@Software: GoLand
*/
package linkNode

import (
	"jdbook/offer/linkNode"
	"testing"

	"github.com/isdamir/gotype"
)

func TestAddTwoNumbers(t *testing.T) {
	l1, l2 := CreteNodeT()
	gotype.PrintNode("l1:", l1)
	gotype.PrintNode("l2:", l2)
	numbers := linkNode.AddTwoNumbers(l1, l2)
	gotype.PrintNode("numbers:", numbers)
}

func CreteNodeT() (l1 *gotype.LNode, l2 *gotype.LNode) {
	l1 = &gotype.LNode{}
	l2 = &gotype.LNode{}
	cur := l1
	for i := 0; i < 7; i++ {
		cur.Next = &gotype.LNode{}
		cur.Data = i + 2
		cur = cur.Next
	}
	cur = l2
	for i := 9; i >= 4; i-- {
		cur.Next = &gotype.LNode{}
		cur.Data = i + 1
		cur = cur.Next
	}
	return
}
