package testfile

import (
	"fmt"
	"io"
	"math/rand"
	"testing"
)

func TestTestSize(t *testing.T) {
	const (
		t1 = 1 << 30
		t2 = 32 << 30
	)

	t.Log(t1, t2)
}

func TestIOReader(t *testing.T) {
	parts := uint64(2)
	piece := 34359738368
	readers := make([]io.Reader, parts)
	for i := range readers {
		readers[i] = io.LimitReader(rand.New(rand.NewSource(42+int64(i))), int64(piece))
		p := make([]byte, piece)
		res, err := readers[i].Read(p)
		if err != io.EOF {
			t.Logf("%v", res)
		}
	}
}

func testmultiArgs(sizes ...uint64) {
	for i, size := range sizes {
		fmt.Println(i, size)
	}
	fmt.Println(len(sizes))
}

func TestMultiArgs(t *testing.T) {
	sizes := uint64(10245455)
	testmultiArgs(sizes)

	t.Log(io.LimitReader(rand.New(rand.NewSource(42)), int64(sizes)))
}
