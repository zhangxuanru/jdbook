/*
@Time : 2020/4/2 15:08
@Author : zxr
@File : 3.5
@Software: GoLand
*/
package tree

import "github.com/isdamir/gotype"

//如何判断两棵二叉树是否相等
func IsEqualTree(root1 *gotype.BNode, root2 *gotype.BNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil {
		return false
	}
	if root2 == nil && root1 != nil {
		return false
	}
	if root1.Data != root2.Data {
		return false
	}
	leftIsEq := IsEqualTree(root1.LeftChild, root2.LeftChild)
	if leftIsEq == false {
		return false
	}
	rightIsEq := IsEqualTree(root1.RightChild, root2.RightChild)
	if rightIsEq == false {
		return false
	}
	return true
}
