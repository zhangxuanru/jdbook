/*
@Time : 2020/4/8 15:20
@Author : zxr
@File : 4.15
@Software: GoLand
*/
package array

import "fmt"

//如何对数组进行循环移位

//方法一
func RightShift(arr []int, k int) {
	n := len(arr)
	if n == 0 || arr == nil {
		return
	}
	tmp := make([]int, k)
	copy(tmp, arr[n-k:])
	for i := n - k - 1; i >= 0; i-- {
		arr[i+k] = arr[i]
	}
	copy(arr[0:k], tmp)
}

//方法二
func RightShift2(arr []int, k int) {
	n := len(arr)
	if n == 0 || arr == nil {
		return
	}
	for k != 0 {
		k--
		tmp := arr[n-1]
		for i := n - 1; i > 0; i-- {
			arr[i] = arr[i-1]
		}
		arr[0] = tmp
	}
}

//方法三
func RightShift3(arr []int, k int) {
	n := len(arr)
	if n == 0 || arr == nil {
		return
	}
	k = k % n
	reverse(arr, 0, n-k-1)
	fmt.Println("n-k-1:", arr)
	reverse(arr, n-k, n-1)
	fmt.Println("n-1:", arr)
	reverse(arr, 0, n-1)
}

func reverse(arr []int, start, end int) {
	for start < end {
		tmp := arr[end]
		arr[end] = arr[start]
		arr[start] = tmp
		start++
		end--
	}
}
