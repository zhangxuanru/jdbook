/*
@Time : 2020/3/30 14:41
@Author : zxr
@File : 2.1
@Software: GoLand
*/
package stack

//用数组方式实现栈
type SliceStack struct {
	arr       []int
	stackSize int
}

func NewSliceStack() *SliceStack {
	return &SliceStack{arr: make([]int, 0)}
}

func (p *SliceStack) IsEmpty() bool {
	return p.stackSize == 0
}

func (p *SliceStack) Size() int {
	return p.stackSize
}

func (p *SliceStack) Top() int {
	if p.IsEmpty() {
		panic("stack is nill")
	}
	return p.arr[p.stackSize-1]
}

func (p *SliceStack) Pop() int {
	if p.stackSize > 0 {
		p.stackSize--
		ret := p.arr[p.stackSize]
		p.arr = p.arr[:p.stackSize]
		return ret
	}
	panic("stack is nil")
}

func (p *SliceStack) Push(val int) {
	p.arr = append(p.arr, val)
	p.stackSize++
}
