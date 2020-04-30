/*
@Time : 2020/4/3 15:58
@Author : zxr
@File : rotatearr
@Software: GoLand
*/
package array

//旋转数组
/**
nums := [1,2,3,4,5,6,7] k=3
nums = [5,6,7,1,2,3,4]
**/
func RotateArr(arr []int, k int) []int {
	n := len(arr)
	r := arr[n-k:]
	r = append(r, arr[0:n-k]...)
	return r
}

//方法二
func RotateArr2(arr []int, k int) {
	if arr == nil || k <= 0 || k >= len(arr)-1 {
		return
	}
	swap(arr, 0, k)
	swap(arr, k+1, len(arr)-1)
	swap(arr, 0, len(arr)-1)
}

func swap(arr []int, low, high int) {
	for ; low < high; low, high = low+1, high-1 {
		arr[low], arr[high] = arr[high], arr[low]
	}
}
