package usebufferio

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

const comment = "Package bufio implements buffered I/O. " +
	"It wraps an io.Reader or io.Writer object, " +
	"creating another object (Reader or Writer) that " +
	"also implements the interface but provides buffering and " +
	"some help for textual I/O."

func TestBufferReader(t *testing.T) {
	basicReader := strings.NewReader(comment)
	fmt.Printf("The size of basic reader:%d\n", basicReader.Size())
	fmt.Println()

	fmt.Println("New a buffered reader ...")
	reader1 := bufio.NewReader(basicReader)
	fmt.Printf("The default size of buffered reader:%d\n", reader1.Size())

	// 此时缓冲区中还没有写入数据
	fmt.Printf("The number of unread bytes in the buffer:%d\n", reader1.Buffered())
	fmt.Println()

	// 没有改变可读计数
	bytes, err := reader1.Peek(214)
	if err != nil {
		fmt.Printf("error:%v\n", err)
	}
	fmt.Printf("Peeked contents(%d):%q\n", len(bytes), bytes)
	fmt.Printf("The number of unread bytes in the buffer:%d\n", reader1.Buffered())

	buf1 := make([]byte, 7)
	n,err := reader1.Read(buf1)
	if err != nil{
		fmt.Printf("error:%v\n", err)
	}
	fmt.Printf("Read contents(%d):%q\n", n,buf1)
	fmt.Printf("The number of unread bytes in the buffer:%d\n", reader1.Buffered())
	fmt.Println()

	fmt.Printf("Reset the basic reader (size:%d) ...\n", len(comment))
	basicReader.Reset(comment)
	fmt.Printf("Reset the buffered reader (size :%d) ...\n", reader1.Size())
	reader1.Reset(basicReader)
	peekNum := len(comment) + 1
	fmt.Printf("Peek %d bytes ...\n", peekNum)
	bytes,err = reader1.Peek(peekNum)
	if err != nil{
		fmt.Printf("error :%v\n", err)
	}
	fmt.Printf("The number of peeked bytes:%d,%q\n", len(bytes),bytes)
	fmt.Println()

	fmt.Printf("Reset the basic reader (size:%d)...\n", len(comment))
	basicReader.Reset(comment)
	size := 300
	fmt.Printf("New a buffered reader with size %d ...\n", size)
	reader2 := bufio.NewReaderSize(basicReader, size)
	peekNum = size + 1
	fmt.Printf("Peek %d bytes ...\n", peekNum)
	bytes,err = reader2.Peek(peekNum)
	if err != nil{
		fmt.Printf("error:%v\n", err)
	}
	fmt.Printf("The number of peeked bytes:%d\n", len(bytes))
}
