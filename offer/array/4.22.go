/*
@Time : 2020/4/8 18:38
@Author : zxr
@File : 4.22
@Software: GoLand
*/
package array

//如何从三个有序数组中找出它们的公共元素
func FindCommon(arr1, arr2, arr3 []int) []int {
	ret := make([]int, 0)
	if arr1 == nil || arr2 == nil || arr3 == nil {
		return ret
	}
	n1 := len(arr1)
	n2 := len(arr2)
	n3 := len(arr3)
	for i, j, k := 0, 0, 0; i < n1 && j < n2 && k < n3; {
		if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
			ret = append(ret, arr1[i])
			i++
			j++
			k++
		} else if arr1[i] < arr2[j] {
			i++
		} else if arr2[j] < arr3[k] {
			j++
		} else {
			k++
		}
	}
	return ret
}
