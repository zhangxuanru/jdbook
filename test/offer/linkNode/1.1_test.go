/*
@Time : 2020/3/26 14:44
@Author : zxr
@File : 1.1_test
@Software: GoLand
*/
package linkNode

import (
	"fmt"
	"jdbook/offer/linkNode"
	"testing"

	"github.com/isdamir/gotype"
)

//链表逆序
func TestReverse1(t *testing.T) {
	head := &gotype.LNode{}
	gotype.CreateNode(head, 8)
	reverse1 := linkNode.Reverse1(head)
	for reverse1 != nil {
		fmt.Println(reverse1.Data)
		reverse1 = reverse1.Next
	}
}

func TestReverse2(t *testing.T) {
	head := &gotype.LNode{}
	gotype.CreateNode(head, 8)
	linkNode.Reverse2(head)
	for head != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
}

func TestReverse3(t *testing.T) {
	head := &gotype.LNode{}
	gotype.CreateNode(head, 8)
	linkNode.Reverse3(head)
	for head != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
}

func TestReverse4(t *testing.T) {
	head := &gotype.LNode{}
	gotype.CreateNode(head, 8)
	linkNode.Reverse4(head)
	//for head != nil {
	//	fmt.Println(head.Data)
	//	head = head.Next
	//}
}
