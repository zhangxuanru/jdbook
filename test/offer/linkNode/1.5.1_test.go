/*
@Time : 2020/3/27 19:25
@Author : zxr
@File : 1.5.1_test
@Software: GoLand
*/
package linkNode

import (
	"jdbook/offer/linkNode"
	"testing"

	"github.com/isdamir/gotype"
)

func TestRouteK(t *testing.T) {
	head := &gotype.LNode{}
	gotype.CreateNode(head, 8)

	gotype.PrintNode("haha:", head)

	linkNode.RouteK(head, 3)

	gotype.PrintNode("haha:", head)
	//gotype.PrintNode("slow:", slow)
}
