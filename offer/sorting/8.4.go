/*
@Time : 2020/4/28 11:53
@Author : zxr
@File : 8.4
@Software: GoLand
*/
package sorting

//归并排序
func McrgeSort(arr []int, low, high int) {
	if low < high {
		q := (low + high) / 2
		McrgeSort(arr, low, q)
		McrgeSort(arr, q+1, high)
		Merge(arr, low, q, high)
	}
}

func Merge(arr []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1)
	R := make([]int, n2)
	for i, k := 0, p; i < n1; i, k = i+1, k+1 {
		L[i] = arr[k]
	}
	for i, k := 0, q+1; i < n2; i, k = i+1, k+1 {
		R[i] = arr[k]
	}
	var i, k, j int
	for i, k, j = 0, p, 0; i < n1 && j < n2; k++ {
		if L[i] > R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
	}
	if i < n1 {
		for j = i; j < n1; j, k = j+1, k+1 {
			arr[k] = L[j]
		}
	}
	if j < n2 {
		for i = j; i < n2; i, k = i+1, k+1 {
			arr[k] = R[i]
		}
	}
}
