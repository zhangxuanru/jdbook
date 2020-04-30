/*
@Time : 2020/4/24 15:33
@Author : zxr
@File : 6.8_test
@Software: GoLand
*/
package matins

import (
	"fmt"
	"jdbook/offer/matins"
	"testing"
)

func TestSearchNum(t *testing.T) {
	num := matins.SearchNum(1500)
	fmt.Println("第1500个序列是：", num)
}
