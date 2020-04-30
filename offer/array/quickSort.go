/*
@Time : 2020/4/1 11:53
@Author : zxr
@File : quickSort
@Software: GoLand
*/
package array

//快速排序
func QuickSort(nums []int) {
	sortArr(nums)
}

func sortArr(nums []int) {
	Sort(nums)
}

func Sort(nums []int) {
	left, right := 0, len(nums)-1
	for right > left {
		// 右边部分放大于
		if nums[right] > nums[0] {
			right--
			continue
		}
		// 左边部分放小于等于
		if nums[left] <= nums[0] {
			left++
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	if nums[right] < nums[0] {
		nums[0], nums[right] = nums[right], nums[0]
	}
	if len(nums[:right]) > 1 {
		sortArr(nums[:right])
	}
	if len(nums[right+1:]) > 1 {
		sortArr(nums[right+1:])
	}
}
