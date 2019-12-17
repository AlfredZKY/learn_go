package test

import (
	"errors"
	"flag"
	"fmt"
	"math"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object")
}

// 用hello生成问候语
func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf("Hello,%s!", name), nil
}

func introduce() string {
	return "Welcome to my Golang column."
}

// The Sieve of Eratosthens 爱拉托逊斯筛选法 思想：对于不超过n的每个非负整数P，删除2P, 3P...，当处理
// 完所有数之后，还没有被删除的就是素数。

// GetPrimes get all less than max prime number
func GetPrimes(max int) []int {
	if max <= 1 {
		return []int{}
	}
	marks := make([]bool, max)
	var count int
	squareRoot := int(math.Sqrt(float64(max)))
	//fmt.Println(squareRoot)
	for i := 2; i <= squareRoot; i++ {
		if marks[i] == false {
			for j := i * i; j < max; j += i {
				if marks[j] == false {
					marks[j] = true
					count++
				}
			}
		}
	}
	primes := make([]int, max-count)
	for i := 2; i < max; i++ {
		if marks[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}