package processlink

import (
	"learn_go/dataStruct/linkstruct"
	"learn_go/datastruct/stacks"
)

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

// CreateCycle 创建一个新链表
func CreateCycle() *linkstruct.LinkNode {
	head := linkstruct.New(0)
	Cycle := head.Loop()
	return Cycle
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

		cur.Next = nil

		// 斩断当前的节点的连接
		cur.Next = pre

		// 更新前驱指针
		pre = cur

		// 更新当前节点
		cur = temp
	}
	return pre
}

// ReverserSignleListRecursive 翻转单链表 方法2 采用递归的方法
func ReverserSignleListRecursive(head *linkstruct.LinkNode) *linkstruct.LinkNode {
	return reverse(nil, head)
}

func reverse(prev, cur *linkstruct.LinkNode) *linkstruct.LinkNode {
	// 采用递归法，首先判断退出递归的条件
	if cur == nil {
		return prev
	}

	head := reverse(cur, cur.Next)

	// cur.Next这时候是尾节点 prev是前一个节点 相当于前后节点互换了指针方向
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
		// prev.Next.Next.Next, prev.Next.Next, prev.Next, prev = prev.Next, prev.Next.Next.Next, prev.Next.Next, prev.Next
		prev.Next.Next.Next, prev.Next.Next, prev.Next, prev = prev.Next, prev.Next.Next.Next, prev.Next.Next, prev.Next
	}
	return hint.Next
}

// SwapPairsyTwo 两两进行节点交换
func SwapPairsyTwo(head *linkstruct.LinkNode) *linkstruct.LinkNode {
	// 申请一个空节点
	p := new(linkstruct.LinkNode)
	q := p

	for head != nil && head.Next != nil {
		nextHead := head.Next.Next
		p.Next = head.Next
		head.Next = nil
		p.Next.Next = head
		p = p.Next.Next
		head = nextHead
	}

	if head != nil {
		p.Next = head
	}

	return q.Next
}

// SwapPairsyThree 递归交换
func SwapPairsyThree(head *linkstruct.LinkNode) *linkstruct.LinkNode {
	// 1.终止条件:当前没有节点或者只有一个节点，肯定就不需要交换了
	if head == nil || head.Next == nil {
		return head
	}

	// 2 节点交换 两个挨着的节点用于交换 head 和 head.Next
	firstNode, secondNode := head, head.Next
	firstNode.Next = SwapPairsyThree(secondNode.Next)
	secondNode.Next = firstNode
	return secondNode
}

// 题目
// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。 方法 有迭代 递归 栈等解决

//ReverseGroup 翻转指定个数的链表节点
func ReverseGroup(head *linkstruct.LinkNode, k int) *linkstruct.LinkNode {
	// 设置前驱节点
	dummy := &linkstruct.LinkNode{Val: -1, Next: head}
	prev := dummy
	cur := prev.Next

	for {
		// 先判断链表的节点是否满足要求
		n := k
		nextPart := cur
		for nextPart != nil && n > 0 {
			nextPart = nextPart.Next
			n--
		}

		if n > 0 {
			break
		} else {
			n = k
		}
		// 保存下一个pre节点
		nextPre := cur
		for n > 0 {
			// 保存下一跳元素
			temp := cur.Next

			// 斩断前尘
			cur.Next = nextPart

			// 下个头为当前元素
			nextPart = cur
			cur = temp
			n--
		}
		// n次翻转完毕
		prev.Next = nextPart

		// 设置下一个prev
		prev = nextPre
		cur = prev.Next
	}
	return dummy.Next
}

// ReverseGroupRecursive 递归
func ReverseGroupRecursive(head *linkstruct.LinkNode, k int) *linkstruct.LinkNode {
	cur := head
	count := 0
	for count != k && cur != nil {
		cur = cur.Next
		count++
	}
	if count == k {
		cur = ReverseGroupRecursive(cur, k)
		for count > 0 {

		}
	}
	return nil
}

// ReverseGroupStack 用栈
func ReverseGroupStack(head *linkstruct.LinkNode, k int) *linkstruct.LinkNode {
	dummy := &linkstruct.LinkNode{Val: -1, Next: nil}
	p := dummy
	_ = p
	tmp := head
	flagBool := false
	var stack stacks.Stack
	var temp *linkstruct.LinkNode
	for head != nil{
		count := k
		temp = tmp
		for count != 0 && tmp != nil {
			if tmp == nil && count > 0{
				flagBool = true
				break
			}
			value := &linkstruct.LinkNode{Val: -1, Next: nil}
			value.Val = tmp.Val
			stack.Push(value)
			tmp = tmp.Next
			head = head.Next
			count--
		}

		if flagBool {
			break
		}

		if count == 0 {
			for stack.Len() != 0 {
				value,_:= stack.Pop()
				// p.Next = interface{}(value).(*linkstruct.LinkNode)
				p.Next = value.(*linkstruct.LinkNode)
				p = p.Next
			}
		}
		
	}
	if tmp != nil {
		p.Next = temp
	}
	return dummy.Next
}

// 题目
// 给定一个链表，判断链表中是否有环。
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

// HasCyclePointer 判断一个指针是否有环
func HasCyclePointer(head *linkstruct.LinkNode) bool {
	if head == nil {
		return false
	}
	quick, slow := head, head

	for quick != nil && quick.Next != nil {
		if quick == slow {
			return true
		}
		quick = quick.Next.Next
		slow = slow.Next
	}
	return false
}

// HasCycleMap 利用哈希表存储
func HasCycleMap(head *linkstruct.LinkNode) bool {
	if head == nil {
		return false
	}

	hash := make(map[*linkstruct.LinkNode]linkstruct.Elem)

	for head != nil {
		if _, exist := hash[head]; exist {
			return true
		}
		hash[head] = head.Val
		head = head.Next
	}
	return false
}

// HasCycleSomeVal 把链表中节点的值全部置为全部一样值
func HasCycleSomeVal(head *linkstruct.LinkNode) bool {
	if head == nil {
		return false
	}

	for head != nil {
		if head.Val == 9999999 {
			return true
		}
		head.Val = 9999999
		head = head.Next
	}
	return false
}


// 题目
// 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
// 说明：不允许修改给定的链表。

// DetectCycle 判断一个链表是否有环，并返回环的位置 利用快慢指针解题
func DetectCycle(head *linkstruct.LinkNode)*linkstruct.LinkNode{
	if head == nil && head.Next == nil {
		return nil
	}
	CycleFlag := false

	// 快指针一次走两步，慢指针一次走一步
	quick,slow := head,head
	for quick != nil && quick.Next != nil {
		quick ,slow= quick.Next.Next,slow.Next
		if quick == slow {
			// 证明有环
			CycleFlag = true
			break
		}
	}
	if !CycleFlag {
		return nil
	}
	quick = head
	for quick != slow {
		quick,slow = quick.Next,slow.Next
	}
	return slow
}