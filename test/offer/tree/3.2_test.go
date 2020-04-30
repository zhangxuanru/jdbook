/*
@Time : 2020/4/1 17:36
@Author : zxr
@File : 3.2_test
@Software: GoLand
*/
package tree

import (
	"fmt"
	"jdbook/offer/tree"
	"testing"

	"github.com/isdamir/gotype"
)

func TestArrayToTree(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("array:", data)
	root := tree.ArrayToTree(data, 0, len(data)-1)
	fmt.Println("tree:")
	gotype.PrintTreeMidOrder(root)
	fmt.Println()
}
