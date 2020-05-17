package processtack

// 题目
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串

// IsValid 利用栈判断字符串是否有效
func IsValid(s string) bool {
	// 注意空字符串可被认为是有效字符串
	if s == "" {
		return true
	}

	// 定义一个栈 存储字符
	var stack []uint8

	// 定义一个字典存储括号对
	m := map[uint8]uint8{
		'}': '{',
		')': '(',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			// 入栈操作
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if m[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 题目
//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。

// MinStack 最小栈
type MinStack struct {
	stack []int
	min   []int
}

// Constructor ** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

// Push 入栈
func (minstack *MinStack) Push(x int) {
	minstack.stack = append(minstack.stack, x)

	if len(minstack.min) > 0 {
		min := minstack.min[len(minstack.min)-1]
		if min > x {
			minstack.min = append(minstack.min, x)
		} else {
			minstack.min = append(minstack.min, min)
		}
	} else {
		minstack.min = append(minstack.min, x)
	}
}

// Pop 出栈
func (minstack *MinStack) Pop() {
	minstack.stack = minstack.stack[0 : len(minstack.stack)-1]
	minstack.min = minstack.min[0 : len(minstack.min)-1]
}

// Top 栈顶元素
func (minstack *MinStack) Top() int {
	return minstack.stack[len(minstack.stack)-1]
}

// GetMin 获取最小值
func (minstack *MinStack) GetMin() int {
	return minstack.min[len(minstack.min)-1]
}

// 题目
// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// LargestRectangleArea 求出最大面积
func LargestRectangleArea(heights []int) int {
	// 前插一个辅助元素
	heights = append([]int{-2}, heights...)

	// 存储最大值
	res := 0
	size := len(heights)

	// 用栈存储heights 的下标,第一个下标为0
	s := make([]int, 1, size)
	i := 1

	for i < len(heights) {
		if heights[i] > heights[len(s)-1] {
			s = append(s, i)
			i++
			continue
		}
		res = max(res, heights[len(s)-1]*(i-s[len(s)-2]-1))
		s = s[:len(s)-1]
	}
	return res
}

// 题目
// 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 返回滑动窗口中的最大值。

// MaxSlidingWindow 滑动窗口每滑动一次取出最大值，返回数组
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	// 维护一个递减队列，存储的是下标
	var Q = make([]int, 0, len(nums))

	// 存储结果的数组
	res := make([]int, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		// 下标对应的最大值放入队列头
		for len(Q) != 0 && nums[i] >= nums[Q[len(Q)-1]] {
			// 队列递减
			Q = Q[:len(Q)-1]
		}

		// 当前元素下标入栈
		Q = append(Q, i)

		// // 当窗口离开first元素时，栈中第一个元素出栈
		// if Q[0] == i-k {
		// 	Q = Q[1:]
		// }

		// 窗口充满了k个元素时取出对队列头
		if i+1-k >= 0 {
			res[i+1-k] = nums[Q[0]]
		}
	}
	return res
}

// 题目
// 设计实现双端队列。
// 你的实现需要支持以下操作：

// MyCircularDeque(k)：构造函数,双端队列的大小为k。
// insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
// insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
// deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
// deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
// getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
// getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
// isEmpty()：检查双端队列是否为空。
// isFull()：检查双端队列是否满了。

// MyCircularDeque 节点
type MyCircularDeque struct {
    cache []int
    capacity int
    length int
    front int
    rear int
}

// ConstructorDCQ 创建节点
func ConstructorDCQ(k int) MyCircularDeque {
	return MyCircularDeque{
		cache: make([]int,k),
		capacity: k,
		front: 1,
	}
}

// InsertFront 插入头结点
func (dq *MyCircularDeque) InsertFront(value int) bool {
	if dq.length == dq.capacity{
		return false
	}
	dq.length++
	dq.front--
	if dq.front == -1{
		dq.front = dq.capacity-1
	}
	dq.cache[dq.front] = value
	return true
}

// InsertLast 插入尾结点
func (dq *MyCircularDeque) InsertLast(value int) bool {
	if dq.length == dq.capacity {
		return false
	}
	dq.length++
	dq.rear++
	if dq.rear == dq.capacity{
		dq.rear = 0
	}
	dq.cache[dq.rear] = value
	return true
}

// DeleteFront 删除头结点
func (dq *MyCircularDeque) DeleteFront() bool {
	if dq.length == 0 {
		return false
	}
	dq.length--
	dq.front++
	if dq.front == dq.capacity{
		dq.front = 0
	}
	return true
}

// DeleteLast 删除尾结点
func (dq *MyCircularDeque) DeleteLast() bool {
	if dq.length == 0 {
		return false
	}
	dq.length--
	dq.rear--
	if dq.rear == -1{
		dq.rear = dq.capacity -1
	}
	return true
}

// GetFront 获取头元素
func (dq *MyCircularDeque) GetFront() int {
	if dq.length == 0{
		return -1
	}
	return dq.cache[dq.front]
}

// GetRear 获取尾元素
func (dq *MyCircularDeque) GetRear() int {
	if dq.length == 0{
		return -1
	}
	return dq.cache[dq.rear]
}

// IsEmpty 判断队列是否为空
func (dq *MyCircularDeque) IsEmpty() bool {
	return dq.length == 0
}

// IsFull 判断队列是否为满
func (dq *MyCircularDeque) IsFull() bool {
	return dq.length == dq.capacity
}


