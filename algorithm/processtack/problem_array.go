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
	// 在最前面插入-2，最后面插入-1 便于遍历
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)

	size := len(heights)
	// 定义一个栈
	s := make([]int,1,size)
	res := 0

	// 从第二位位置开始
	i := 1
	for i < len(heights) {
		// 和栈顶元素的下标比较
		if heights[i] > heights[s[len(s)-1]]{
			s = append(s,i)
			i++
			continue
		}
		res = max(res,heights[s[len(s)-1]]*(i-s[len(s)-2]-1))

		// 弹出栈顶元素
		s = s[:len(s)-1]
	}
	return res
}
