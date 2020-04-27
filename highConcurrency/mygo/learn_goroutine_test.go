package mygo


import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	time.Sleep(1 * time.Second)
}

func TestAsynExecute(t *testing.T) {
	AsynExecute()
}

func TestGorotineChannel(t *testing.T) {
	GorotineChannel()
}


func TestGorotineOrderExecute(t *testing.T) {
	GorotineOrderExecute()
}
