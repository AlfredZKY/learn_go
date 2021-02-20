package main

import (
	"fmt"
)

func line(){
	var a int = 0
	switch a {
	case 0:
		{
			goto A
		A:
		}
	case 1:
		goto B
	B:
		; // ; 分号断行 否则编译不过 B:后必须跟一个表达式 
		  // lines.go:14:2: expected statement, found 'case'
		  // lines.go:20:1: expected declaration, found '}'
	case 2:
		goto C
	C:
	}
}

func adds(){
	var a = 0
	fmt.PrintLn(a++) // ./lines.go:29:15: syntax error: unexpected ++, expecting comma or )
	fmt.PrintLn(a--) // ./lines.go:30:15: syntax error: unexpected --, expecting comma or )
}

func main() {
	// line()
	adds()
	
}


// %!xxd 在vim中使用16进制查看内容