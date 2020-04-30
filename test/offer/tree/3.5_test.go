/*
@Time : 2020/4/2 15:09
@Author : zxr
@File : 3.5_test
@Software: GoLand
*/
package tree

import (
	"fmt"
	"jdbook/offer/tree"
	"testing"
)

//如何判断两棵二叉树是否相等
func TestIsEqualTree(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	root1 := tree.ArrayToTree(nums1, 0, len(nums1)-1)
	root2 := tree.ArrayToTree(nums2, 0, len(nums2)-1)
	equalTree := tree.IsEqualTree(root1, root2)
	fmt.Println("两棵二叉树是否相等:", equalTree)
}
