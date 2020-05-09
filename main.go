package main


import (
	// "time"
	"fmt"
	"learn_go/algorithm/proarrays"
	"learn_go/highConcurrency/channels"
	"learn_go/highConcurrency/useselect"
	"learn_go/learnfunc"
)

func f1() {
	for {
		fmt.Println("call f1...")
	}
}

func f2() {
	fmt.Println("call f2...")
}

type operate func(x int, y int) int

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := proarrays.ThreadNums(nums)
	fmt.Println(res)

	channels.UseChannelPanic()

	// // go f1()
	// go f2()
	// ch := make(chan int)
	// <- ch
	channels.UseChanSelect()
	channels.DetermineChanClose()
	channels.GetSumArray()
	useselect.CalcFiboni()

	var p learnfunc.Printer

	p = learnfunc.PrintToStd
	p("something")

	op := func(x, y int) int {
		return x + y
	}

	a, b := 1, 4
	result, err := learnfunc.Calculate(a, b, op)
	if err == nil {
		fmt.Println(result)
	}

}

