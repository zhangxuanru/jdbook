/*
@Time : 2020/3/30 11:59
@Author : zxr
@File : 1.10
@Software: GoLand
*/
package linkNode

import "github.com/isdamir/gotype"

//删除单链表中的某个节点
func RemoveNode(head *gotype.LNode, p int) {
	node := head
	i := 0
	for node != nil {
		i++
		if i == p-1 {
			node.Next = node.Next.Next
		}
		node = node.Next
	}
}
