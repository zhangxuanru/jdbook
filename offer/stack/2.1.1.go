/*
@Time : 2020/3/30 14:57
@Author : zxr
@File : 2.1.1
@Software: GoLand
*/
package stack

import "github.com/isdamir/gotype"

//用链表方式实现栈
type LinkedStack struct {
	head *gotype.LNode
	size int
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{head: &gotype.LNode{}, size: 0}
}

func (p *LinkedStack) IsEmpty() bool {
	return p.head.Next == nil
}

func (p *LinkedStack) Size() int {
	return p.size
}

func (p *LinkedStack) Top() int {
	if p.head.Next != nil {
		return p.head.Next.Data.(int)
	}
	panic("stack is nil")
}

func (p *LinkedStack) Pop() int {
	tmp := p.head.Next
	if tmp != nil {
		p.head.Next = tmp.Next
		return tmp.Data.(int)
	}
	return -1
}

func (p *LinkedStack) Push(val int) {
	node := &gotype.LNode{Data: val, Next: p.head.Next}
	p.head.Next = node
	p.size++
}
