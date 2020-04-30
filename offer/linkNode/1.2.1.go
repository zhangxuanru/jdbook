/*
@Time : 2020/3/26 18:56
@Author : zxr
@File : 1.2.1
@Software: GoLand
*/
package linkNode

import (
	"fmt"

	"github.com/isdamir/gotype"
)

//如何从有序链表中移除重复项
func RemoveNodeDup(node *gotype.LNode) {
	if node == nil {
		return
	}
	currNode := node
	fmt.Printf("%p----%p\n\n", currNode, node)
	for ; node != nil; node = node.Next {
		if currNode.Data == currNode.Next.Data {
			currNode.Next = currNode.Next.Next
		} else {
			currNode = currNode.Next
		}
		fmt.Printf("currNode point:%p\n\n", currNode)
	}
	fmt.Printf("%p----%p\n\n", currNode, node)
}
