/*
@Time : 2020/4/24 15:48
@Author : zxr
@File : 6.9_test
@Software: GoLand
*/
package matins

import (
	"fmt"
	"jdbook/offer/matins"
	"log"
	"strconv"
	"testing"
)

//如何把十进制数分别以二进制和十六进制输出
func TestIntToBinAry(t *testing.T) {
	matins.IntToBinAry(18)
	formatInt := strconv.FormatInt(18, 2)
	fmt.Println("--", formatInt)
	fmt.Println(DecBin(18))

	//-------------十六进制
	fmt.Println("")

	matins.IntTo16Binary(18)
	s := strconv.FormatInt(18, 16)
	fmt.Println("--", s)
}

func DecBin(n int64) string {
	if n < 0 {
		log.Println("Decimal to binary error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	s := ""
	for q := n; q > 0; q = q / 2 {
		m := q % 2
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}
