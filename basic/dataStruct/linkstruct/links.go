package linkstruct

import (
	"errors"
	"fmt"
)

// Elem 别名
type Elem int

// LinkNode link elements
type LinkNode struct {
	Val  Elem
	Next *LinkNode
}

// New 初始化头节点
func New(i Elem) *LinkNode {
	return &LinkNode{i, nil}
}

// Insert 在指定位置插入元素
func (head *LinkNode) Insert(i int, e Elem) bool {
	p := head
	j := 1
	for nil != p && j < i {
		p = p.Next
		j++
	}
	if nil == p || j > i {
		fmt.Println("pls check i:", i)
		return false
	}
	// 构造节点
	s := &LinkNode{Val: e}
	s.Next = p.Next
	p.Next = s
	return true
}

// Append 直接插入到尾节点
func (head *LinkNode) Append(e Elem) {
	p := head
	// 不能指到尾节点，尾节点的上一个节点
	for p.Next != nil {
		p = p.Next
	}
	newNode := &LinkNode{Val: e}
	p.Next = newNode
}

// Traverse 遍历链表
func (head *LinkNode) Traverse() {
	point := head
	for point != nil {
		fmt.Printf("%d\t", point.Val)
		point = point.Next
	}
	fmt.Println("\n--------done--------")
}

// Delete 删除指定位置元素
func (head *LinkNode) Delete(i int) bool {
	p := head.Next
	j := 1
	for nil != p && j < i {
		p = p.Next
		j++
	}
	if nil == p || j > i {
		fmt.Println("pls check i:", i)
		return false
	}
	p.Next = p.Next.Next
	return true
}

// DeleteTail 删除元素
func (head *LinkNode) DeleteTail() {
	p := head
	for nil != p.Next.Next {
		p = p.Next

	}
	p.Next = p.Next.Next
}

// Get 获取指定位置的值
func (head *LinkNode) Get(i Elem) (Elem, error) {
	p := head
	if i < 0 {
		return -2, errors.New("not find")
	}
	var j Elem = 1
	for ; j < i; j++ {
		if nil == p {
			return -1, errors.New("not find")
		}
		p = p.Next
	}
	return p.Val, nil
}

// Loop 新建一个循环链表
func (head *LinkNode) Loop() *LinkNode {
	linkList := New(1)
	linkList1 := New(2)
	linkList2 := New(3)
	linkList3 := New(4)
	linkList4 := New(5)
	linkList5 := New(6)
	linkList.Next = linkList1
	linkList1.Next = linkList2
	linkList2.Next = linkList3
	linkList3.Next = linkList4
	linkList4.Next = linkList5
	linkList5.Next = linkList3
	return linkList
}
