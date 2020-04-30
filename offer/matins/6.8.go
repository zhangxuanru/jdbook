/*
@Time : 2020/4/24 15:29
@Author : zxr
@File : 6.8
@Software: GoLand
*/
package matins

//如何求有序数列的第1500个数的值（一个有序数列，数列中的每一个值都能被2或3或5整除。1是这个序列中的
// 第一个值，求第1500个元素是多少）

func SearchNum(n int) int {
	count := 0
	i := 0
	for {
		i++
		if i%2 == 0 || i%3 == 0 || i%5 == 0 {
			count++
		}
		if count == n {
			break
		}
	}
	return i
}
