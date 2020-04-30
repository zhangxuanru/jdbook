/*
@Time : 2020/4/24 10:31
@Author : zxr
@File : 2point_test
@Software: GoLand
*/
package t

import (
	"fmt"
	"testing"
)

func TestTwoPointFind(t *testing.T) {
	nums := []int{3, 4, 5, 7, 23, 56, 90}
	find := TwoPointFind(nums, 56)
	fmt.Println("is Fomd 56:", find)

	find = TwoPointFind(nums, 132)
	fmt.Println("is Fomd 132:", find)
}

//二分查找
func TwoPointFind(nums []int, n int) bool {
	low := 0
	high := len(nums)
	for low < high {
		mid := (low + high) / 2
		if nums[mid] == n {
			return true
		}
		if nums[mid] > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
