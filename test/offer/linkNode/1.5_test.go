/*
@Time : 2020/3/27 17:54
@Author : zxr
@File : 1.5_test
@Software: GoLand
*/
package linkNode

import (
	"fmt"
	"jdbook/offer/linkNode"
	"testing"

	"github.com/isdamir/gotype"
)

func TestFindLinkNodeByKey(t *testing.T) {
	head := &gotype.LNode{}
	gotype.CreateNode(head, 10)

	node := &gotype.LNode{nil, &gotype.LNode{1, &gotype.LNode{2, &gotype.LNode{3, &gotype.LNode{4, &gotype.LNode{5, &gotype.LNode{6, &gotype.LNode{7, nil}}}}}}}}

	c := node.Next
	for c != nil {
		fmt.Println("val:", c.Data)
		c = c.Next
	}
	//gotype.PrintNode("old node:", head)
	key := linkNode.FindLinkNodeByKey(node, 3)
	fmt.Println(key)

	//key2 := linkNode.FindLinkNodeByKey2(node, 3)
	//fmt.Println(key2.Data)
}
