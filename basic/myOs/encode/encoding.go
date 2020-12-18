package encode

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"strconv"
)

// BinaryReadParctice binary data read
func BinaryReadParctice() {
	//参数列表：
	// 1）r  可以读出字节流的数据源
	// 2）order  特殊字节序，包中提供大端字节序和小端字节序
	// 3）data  需要解码成的数据
	// 返回值：error  返回错误
	// 功能说明：Read从r中读出字节数据并反序列化成结构数据。data必须是固定长的数据值或固定长数据的slice。从r中读出的数据可以使用特殊的 字节序来解码，并顺序写入value的字段。当填充结构体时，使用(_)名的字段讲被跳过。

	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		log.Fatalln("Binary.Read failed:", err)
	}
	fmt.Println(pi)
}

// BinaryWriteParctice binary data write
func BinaryWriteParctice() {
	//参数列表：
	// 1）w  可写入字节流的数据
	// 2）order  特殊字节序，包中提供大端字节序和小端字节序
	// 3）data  需要解码的数据
	// 返回值：error  返回错误
	// 功能说明：
	// Write讲data序列化成字节流写入w中。data必须是固定长度的数据值或固定长数据的slice，或指向此类数据的指针。写入w的字节流可用特殊的字节序来编码。另外，结构体中的(_)名的字段讲忽略。

	buf := new(bytes.Buffer)
	pi := math.Pi

	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.Bytes())
}

// BinarySizeParctice binary data size  Size讲返回数据系列化之后的字节长度，数据必须是固定长数据类型、slice和结构体及其指针。
func BinarySizeParctice() {
	//参数列表：v  需要计算长度的数据
	// 返回值：int 数据序列化之后的字节长度
	// 功能说明：
	// Size讲返回数据系列化之后的字节长度，数据必须是固定长数据类型、slice和结构体及其指针。

	var a int
	p := &a
	b := [10]int64{1}
	s := "asda"
	bs := make([]byte, 10)

	fmt.Println(binary.Size(a))  //-1
	fmt.Println(binary.Size(p))  //-1
	fmt.Println(binary.Size(b))  //80
	fmt.Println(binary.Size(s))  //-1
	fmt.Println(binary.Size(bs)) //10
}

// BinaryPutUvarintParctice put uint64 in buf
func BinaryPutUvarintParctice() {
	// 参数列表：
	// 1）buf  需写入的缓冲区
	// 2）x  uint64类型数字
	// 返回值：
	// 1）int  写入字节数。
	// 2）panic  buf过小。
	// 功能说明：
	// PutUvarint主要是讲uint64类型放入buf中，并返回写入的字节数。如果buf过小，PutUvarint将抛出panic。
	u16 := 1234
	u64 := 0x1020304040302010
	sbuf := make([]byte, 4)
	buf := make([]byte, 10)

	ret := binary.PutUvarint(sbuf, uint64(u16))
	fmt.Println(ret, len(strconv.Itoa(u16)), sbuf)

	// 会发现转成二进制来传输数据，比直接转字符串之后转[]byte这种方式传更节省传输空间
	ret = binary.PutUvarint(buf, uint64(u64))
	fmt.Println(ret, len(strconv.Itoa(u64)), buf)
}

// BinaryPutvarintParctice put int64 in buf
func BinaryPutvarintParctice() {
	//参数列表：
	// 1）buf  需要写入的缓冲区
	// 2）x int64类型数字
	// 返回值：
	// 1）int  写入字节数
	// 2）panic  buf过小
	// 功能说明：
	// PutVarint主要是讲int64类型放入buf中，并返回写入的字节数。如果buf过小，PutVarint将抛出panic。
	i16 := 1234
	i64 := -1234567890
	sbuf := make([]byte, 4)
	buf := make([]byte, 10)

	ret := binary.PutVarint(sbuf, int64(i16))
	fmt.Println(ret, len(strconv.Itoa(i16)), sbuf)

	ret = binary.PutVarint(buf, int64(i64))
	fmt.Println(ret, len(strconv.Itoa(i64)), buf)
}

// BinaryUvarintParctice get uint64 data from buf
func BinaryUvarintParctice() {
	// 参数列表：buf  需要解码的缓冲区
	// 返回值：
	// 1）uint64  解码的数据。
	// 2）int  解析的字节数。
	// 功能说明:
	// Uvarint是从buf中解码并返回一个uint64的数据,及解码的字节数(>0)。如果出错,则返回数据0和一个小于等于0的字节数n,其意义为:
	// 1)n == 0: buf太小
	// 2)n < 0: 数据太大,超出uint64最大范围,且-n为已解析字节数
	sbuf := []byte{}
	buf := []byte{144, 192, 192, 132, 136, 140, 144, 16, 0, 1, 1}
	bbuf := []byte{144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1}

	num, ret := binary.Uvarint(sbuf)
	fmt.Println(num, ret)

	num, ret = binary.Uvarint(buf)
	fmt.Println(num, ret)

	num, ret = binary.Uvarint(bbuf)
	fmt.Println(num, ret)
}

// BinaryPutVarintParctice get int64 data from buf
func BinaryPutVarintParctice() {
	// 	参数列表: buf  需要解码的缓冲区
	// 返回值:
	// 1) int64 解码的数据
	// 2) int  解析的字节数
	// 功能说明:
	// Varint是从buf中解码并返回一个int64的数据,及解码的字节数(>0).如果出错,则返回数据0和一个小于等于0的字节数n,其意义为:
	// 1) n == 0: buf太小
	// 2) n < 0: 数据太大,超出64位,且-n为已解析字节数
	var sbuf []byte
	var buf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 16, 0, 1, 1}
	var bbuf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1}

	num, ret := binary.Varint(sbuf)
	fmt.Println(num, ret) //0 0

	num, ret = binary.Varint(buf)
	fmt.Println(num, ret) //580990878187261960 9

	num, ret = binary.Varint(bbuf)
	fmt.Println(num, ret) //0 -11
}

// BinaryReadUvarintParctice read uint64 data from buf
func BinaryReadUvarintParctice() {
	// 参数列表:
	// 返回值:
	// 1) uint64  解析出的数据
	// 2) error  返回的错误
	// 功能说明:
	// ReadUvarint从r中解析并返回一个uint64类型的数据及出现的错误.
	// 功能说明:
	// ReadUvarint从r中解析并返回一个uint64类型的数据及出现的错误.
	var sbuf []byte
	var buf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 16, 0, 1, 1}
	var bbuf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1}

	num, err := binary.ReadUvarint(bytes.NewBuffer(sbuf))
	fmt.Println(num, err) //0 EOF

	num, err = binary.ReadUvarint(bytes.NewBuffer(buf))
	fmt.Println(num, err) //1161981756374523920 <nil>

	num, err = binary.ReadUvarint(bytes.NewBuffer(bbuf))
	fmt.Println(num, err) //4620746270195064848 binary: varint overflows a 64-bit integer

}

// BinaryReadVarintParctice read int64 data from buf
func BinaryReadVarintParctice() {
	// 	参数列表: r  实现ByteReader接口的对象
	// 返回值:
	// 1) int64  解析出的数据
	// 2) error  返回的错误
	// 功能说明:
	// ReadVarint从r中解析并返回一个int64类型的数据及出现的错误.
	var sbuf []byte
	var buf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 16, 0, 1, 1}
	var bbuf []byte = []byte{144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1}

	num, err := binary.ReadVarint(bytes.NewBuffer(sbuf))
	fmt.Println(num, err) //0 EOF

	num, err = binary.ReadVarint(bytes.NewBuffer(buf))
	fmt.Println(num, err) //580990878187261960 <nil>

	num, err = binary.ReadVarint(bytes.NewBuffer(bbuf))
	fmt.Println(num, err) //2310373135097532424 binary: varint overflows a 64-bit integer
}
