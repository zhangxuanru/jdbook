/*
@Time : 2020/4/24 12:20
@Author : zxr
@File : 6.3_test
@Software: GoLand
*/
package matins

import (
	"fmt"
	"jdbook/offer/matins"
	"testing"
)

//如何实现不用除法操作符实现两个正整数的除法
func TestBow(t *testing.T) {
	res, rem := matins.Bow(13, 4)
	fmt.Printf("13/4商为：%d 余数为：%d\n", res, rem)

	n := 30
	for i, j := 0, 1; i < n; i, j = i+j, i {
		fmt.Println("i=", i)
	}
}
