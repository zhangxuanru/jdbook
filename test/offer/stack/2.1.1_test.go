/*
@Time : 2020/3/30 15:19
@Author : zxr
@File : 2.1.1_test
@Software: GoLand
*/
package stack

import (
	"fmt"
	"jdbook/offer/stack"
	"testing"
)

func TestLinkedStack(t *testing.T) {
	linkedModel()
}

func linkedModel() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
	}()
	head := stack.NewLinkedStack()
	head.Push(1)
	fmt.Println("栈顶元素为：", head.Top())
	fmt.Println("栈大小为：", head.Size())
	head.Pop()
	fmt.Println("弹栈成功：", head.Size())

	head.Push(4)
	head.Push(5)
	head.Push(6)

	fmt.Println("栈大小为：", head.Size())
	pop1 := head.Pop()
	pop2 := head.Pop()
	pop4 := head.Pop()
	fmt.Println("栈Pop：", pop1, "---", pop2, "---", pop4)

}
