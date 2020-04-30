/*
@Time : 2020/4/3 17:37
@Author : zxr
@File : 4.6
@Software: GoLand
*/
package array

import (
	"fmt"
	"sort"
)

//找出数组中第k小的数

func FindSmallKNumber(nums []int, k int) int {
	n := len(nums)
	if nums == nil || k >= n {
		return -1
	}
	//排序
	//sort.Ints(nums)

	//倒序排序
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println("sort nums", nums)

	//快速排序
	quickSort(nums, 0, len(nums)-1)
	fmt.Println("nums:", nums)
	return nums[k-1]
}

//快速排序
func quickSort(nums []int, low, high int) {
	if nums == nil || low > high {
		return
	}
	mid := 0
	for low < high {
		if nums[high] > nums[mid] {
			high--
			continue
		}
		if nums[low] < nums[mid] {
			low++
			continue
		}
		nums[low], nums[high] = nums[high], nums[low]
	}
	if nums[high] < nums[mid] {
		nums[mid], nums[high] = nums[high], nums[mid]
	}
	if len(nums[:high]) > 1 {
		quickSort(nums[:high], 0, len(nums[:high])-1)
	}
	if len(nums[high+1:]) > 1 {
		quickSort(nums[high+1:], high, len(nums[high+1:])-1)
	}
}
