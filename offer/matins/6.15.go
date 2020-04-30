/*
@Time : 2020/4/24 17:46
@Author : zxr
@File : 6.15
@Software: GoLand
*/
package matins

import "fmt"

//如何不使用循环输出1到100
func PrintNum(n int) {
	if n > 0 {
		PrintNum(n - 1)
		fmt.Println(n)
	}
}
