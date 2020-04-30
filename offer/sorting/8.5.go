/*
@Time : 2020/4/28 14:36
@Author : zxr
@File : 8.5
@Software: GoLand
*/
package sorting

//快速排序
func QS(arr []int) {
	low := 0
	high := len(arr) - 1
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
		arr[low], arr[high] = arr[high], arr[low]
	}
	if arr[high] < temp {
		arr[high], arr[0] = arr[0], arr[high]
	}
	if len(arr[:high]) > 0 {
		QS(arr[:high])
	}
	if len(arr[high+1:]) > 0 {
		QS(arr[high+1:])
	}
}

func QS1(arr []int, low, high int) {
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
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	if arr[i] != temp {
		arr[i], arr[low] = arr[low], arr[i]
	}
	QS1(arr, low, j-1)
	QS1(arr, j+1, high)
}
