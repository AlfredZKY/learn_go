package mystring

import "testing"

func TestString(t*testing.T){
	var s string
	t.Log("x"+s+"x")	// 初始化为默认零值""

	s = "hello"
	//s[1] = "3"	//string是不可变的byte slice
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(len(s)) // 是byte数组

	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t*testing.T){
	s:="中华人名共和国"
	for _,c:=range s{
		t.Logf("%[1]c %[1]d %[1]x",c)
	}
}


func checkHostName(hostname string) bool {
	if hostname == "miner" ||  hostname == "c2_1" || hostname == "c2_2" || hostname == "c2_3" ||hostname == "c2_4" || hostname == "c2_5" || hostname == "c2_6"{
		return true
	}
	return false
}

func TestStringFind(t*testing.T){
	var srcStr = [...]string{"pc1","pc2","c1","c2_"}
	t.Log(srcStr)
	for _,v := range srcStr {
		res := checkHostName(v)
		if res {
			t.Log(v)
		}
	}
}