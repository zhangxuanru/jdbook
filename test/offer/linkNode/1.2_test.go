/*
@Time : 2020/3/26 16:33
@Author : zxr
@File : 1.2_test
@Software: GoLand
*/
package linkNode

import (
	"jdbook/offer/linkNode"
	"testing"

	"github.com/isdamir/gotype"
)

//如何从无序链表中移除重复项
func TestRemoveDup1(t *testing.T) {
	node := &gotype.LNode{nil, &gotype.LNode{1, &gotype.LNode{3, &gotype.LNode{1, &gotype.LNode{5, &gotype.LNode{5, &gotype.LNode{7, nil}}}}}}}

	//fun2
	linkNode.RemoveDup2(node)

	//fun1
	linkNode.RemoveDup1(node)

	gotype.PrintNode("RemoveDup1:", node)
}
