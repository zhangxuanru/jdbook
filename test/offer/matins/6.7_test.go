/*
@Time : 2020/4/24 15:26
@Author : zxr
@File : 6.7_test
@Software: GoLand
*/
package matins

import (
	"fmt"
	"jdbook/offer/matins"
	"testing"
)

//如何比较a,b两个数的大小， 不能用大于，小于 语句
func TestMax1(t *testing.T) {
	max1 := matins.Max1(160, 20)
	fmt.Println("最大数是:", max1)
}
