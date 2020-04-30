/*
@Time : 2020/3/26 18:56
@Author : zxr
@File : 1.2.1_test
@Software: GoLand
*/
package linkNode

import (
	"fmt"
	"jdbook/offer/linkNode"
	"testing"

	"github.com/isdamir/gotype"
)

//如何从有序链表中移除重复项

func TestRemoveNodeDup(t *testing.T) {
	node := &gotype.LNode{nil, &gotype.LNode{1, &gotype.LNode{2, &gotype.LNode{2, &gotype.LNode{3, &gotype.LNode{3, &gotype.LNode{5, &gotype.LNode{5, nil}}}}}}}}

	gotype.PrintNode("old:", node)
	linkNode.RemoveNodeDup(node)

	fmt.Printf("node:%p\n", node)

	gotype.PrintNode("haha:", node)
}
