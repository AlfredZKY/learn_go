package main

import (
	"flag"
	"fmt"
	"strings"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	var a int = 1
	b := a << 30
	c := 64 << 30
	fmt.Println(b, c)
	fmt.Println(c / 1024 / 1024 / 1024)
	fmt.Println(strings.HasSuffix("NLT_abc", "abc"))
}
