/*
@Time : 2020/4/24 16:42
@Author : zxr
@File : 6.10
@Software: GoLand
*/
package matins

//如何求二进制中1的个数
func Find1(n int) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}

func Find11(n int) int {
	count := 0
	for n > 0 {
		if n != 0 {
			n = n & (n - 1)
		}
		count++
	}
	return count
}
