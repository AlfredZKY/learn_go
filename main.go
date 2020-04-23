package main

import (
	"fmt"
	"learn_go/algorithm"
)

func main(){
	nums := []int{-1,0,1,2,-1,-4}
	res := algorithm.ThreeSum(nums)
	fmt.Println(res)
}