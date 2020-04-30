/*
@Time : 2020/3/30 14:42
@Author : zxr
@File : 2.1_test
@Software: GoLand
*/
package stack

import (
	"fmt"
	"jdbook/offer/stack"
	"testing"
)

func TestStack1(t *testing.T) {
	sliceMode()
}
func sliceMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
	fmt.Println("slice init stack")
	sliceStack := stack.NewSliceStack()
	sliceStack.Push(1)
	fmt.Println("栈顶元素为：", sliceStack.Top())
	fmt.Println("栈的大小为:", sliceStack.Size())
	sliceStack.Top()
	fmt.Println("弹栈成功:", sliceStack.Size())
	sliceStack.Top()
}
