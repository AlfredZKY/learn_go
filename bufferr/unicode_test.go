package other

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestUnicode(t *testing.T) {
	str := "Go 爱好者 "
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))

	for i, c := range str {
		fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	}
	restype := reflect.TypeOf(str)
	fmt.Println(restype)

	var builder1 strings.Builder
	restype = reflect.TypeOf(builder1)
	fmt.Println(restype)
}

func TestString(t *testing.T) {
	var builder1 strings.Builder
	builder1.WriteString("hello world!!!")
	fmt.Printf("the first output(%d):\n%q\n", builder1.Len(), builder1.String())
	fmt.Println()

	builder1.WriteByte(' ')
	builder1.WriteString("it minimizes memory copying.the zero value is ready to use")
	builder1.Write([]byte{'\n', '\n'})
	builder1.WriteString("Do not copy a non-zero Builder.")
	fmt.Printf("the second output(%d):\n\"%s\"\n", builder1.Len(), builder1.String())

	fmt.Println("Grow the builder")
	builder1.Grow(10)
	fmt.Printf("The length of contents in the builder is %d.\n", builder1.Len())
	fmt.Println()

	fmt.Println("Reset the bulder")
	builder1.Reset()
	fmt.Printf("the third outpt(%d):\n%q\n", builder1.Len(), builder1.String())
}

func TestStringBuiler(t *testing.T) {
	var builder1 strings.Builder
	builder1.Grow(1)

	// 利用函数传值，
	f1 := func(b strings.Builder) {
		//b.Grow(1)
	}

	f1(builder1)

	ch1 := make(chan strings.Builder, 1)
	ch1 <- builder1

	// 利用通道传值
	builder2 := <-ch1
	//builder2.Grow(1)
	_ = builder2

	// 利用复制传值
	builder3 := builder1
	//builder3.Grow(1)

	_ = builder3

	f2 := func(bp *strings.Builder) {
		(*bp).Grow(1) //利用指针，这里虽然不会引发panic，但不是并发安全的
		builder4 := *bp
		//builder4.Grow(1)	// 这里还是会引发panic
		_ = builder4
	}

	f2(&builder1)

	builder1.Reset()
	builder5 := builder1
	builder5.Grow(1)
	builder5.WriteString("hello")
}

func TestStringsReader(t *testing.T) {
	reader1 := strings.NewReader(
		"NewReader returns a new Reader reading from s. " +
			"It is similar to bytes.NewBufferString but more efficient and read-only.")
	fmt.Printf("the size of reader:%d\n", reader1.Size())
	fmt.Printf("The reading index in reader:%d\n", (reader1.Size() - int64(reader1.Len())))

	buf1 := make([]byte, 47)
	n, _ := reader1.Read(buf1)
	fmt.Printf("%d bytes were read.(call Read)\n", n)
	fmt.Printf("The reading index in reader:%d\n", reader1.Size()-int64(reader1.Len()))
	fmt.Println()

	buf2 := make([]byte, 21)
	offset1 := int64(64)
	n, _ = reader1.ReadAt(buf2, offset1)
	fmt.Printf("%d bytes were read.(call ReadAt,offset:%d)\n", n, offset1)
	fmt.Printf("The reading index in reader:%d\n", reader1.Size()-int64(reader1.Len()))
	fmt.Println()

	offset2 := int64(17)
	expectedIndex := reader1.Size() - int64(reader1.Len()) + offset2
	fmt.Printf("Seek with offset %d and whence %d ...\n", offset2, io.SeekCurrent)
	readingIndex,_:=reader1.Seek(offset2,io.SeekCurrent)
	fmt.Printf("The reading index in reader:%d(returned by Seek)\n", readingIndex)
	fmt.Printf("The reading index in reader:%d(computed by me)\n", expectedIndex)

	n,_=reader1.Read(buf2)
	fmt.Printf("%d bytes were read.(call Read)\n", n)
	fmt.Printf("the reading index in reader: %d\n", reader1.Size()-int64(reader1.Len()))
}