package testfile

import "testing"

func TestTestSize(t*testing.T){
	const (
		t1 = 1 << 30
		t2 = 32<< 30
	)

	t.Log(t1,t2)
}