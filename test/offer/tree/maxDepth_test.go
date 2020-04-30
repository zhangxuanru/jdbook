/*
@Time : 2020/4/2 11:13
@Author : zxr
@File : maxDepth_test
@Software: GoLand
*/
package tree

import (
	"fmt"
	"jdbook/offer/tree"
	"testing"
)

//求二叉树最大深度
func TestMaxDepth(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := tree.ArrayToTree(nums, 0, len(nums)-1)
	depth := tree.MaxDepth(root)
	fmt.Println("二叉树最大深度:", depth)
}
