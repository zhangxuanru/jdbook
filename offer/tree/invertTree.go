/*
@Time : 2020/4/2 18:44
@Author : zxr
@File : invertTree
@Software: GoLand
*/
package tree

import "github.com/isdamir/gotype"

//反转二叉树
func invertTree(root *gotype.BNode) *gotype.BNode {
	if root == nil {
		return root
	}
	root.LeftChild, root.RightChild = invertTree(root.RightChild), invertTree(root.LeftChild)
	return root
}
