/*
@Time : 2020/4/28 11:43
@Author : zxr
@File : 8.3
@Software: GoLand
*/
package sorting

//冒泡排序
func BubSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
