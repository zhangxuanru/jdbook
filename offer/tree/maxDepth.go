/*
@Time : 2020/4/2 11:11
@Author : zxr
@File : maxDepth
@Software: GoLand
*/
package tree

import (
	"github.com/isdamir/gotype"
)

//求二叉树最大深度
func MaxDepth(root *gotype.BNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.LeftChild)
	right := MaxDepth(root.RightChild)
	if left > right {
		return left + 1
	}
	return right + 1
}
