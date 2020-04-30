/*
@Time : 2020/4/1 18:43
@Author : zxr
@File : 3.3
@Software: GoLand
*/
package tree

import (
	"fmt"
	queue2 "jdbook/offer/queue"

	"github.com/isdamir/gotype"
)

//如何从顶部开始逐层打印二叉树结点数据

//方法一（用队列来实现）
func PrintTree(root *gotype.BNode) {
	if root == nil {
		return
	}
	var p *gotype.BNode
	queue := queue2.NewSliceDataQueue()
	queue.EnQueue(root)
	for queue.Size() > 0 {
		p = queue.DeQueue().(*gotype.BNode)
		fmt.Print(p.Data, ",")
		if p.LeftChild != nil {
			queue.EnQueue(p.LeftChild)
		}
		if p.RightChild != nil {
			queue.EnQueue(p.RightChild)
		}
	}
}

//方法二
func PrintTree2(root *gotype.BNode, leven int) int {
	if root == nil || leven < 0 {
		return 0
	} else if leven == 0 {
		fmt.Print(root.Data, ",")
		return 1
	} else {
		return PrintTree2(root.LeftChild, leven-1) + PrintTree2(root.RightChild, leven-1)
	}
}

//test print tree
func T2(root *gotype.BNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Data, ",")
	if root.LeftChild != nil {
		T2(root.LeftChild)
	}
	if root.RightChild != nil {
		T2(root.RightChild)
	}
}
