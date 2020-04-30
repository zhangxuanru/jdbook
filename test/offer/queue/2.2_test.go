/*
@Time : 2020/3/30 15:49
@Author : zxr
@File : 2.2_test
@Software: GoLand
*/
package queue

import (
	"fmt"
	"jdbook/offer/queue"
	"testing"
)

func TestSliceQueue(t *testing.T) {
	SliceQueue()
}

func SliceQueue() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
	}()
	fmt.Println("slice 构建队列")
	sliceQueue := queue.NewSliceQueue()

	sliceQueue.Push(1)
	sliceQueue.Push(2)
	sliceQueue.Push(3)
	fmt.Println("队列首元素:", sliceQueue.GetFront())
	fmt.Println("队列尾元素:", sliceQueue.GetBack())
	fmt.Println("队列大小:", sliceQueue.Size())
}
