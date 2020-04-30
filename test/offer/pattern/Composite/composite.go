/*
@Time : 2020/4/29 18:27
@Author : zxr
@File : composite
@Software: GoLand
*/
package Composite

import "fmt"

//组合模式
type Component interface {
	Traverse()
}

type Leaf struct {
	value int
}

func NewLeaf(i int) *Leaf {
	return &Leaf{i}
}

func (l *Leaf) Traverse() {
	fmt.Println("Leaf value:", l.value)
}

type Cat struct {
	Name string
}

func NewCat(name string) *Cat {
	return &Cat{name}
}

func (c *Cat) Traverse() {
	fmt.Println("cat name:", c.Name)
}

type Composite struct {
	children []Component
}

func NewComposite() *Composite {
	return &Composite{children: make([]Component, 0)}
}
func (c *Composite) Add(coment Component) {
	c.children = append(c.children, coment)
}

func (c *Composite) Traverse() {
	for idx, _ := range c.children {
		c.children[idx].Traverse()
	}
}
