package linkstruct

import (
	"errors"
	"fmt"
)

// Elem 别名
type Elem int

// LinkNode link elements
type LinkNode struct {
	val  Elem
	next *LinkNode
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
		p = p.next
		j++
	}
	if nil == p || j > i {
		fmt.Println("pls check i:", i)
		return false
	}
	// 构造节点
	s := &LinkNode{val: e}
	s.next = p.next
	p.next = s
	return true
}

// Append 直接插入到尾节点
func (head *LinkNode) Append(e Elem) {
	p := head
	// 不能指到尾节点，尾节点的上一个节点
	for p.next != nil {
		p = p.next
	}
	newNode := &LinkNode{val: e}
	p.next = newNode
}

// Traverse 遍历链表
func (head *LinkNode) Traverse() {
	point := head
	for point != nil {
		fmt.Printf("%d\t", point.val)
		point = point.next
	}
	fmt.Println("\n--------done--------")
}

// Delete 删除指定位置元素
func (head *LinkNode) Delete(i int) bool {
	p := head.next
	j := 1
	for nil != p && j < i {
		p = p.next
		j++
	}
	if nil == p || j > i {
		fmt.Println("pls check i:", i)
		return false
	}
	p.next = p.next.next
	return true
}

// DeleteTail 删除元素
func (head *LinkNode) DeleteTail() {
	p := head
	for nil != p.next.next {
		p = p.next

	}
	p.next = p.next.next
}

// Get 获取指定位置的值
func (head *LinkNode) Get(i int) (Elem, error) {
	p := head
	if i < 0 {
		return -2, errors.New("not find")
	}
	for j := 1; j < i; j++ {
		if nil == p {
			return -1, errors.New("not find")
		}
		p = p.next
	}
	return p.val, nil
}

// Loop 新建一个循环链表
func (head *LinkNode) Loop() *LinkNode {
	linkList := New(1)
	linkList1 := New(2)
	linkList2 := New(3)
	linkList3 := New(4)
	linkList4 := New(5)
	linkList5 := New(6)
	linkList.next = linkList1
	linkList1.next = linkList2
	linkList2.next = linkList3
	linkList3.next = linkList4
	linkList4.next = linkList5
	linkList5.next = linkList3
	return linkList
}
