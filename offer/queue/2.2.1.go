/*
@Time : 2020/3/30 16:03
@Author : zxr
@File : 2.2.1
@Software: GoLand
*/
package queue

import (
	"fmt"

	"github.com/isdamir/gotype"
)

//用链表实现队列
type LinkedQueue struct {
	head *gotype.LNode
	end  *gotype.LNode
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{}
}

func (p *LinkedQueue) IsEmpty() bool {
	return p.head == nil
}

func (p *LinkedQueue) Size() int {
	size := 0
	node := p.head
	for node != nil {
		node = node.Next
		size++
	}
	return size
}

func (p *LinkedQueue) Push(val int) {
	node := &gotype.LNode{Data: val}
	if p.head == nil {
		p.head = node
		p.end = node
	} else {
		p.end.Next = node
		p.end = node
	}
}

func (p *LinkedQueue) Pop() {
	if p.head == nil {
		fmt.Println("queue is nil")
		return
	}
	p.head = p.head.Next
	if p.head == nil {
		p.end = nil
	}
}

func (p *LinkedQueue) GetFront() int {
	if p.head == nil {
		fmt.Println("queue is nil")
		return -1
	}
	return p.head.Data.(int)
}

func (p *LinkedQueue) GetBack() int {
	if p.end == nil {
		fmt.Println("queue is nil")
		return -1
	}
	return p.end.Data.(int)
}
