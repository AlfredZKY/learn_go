package opp

import (
	"bytes"
	"math/rand"
	"strconv"
)

// CPUProFile show cpu info
func CPUProFile()error{
	max := 100000000
	var buf bytes.Buffer
	for j := 0;j<max;j++{
		num := rand.Int63n((int64(max)))
		str := strconv.FormatInt(num,10)
		buf.WriteString(str)
	}
	_ = buf.String()
	return nil
}