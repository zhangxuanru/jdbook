/*
@Time : 2020/3/30 15:40
@Author : zxr
@File : 2.2
@Software: GoLand
*/
package queue

import "fmt"

//用数组方式实现队列
type SliceDataQueue struct {
	arr   []interface{}
	front int
	rear  int
}

func NewSliceDataQueue() *SliceDataQueue {
	return &SliceDataQueue{arr: make([]interface{}, 0)}
}

func (p *SliceDataQueue) IsEmpty() bool {
	return p.front == p.rear
}

func (p *SliceDataQueue) Size() int {
	return p.rear - p.front
}

func (p *SliceDataQueue) GetFront() interface{} {
	if p.IsEmpty() {
		fmt.Println("queue is nil")
	}
	return p.arr[p.front]
}

func (p *SliceDataQueue) GetBack() interface{} {
	if p.IsEmpty() {
		fmt.Println("queue is nil")
	}
	return p.arr[p.rear-1]
}

func (p *SliceDataQueue) DeQueue() interface{} {
	var data interface{}
	if p.rear > p.front {
		data = p.arr[0]
		p.rear--
		p.arr = p.arr[1:]
	}
	return data
}

func (p *SliceDataQueue) Push(val int) {
	p.arr = append(p.arr, val)
	p.rear++
}

func (p *SliceDataQueue) EnQueue(data interface{}) {
	p.arr = append(p.arr, data)
	p.rear++
}
