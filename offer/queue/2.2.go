/*
@Time : 2020/3/30 15:40
@Author : zxr
@File : 2.2
@Software: GoLand
*/
package queue

import "fmt"

//用数组方式实现队列
type SliceQueue struct {
	arr   []int
	front int
	rear  int
}

func NewSliceQueue() *SliceQueue {
	return &SliceQueue{arr: make([]int, 0)}
}

func (p *SliceQueue) IsEmpty() bool {
	return p.front == p.rear
}

func (p *SliceQueue) Size() int {
	return p.rear - p.front
}

func (p *SliceQueue) GetFront() int {
	if p.IsEmpty() {
		fmt.Println("queue is nil")
	}
	return p.arr[p.front]
}

func (p *SliceQueue) GetBack() int {
	if p.IsEmpty() {
		fmt.Println("queue is nil")
	}
	return p.arr[p.rear-1]
}

func (p *SliceQueue) DeQueue() {
	if p.rear > p.front {
		p.rear--
		p.arr = p.arr[1:]
	}
}
func (p *SliceQueue) Push(val int) {
	p.arr = append(p.arr, val)
	p.rear++
}
func (p *SliceQueue) EnQueue(data interface{}) {

}
