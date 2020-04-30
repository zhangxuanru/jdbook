/*
@Time : 2020/4/15 15:29
@Author : zxr
@File : 5.3.1_test
@Software: GoLand
*/
package string

import (
	"fmt"
	string2 "jdbook/offer/string"
	"testing"
)

//如何实现单词反转
func TestReverseWorld(t *testing.T) {
	str := "how are you"
	world := string2.ReverseWorld(str)
	fmt.Println("world:", world)
}
