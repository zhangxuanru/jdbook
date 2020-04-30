/*
@Time : 2020/3/30 16:35
@Author : zxr
@File : 2.2.1_test
@Software: GoLand
*/
package queue

import (
	"fmt"
	"jdbook/offer/queue"
	"testing"
)

func TestLinkedQueue(t *testing.T) {
	LinkedQueue()
}

func LinkedQueue() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()

	linkedQueue := queue.NewLinkedQueue()
	linkedQueue.Push(1)
	linkedQueue.Push(2)
	fmt.Println("队列头元素为:", linkedQueue.GetFront())
	fmt.Println("队列尾元素为:", linkedQueue.GetBack())
	fmt.Println("队列大小为:", linkedQueue.Size())

}
