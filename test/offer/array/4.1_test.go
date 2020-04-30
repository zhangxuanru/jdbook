/*
@Time : 2020/4/3 11:33
@Author : zxr
@File : 4.1_test
@Software: GoLand
*/
package array

import (
	"fmt"
	"jdbook/offer/array"
	"testing"
)

func TestFindDupByHash(t *testing.T) {
	nums := []int{1, 3, 4, 2, 5, 3}
	i := array.FindDupByHash(nums)
	fmt.Println("HASH：重复的值是：", i)

	sum := array.FindDupBySum(nums)
	fmt.Println("SUM：重复的值是：", sum)

	sum = array.FindDupByXOR(nums)
	fmt.Println("SUM：重复的值是：", sum)

	byMap := array.FindDupByMap(nums)
	fmt.Println("SUM：重复的值是：", byMap)

}
