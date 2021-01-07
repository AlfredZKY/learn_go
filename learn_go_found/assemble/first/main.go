package main

func main() {
	a := 3
	b := 2
	returnTwo(a, b)
}

// BP：基址指针寄存器(extended base pointer)，也叫帧指针，存放着一个指针，表示函数栈开始的地方。
// SP：栈指针寄存器(extended stack pointer)，存放着一个指针，存储的是函数栈空间的栈顶，也就是函数栈空间分配结束的地方，注意这里是硬件寄存器，不是Plan9中的伪寄存器。
// BP 与 SP 放在一起，一个表示开始（栈顶）、一个表示结束（栈低）。



func returnTwo(a, b int) (c, d int) {
	tmp := 1 // 这一行的主要目的是保证栈桢不为0，方便分析
	c = a + b
	d = b - tmp
	return
}
