/*
@Time : 2020/4/29 18:33
@Author : zxr
@File : composite_test
@Software: GoLand
*/
package Composite

import (
	"fmt"
	"testing"
)

//组合模式
func TestCat_Traverse(t *testing.T) {
	composite := NewComposite()
	for i := 0; i < 3; i++ {
		leaf := NewLeaf(i*2 + 1)
		cat := NewCat(fmt.Sprintf("cat:%d", i))
		composite.Add(leaf)
		composite.Add(cat)
	}
	composite.Traverse()
}
