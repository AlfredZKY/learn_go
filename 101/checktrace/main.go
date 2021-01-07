// example1.go
package main

import "math/rand"

func f1(s []int) {
	_ = s[0] // 第5行： 需要边界检查
	_ = s[1] // 第6行： 需要边界检查
	_ = s[2] // 第7行： 需要边界检查
}

func f2(s []int) {
	_ = s[2] // 第11行： 需要边界检查
	_ = s[1] // 第12行： 边界检查消除了！
	_ = s[0] // 第13行： 边界检查消除了！
}

func f3(s []int, index int) {
	_ = s[index] // 第17行： 需要边界检查
	_ = s[index] // 第18行： 边界检查消除了！
}

func f4(a [5]int) {
	_ = a[4] // 第22行： 边界检查消除了！
}

func f5(s []int) {
	for i := range s {
		_ = s[i]
		_ = s[i:len(s)]
		_ = s[:i+1]
	}
}

func f6(s []int) {
	for i := 0; i < len(s); i++ {
		_ = s[i]
		_ = s[i:len(s)]
		_ = s[:i+1]
	}
}

func f7(s []int) {
	for i := len(s) - 1; i >= 0; i-- {
		_ = s[i]
		_ = s[i:len(s)]
	}
}

func f8(s []int, index int) {
	if index >= 0 && index < len(s) {
		_ = s[index]
		_ = s[index:len(s)]
	}
}

func f9(s []int) {
	if len(s) > 2 {
	    _, _, _ = s[0], s[1], s[2]
	}
}



func fa() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	index := rand.Intn(7)
	_ = s[:index] // 第9行： 需要边界检查
	_ = s[index:] // 第10行： 边界检查消除了！
}

func fb(s []int, index int) {
	_ = s[:index] // 第14行： 需要边界检查
	_ = s[index:] // 第15行： 需要边界检查（不够智能？）
}

func fc() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	s = s[:4]
	index := rand.Intn(7)
	_ = s[:index] // 第22行： 需要边界检查
	_ = s[index:] // 第23行： 需要边界检查（不够智能？）
}

func main() {}