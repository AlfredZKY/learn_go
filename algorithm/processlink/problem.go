package processlink

import "learn_go/dataStruct/linkstruct"

var (
	head = linkstruct.New(0)
)

// CreateSingleLink 创建一个新链表
func CreateSingleLink() *linkstruct.LinkNode {
	head := linkstruct.New(0)
	head.Append(1)
	head.Append(2)
	head.Append(4)
	head.Append(5)
	head.Append(6)
	return head
}

// 题目
// 反转一个单链表。
// 示例:
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL

// ReverserSignleListIter 翻转单链表 方法1 采用迭代的方法
func ReverserSignleListIter(head *linkstruct.LinkNode) *linkstruct.LinkNode {
	// 设置一个前置指针
	var pre *linkstruct.LinkNode = nil

	cur := head

	for cur != nil {
		// 保存下一跳的指针
		temp := cur.Next

		// 起初是斩断前程，后面就是指向前节点
		cur.Next = pre

		// 更新前驱指针
		pre = cur

		// 更新下一跳的指针，方便访问下一个指针,也就是恢复头指针
		cur = temp
	}
	return pre
}

// ReverserSignleListReceive 翻转单链表 方法2 采用递归的方法
func ReverserSignleListReceive(head *linkstruct.LinkNode) *linkstruct.LinkNode {
	return reverse(nil, head)
}

func reverse(prev, cur *linkstruct.LinkNode) *linkstruct.LinkNode {
	// 采用递归法，首先判断退出递归的条件
	if cur == nil {
		return prev
	}
	head := reverse(cur, cur.Next)
	cur.Next = prev
	return head
}

// 题目
// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// SwapPairsyOne 两两进行节点交换
func SwapPairsyOne(head *linkstruct.LinkNode) *linkstruct.LinkNode {
	// 创建一个前置指针
	var prev *linkstruct.LinkNode = &linkstruct.LinkNode{Val: -1, Next: head}

	// 在创建一个根节点，它的下一个节点就是头结点
	var hint *linkstruct.LinkNode = prev

	// go 代码赋值的特殊性 一行代码完成交换
	// 例如 prev->1->2->3->4->nil
	// 实际上是一次性跳过两个节点，这样才能交换 遍历链表时跳出循环的条件 prev的下一个不为空，并且下一个的下一个也不为空

	// 1. prev.Next 也就是1，赋值给2的下一个元素，也就是把1接在2的后面
	// 2. 把2的下一个元素，也就是3，赋值给1的下一个元素，也就是是把1接在3的后面
	// 3. 把1的下一个元素，也就是2，赋值给prev的下一个元素，因为2和1已经调换顺序，需要把prev重新连接在2的前面
	// 4. 把prev的下一个元素，也就是1，赋值给prev，注意这里的prev还没有发生改变，这一行是同时生效，所以现在prev的下一个元素还是1，所以看似把prev.Next赋值
	// 给了prev是一跳，其实是两跳，因为1和2换了位置，我们又跳到了1，所以是两跳

	for prev.Next != nil && prev.Next.Next != nil {
		prev.Next.Next.Next, prev.Next.Next, prev.Next, prev = prev.Next, prev.Next.Next.Next, prev.Next.Next, prev.Next
	}
	return hint.Next
}
