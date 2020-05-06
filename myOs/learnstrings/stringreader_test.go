package learnstrings

import (
	"strings"
	"testing"
)

var readers = strings.NewReader(
	"NewReader returns a new Reader reading from s. " +
		"It is similar to bytes.NewBufferString but more efficient and read-only.")

func TestStringReader(t *testing.T) {

	t.Logf("The size of reader is %d\n",readers.Size())	
	t.Logf("The reading index in reader %d\n",readers.Size()-int64(readers.Len()))

	buf1 := make([]byte,47)
	n,_ := readers.Read(buf1)
	t.Logf("%d bytes were read.(call Read)\n",n)
	t.Logf("The reading index in reader %d\n",readers.Size()-int64(readers.Len()))
	t.Log(string(buf1))
}

func TestStringReadAt(t*testing.T){
	buf := make([]byte,21)
	offset := int64(64)
	
	// 在64的位置写入21个字节
	n,_ := readers.ReadAt(buf,offset)
	t.Logf("%d bytes were read. (call ReadAt,offset:%d)\n",n,offset)
	t.Log(string(buf))
	t.Log(len(buf))
	t.Logf("The reading index in reader %d\n",readers.Size()-int64(readers.Len()))


}
