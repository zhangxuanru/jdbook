/*
@Time : 2020/4/28 16:00
@Author : zxr
@File : sortList
@Software: GoLand
*/
package sorting

//排序汇总
func SelectSort2(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		temp := arr[i]
		flag := i
		for j := i + 1; j < n; j++ {
			if arr[j] < temp {
				temp = arr[j]
				flag = j
			}
		}
		if flag != i {
			arr[flag], arr[i] = arr[i], arr[flag]
		}
	}
}

func InsertSort2(arr []int) {
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

func BubSort2(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func QuickSort2(arr []int) {
	if len(arr) == 0 {
		return
	}
	low := 0
	high := len(arr) - 1
	i := low
	j := high
	temp := arr[low]
	for i < j {
		if arr[j] >= temp {
			j--
			continue
		}
		if arr[i] <= temp {
			i++
			continue
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	if arr[j] < temp {
		arr[j], arr[low] = arr[low], arr[j]
	}
	if len(arr[:j]) > 0 {
		QuickSort2(arr[:j])
	}
	if len(arr[j+1:]) > 0 {
		QuickSort2(arr[j+1:])
	}
}

func QuickSort3(arr []int, low, high int) {
	if low >= high {
		return
	}
	i := low
	j := high
	temp := arr[low]
	for i < j {
		for arr[j] >= temp && i < j {
			j--
		}
		for arr[i] <= temp && i < j {
			i++
		}
		if i != j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	if arr[i] != temp {
		arr[i], arr[low] = temp, arr[i]
	}
	QuickSort3(arr, low, j)
	QuickSort3(arr, j+1, high)
}
