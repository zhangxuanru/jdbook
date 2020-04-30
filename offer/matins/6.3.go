/*
@Time : 2020/4/24 12:20
@Author : zxr
@File : 6.3
@Software: GoLand
*/
package matins

//如何实现不用除法操作符实现两个正整数的除法
func Bow(n1, n2 int) (res, rem int) {
	sum := 0
	//被除数减除数，直到相减结果小于除数为止
	for n1 >= n2 {
		n1 = n1 - n2
		sum++
	}
	res = sum
	rem = n1
	return
}
