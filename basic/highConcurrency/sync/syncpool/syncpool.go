package syncpool

import (
	"bytes"
	"sync"
)

// bufPool 代表存放数据块缓冲区的临时对象池
var bufPool sync.Pool

// delimiter 代表预定义的定界符
var delimiter = byte('\n')

// Buffer 代表了一个简易的数据块缓冲区的接口
type Buffer interface {
	Delimiter() byte                   // Delimiter 用于获取数据块之间的界定符
	Write(contents string) (err error) // 用于写一个数据块
	Read() (contents string,err error)  // 用于读取一个数据块
	Free()                             //释放当前的缓冲区
}

// myBuffer 代表了数据块缓冲区的实现
type myBuffer struct {
	buf       bytes.Buffer
	delimiter byte
}

func (b *myBuffer) Delimiter() byte {
	return b.delimiter
}

func (b *myBuffer) Write(contents string) (err error) {
	if _, err = b.buf.WriteString(contents); err != nil {
		return
	}
	return b.buf.WriteByte(b.delimiter)
}

func (b*myBuffer)Read()(contents string,err error){
	return b.buf.ReadString(b.delimiter)
}

func (b*myBuffer) Free(){
	bufPool.Put(b)
}

// GetBuffer 用于获取一个数据块缓冲区
func GetBuffer()Buffer{
	// 将Buffer 转化为Get()类型，相当于接口绑定。
	return bufPool.Get().(Buffer)
}

func init(){
	bufPool = sync.Pool{
		New:func()interface{} {
			return &myBuffer{delimiter:delimiter}
		},
	}
}