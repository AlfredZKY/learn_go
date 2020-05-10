package common

import (
	"fmt"
	"testing"
)

func TestCreateFile(t *testing.T) {
	CreateFile("./op", "test.txt")
}

func testFunc() error {
	times := 10000000
	for i := 0; i < times; i++ {
		fmt.Println(i)
	}
	return nil
}

func TestExecute(t *testing.T) {
	Execute(testFunc,4)
}
