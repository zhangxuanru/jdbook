/*
@Time : 2020/4/24 15:47
@Author : zxr
@File : 6.9
@Software: GoLand
*/
package matins

import (
	"fmt"
)

//如何把十进制数分别以二进制和十六进制输出

//十进制转成二进制
func IntToBinAry(n int) {
	l := ""
	for n > 0 {
		m := n % 2
		n = n / 2
		l = fmt.Sprintf("%v%s", m, l)
	}
	fmt.Println(l)
}

//十进制转成十六进制
func IntTo16Binary(n int) {
	str := ""
	for n > 0 {
		m := n % 16
		n = n / 16
		str = fmt.Sprintf("%v%s", m, str)
	}
	fmt.Println("十六进制:", str)
}
