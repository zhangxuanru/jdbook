/*
@Time : 2020/4/1 18:44
@Author : zxr
@File : 3.3
@Software: GoLand
*/
package tree

import (
	"fmt"
	"jdbook/offer/tree"
	"testing"
)

//如何从顶部开始逐层打印二叉树结点数据
func TestPrintTree(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := tree.ArrayToTree(nums, 0, len(nums)-1)
	//fun1
	tree.PrintTree(root)

	fmt.Println()
	fmt.Println("------------")
	//tree.T2(root)

	for i := 0; i <= 3; i++ {
		//fun2
		tree.PrintTree2(root, i)
	}
}
