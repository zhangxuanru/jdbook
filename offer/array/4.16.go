/*
@Time : 2020/4/8 16:39
@Author : zxr
@File : 4.16
@Software: GoLand
*/
package array

//如何在有规律的二维数组中进行高效的数据查找
func FindWithBinary(arr [][]int, data int) bool {
	n := len(arr)
	if arr == nil || n == 0 {
		return false
	}
	colNum := len(arr[0])
	for i, j := 0, colNum-1; i < n && j >= 0; {
		if arr[i][j] == data {
			return true
		} else if arr[i][j] > data {
			j--
		} else {
			i++
		}
	}
	return false
}
