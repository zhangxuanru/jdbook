/*
@Time : 2020/4/1 17:35
@Author : zxr
@File : 3.2
@Software: GoLand
*/
package tree

import (
	"github.com/isdamir/gotype"
)

//如何把一个有序的整数数组放到二叉树中
func ArrayToTree(arr []int, start, end int) *gotype.BNode {
	var root *gotype.BNode
	if end >= start {
		root = gotype.NewBNode()
		mid := (start + end + 1) / 2
		root.Data = arr[mid]
		root.LeftChild = ArrayToTree(arr, start, mid-1)
		root.RightChild = ArrayToTree(arr, mid+1, end)
	}
	return root
}
