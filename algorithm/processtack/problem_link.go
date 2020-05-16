package processtack

// MinStackLink 节点元素
type MinStackLink struct {
	Min  int
	Val  int
	Next *MinStackLink
}

// ConstructorLink 初始化节点
func ConstructorLink() MinStackLink {
	return MinStackLink{}
}

// Push 入栈操作
func (stack *MinStackLink) Push(x int) {
	// 构建节点
	temp := &MinStackLink{x, x, nil}
	if stack.Next == nil {
		stack.Next = temp
	} else if stack.Next.Min < x {
		temp.Min = stack.Next.Min
		temp.Next = stack.Next
		stack.Next = temp
	} else {
		temp.Next = stack.Next
		stack.Next = temp
	}
}

// Pop 出栈操作
func (stack *MinStackLink) Pop() {
	if stack.Next != nil {
		stack.Next = stack.Next.Next
	}
}

// Top 栈顶元素
func (stack *MinStackLink) Top() int {
	if stack.Next != nil {
		return stack.Next.Val
	}
	return 0
}

// GetMin 得到栈中最小元素
func (stack *MinStackLink) GetMin() int {
	if stack.Next != nil {
		return stack.Next.Min
	}
	return (1<<31) - 1
}
