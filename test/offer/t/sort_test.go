/*
@Time : 2020/4/23 11:09
@Author : zxr
@File : sort
@Software: GoLand
*/
package t

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	nums := []int{6, 7, 34, 1, 3, 71, 23, 2}
	bub(nums)
	fmt.Println("冒泡:", nums)

	nums1 := []int{9, 7, 34, 1, 6, 65, 21, 2}
	quickSort(nums1, 0, len(nums1)-1)
	fmt.Println("快排:", nums1)

	nums2 := []int{12, 53, 34, 29, 62, 5, 21, 2}
	quickSort2(nums2)
	fmt.Println("快排2:", nums2)
	nums3 := []int{12, 53, 34, 29, 62, 5, 21, 2}
	selectSort(nums3)
	fmt.Println("选择排序:", nums3)
}

//冒泡排序
func bub(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

//快速排序
func quickSort(arr []int, low, hight int) {
	if low > hight {
		return
	}
	i := low
	j := hight
	temp := arr[low]
	for i < j {
		for temp <= arr[j] && i < j {
			j--
		}
		for temp >= arr[i] && i < j {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	/*
		    i=j
			if arr[j] < temp {
				arr[j], arr[low] = temp, arr[j]
			}
	*/
	if i != low {
		arr[low], arr[i] = arr[i], temp
	}
	quickSort(arr, low, j-1)
	quickSort(arr, j+1, hight)
}

//快速排序方法二
func quickSort2(arr []int) {
	high := len(arr) - 1
	low := 0
	temp := arr[low]
	for low < high {
		if arr[high] >= temp {
			high--
			continue
		}
		if arr[low] <= temp {
			low++
			continue
		}
		if low != high {
			arr[low], arr[high] = arr[high], arr[low]
		}
	}
	if arr[high] < temp {
		arr[high], arr[0] = temp, arr[high]
	}
	if len(arr[:high]) > 0 {
		quickSort2(arr[:high])
	}
	if len(arr[high+1:]) > 0 {
		quickSort2(arr[high+1:])
	}

}

//选择排序
func selectSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		k := i
		for j := k + 1; j < n; j++ {
			if arr[j] < arr[k] {
				k = j
			}
		}
		if k != i {
			arr[k], arr[i] = arr[i], arr[k]
		}
	}
}
