package constant

import (
	"testing"
)

const(
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const(
	Readable = 1<< iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T){
	t.Log(Monday,Tuesday,Wednesday)
}

func TestConstantIo(t *testing.T){
	a:=7 //0111
	a = 1 //0
	t.Log(a&Readable == Readable,a&Writable == Writable,a&Executable==Executable)
}