/*
@Time : 2020/4/28 11:06
@Author : zxr
@File : 8.1
@Software: GoLand
*/
package sorting

//选择排序
func SelectSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		temp := nums[i]
		flag := i
		for j := i + 1; j < n; j++ {
			if nums[j] < temp {
				temp = nums[j]
				flag = j
			}
		}
		if flag != i {
			nums[flag], nums[i] = nums[i], nums[flag]
		}
	}
}
