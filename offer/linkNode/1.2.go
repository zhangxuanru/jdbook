/*
@Time : 2020/3/26 16:30
@Author : zxr
@File : 1.2
@Software: GoLand
*/
package linkNode

import (
	"github.com/isdamir/gotype"
)

//如何从无序链表中移除重复项

// 方法一
func RemoveDup1(node *gotype.LNode) {
	if node == nil || node.Next == nil {
		return
	}
	outerCur := node.Next
	var innerCur *gotype.LNode
	var innerPer *gotype.LNode
	for ; outerCur != nil; outerCur = outerCur.Next {
		for innerCur, innerPer = outerCur.Next, outerCur; innerCur != nil; {
			if outerCur.Data == innerCur.Data {
				innerPer.Next = innerCur.Next
				innerCur = innerCur.Next
			} else {
				innerPer = innerCur
				innerCur = innerCur.Next
			}
		}
	}
}

//方法二
func RemoveDup2(node *gotype.LNode) {
	if node == nil {
		return
	}
	cur := node
	nodeMap := make(map[interface{}]bool)
	for ; cur != nil; cur = cur.Next {
		nodeMap[cur.Data] = true
	}
	curOut := gotype.NewLNode()
	newNode := curOut
	for ; node != nil; node = node.Next {
		if _, exists := nodeMap[node.Data]; exists {
			newNode.Next = gotype.NewLNode()
			newNode.Data = node.Data
			newNode = newNode.Next
			delete(nodeMap, node.Data)
		}
	}
	gotype.PrintNode("curOut:", curOut)
}
