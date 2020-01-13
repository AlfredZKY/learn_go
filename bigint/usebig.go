package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	den := big.NewRat(1, 1024)
	fmt.Printf("res is %v,Type is %T\n", den, den)
	res := big.NewInt(1).SetUint64(1000)
	fmt.Printf("res is %v,Type is %T\n", res, res)

	im := big.NewInt(math.MaxInt64)
	fmt.Printf("Big Int im is: %v\n", im)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(10)
	ret := big.NewInt(1)
	fmt.Printf("Big Int io is: %v\n", io)
	fmt.Printf("Big Int ip is: %v\n", ip)

	ret.Mul(io, ip)
	fmt.Printf("Big Int mul is : %v\n", ret)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
	byt := new(big.Rat).SetInt(big.NewInt(0).SetUint64(20971520))
	//byt := big.NewInt(1).SetUint64(1048576)
	byt.Mul(byt, den)
	byt.Mul(byt, den)
	f, _ := byt.Float64()
	fmt.Printf("byt is %v\n", f)
}
