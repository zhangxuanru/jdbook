/*
@Time : 2020/4/28 11:24
@Author : zxr
@File : 8.2
@Software: GoLand
*/
package sorting

//插入排序
func InsertSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		temp, j := arr[i], i
		if arr[j-1] > temp {
			for j >= 1 && arr[j-1] > temp {
				arr[j] = arr[j-1]
				j--
			}
			arr[j] = temp
		}
	}
}
