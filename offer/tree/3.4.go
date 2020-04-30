/*
@Time : 2020/4/2 12:02
@Author : zxr
@File : 3.4
@Software: GoLand
*/
package tree

import (
	"fmt"
	"math"

	"github.com/isdamir/gotype"
)

//如何求一棵二叉树的最大子树和
var MaxSum = math.MinInt64

func FindMaxSubTree(root *gotype.BNode, maxRoot *gotype.BNode) int {
	if root == nil {
		return 0
	}
	lMax := FindMaxSubTree(root.LeftChild, maxRoot)
	rMax := FindMaxSubTree(root.RightChild, maxRoot)
	sum := lMax + rMax + root.Data.(int)
	fmt.Println("lMax:", lMax, "rMax:", rMax, "data:", root.Data, "sum:", sum)
	if sum > MaxSum {
		MaxSum = sum
		maxRoot.Data = root.Data
	}
	return sum
}
