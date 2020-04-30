/*
@Time : 2020/4/3 14:51
@Author : zxr
@File : 4.2
@Software: GoLand
*/
package array

//如何查找数组中元素的最大值和最小值

//方法一（分治法）
func GetArrMaxAndMin(arr []int, l, r int) (max, min int) {
	if arr == nil {
		return
	}
	m := (l + r) / 2
	if l == r {
		max = arr[l]
		min = arr[l]
		return
	}
	if l+1 == r {
		if arr[l] > arr[r] {
			max = arr[l]
			min = arr[r]
		} else {
			max = arr[r]
			min = arr[l]
		}
		return
	}
	lMax, lMin := GetArrMaxAndMin(arr, l, m)
	rMax, rMin := GetArrMaxAndMin(arr, m+1, r)
	if lMax > rMax {
		max = lMax
	} else {
		max = rMax
	}
	if lMin < rMin {
		min = lMin
	} else {
		min = rMin
	}
	return
}

//方法二（顺序比较法）
func GetArrMaxAndMin2(arr []int) (max, min int) {
	if arr == nil {
		return
	}
	max = arr[0]
	min = arr[0]
	for _, v := range arr {
		if v > max {
			max = v
			continue
		}
		if v < min {
			min = v
			continue
		}
	}
	return
}
