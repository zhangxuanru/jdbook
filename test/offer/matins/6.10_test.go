/*
@Time : 2020/4/24 16:43
@Author : zxr
@File : 6.10_test
@Software: GoLand
*/
package matins

import (
	"fmt"
	"jdbook/offer/matins"
	"strconv"
	"testing"
)

//如何求二进制中1的个数
func TestFind1(t *testing.T) {
	find1 := matins.Find1(7)
	fmt.Println("7 find1:", find1)

	fmt.Println(strconv.IntSize)
}
