/*
@Time : 2020/3/31 11:40
@Author : zxr
@File : 2.10
@Software: GoLand
*/
package array

//如何从数组中找出a+b=c+d的数对
func FindPairs(arr []int) []int {
	arrLen := len(arr)
	numMap := make(map[int][]int)
	var numsRet []int
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			sum := arr[i] + arr[j]
			if ints, exists := numMap[sum]; exists {
				ints = append(ints, arr[i], arr[j])
				numsRet = append(numsRet, ints...)
				return numsRet
			} else {
				numMap[sum] = []int{arr[i], arr[j]}
			}
		}
	}
	return numsRet
}
