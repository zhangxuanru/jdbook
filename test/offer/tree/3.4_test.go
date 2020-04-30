/*
@Time : 2020/4/2 12:20
@Author : zxr
@File : 3.4_test
@Software: GoLand
*/
package tree

import (
	"fmt"
	"jdbook/offer/tree"
	"testing"

	"github.com/isdamir/gotype"
)

//如何求一棵二叉树的最大子树和
func TestFindMaxSubTree(t *testing.T) {
	root := CreateTree()
	maxTree := &gotype.BNode{}
	tree.FindMaxSubTree(root, maxTree)
	fmt.Println("最大子树和:", tree.MaxSum)
	fmt.Println("对应子树根结点为:", maxTree.Data)
}

func CreateTree() *gotype.BNode {
	root := &gotype.BNode{}
	node1 := &gotype.BNode{}
	node2 := &gotype.BNode{}
	node3 := &gotype.BNode{}
	node4 := &gotype.BNode{}
	root.Data = 6
	node1.Data = 3
	node2.Data = -7
	node3.Data = -1
	node4.Data = 9
	root.LeftChild = node1
	root.RightChild = node2
	node1.LeftChild = node3
	node1.RightChild = node4
	return root
}
