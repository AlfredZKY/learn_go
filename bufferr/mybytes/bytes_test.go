package mybytes

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBytesFirst(t *testing.T) {
	var buffer1 bytes.Buffer
	contents := "Simple byte buffer for marshaling data."
	fmt.Printf("Write contents %q...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer:%d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer:%d\n", buffer1.Cap())
	fmt.Println()

	p1 := make([]byte, 7)
	n, _ := buffer1.Read(p1)
	fmt.Printf("%d bytes were read.(call Read)\n", n)
	fmt.Printf("The length of buffer:%d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer:%d\n", buffer1.Cap())
}

func TestBytesSecond(t *testing.T) {
	var contents string
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("The length of new buffer with contents %q:%d\n", contents, buffer1.Len())
	fmt.Printf("The capacity of new buffer with contents %q:%d\n", contents, buffer1.Cap())
	fmt.Println()

	contents = "123456"
	fmt.Printf("Writing contests %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer:%d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer:%d\n", buffer1.Cap())
	fmt.Println()

	contents = "78"
	fmt.Printf("Writing contests %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer:%d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer:%d\n", buffer1.Cap())
	fmt.Println()

	contents = "89"
	fmt.Printf("Writing contests %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer:%d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer:%d\n", buffer1.Cap())
	fmt.Println()

	contents = "abcdefghijk"
	buffer2 := bytes.NewBufferString(contents)
	fmt.Printf("The length of buffer:%d\n", buffer2.Len())
	fmt.Printf("The capacity of buffer:%d\n", buffer2.Cap())
	fmt.Println()

	n := 10
	fmt.Printf("Grow the buffer with %d ...\n", n)
	buffer2.Grow(n)
	fmt.Printf("The length of buffer:%d\n", buffer2.Len())
	fmt.Printf("The capacity of buffer:%d\n", buffer2.Cap())
	fmt.Println()

	var buffer3 bytes.Buffer
	fmt.Printf("The length of buffer:%d\n", buffer3.Len())
	fmt.Printf("The capacity of buffer:%d\n", buffer3.Cap())
	fmt.Println()

	contents = "xyz"
	fmt.Printf("write contes %q ...\n", contents)
	buffer3.WriteString(contents)
	fmt.Printf("The length of buffer:%d\n", buffer3.Len())
	fmt.Printf("The capacity of buffer:%d\n", buffer3.Cap())
	fmt.Println()
}

func TestBytesThird(t *testing.T) {
	contents := "ab"
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("The capacity of new buffer with contents %q:%d\n", contents, buffer1.Cap())
	fmt.Println()

	unreadbytes := buffer1.Bytes()
	fmt.Printf("The unread bytes of the buffer:%v\n", unreadbytes)
	fmt.Println()

	contents = "cdefg"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The capacity of buffer:%d\n", buffer1.Cap())
	fmt.Println("---------------------------------------------")
	unreadbytes = unreadbytes[:cap(unreadbytes)]
	fmt.Printf("The unread bytes of the buffer:%v\n", unreadbytes)
	fmt.Println(buffer1.String())

	value:=byte('X')
	fmt.Printf("Set a byte in the unread bytes to %v ...\n", value)
	unreadbytes[len(unreadbytes)-2] = value
	fmt.Printf("The unread bytes of the buffer:%v\n", buffer1.Bytes())
	fmt.Println()

	fmt.Println("---------------------------------------------")
	contents = "hijklmn"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The unread bytes of the buffer:%d\n", buffer1.Cap())
	fmt.Println()
	unreadbytes = unreadbytes[:cap(unreadbytes)]
	fmt.Printf("The unread bytes of the buffer:%v\n", unreadbytes)
	fmt.Println(buffer1.String())
	fmt.Print("\n\n")
}

func TestBytesNext(t*testing.T){
	contents :="12"
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("The capacity of new buffer with contents %q: %d\n", contents,buffer1.Cap())
	fmt.Println()

	nextBytes := buffer1.Next(2)
	fmt.Printf("The next bytes of the buffer: %v\n",nextBytes)
	fmt.Println()

	contents = "34567"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The capacity of buffer:%d\n", buffer1.Cap())
	fmt.Println()

	nextBytes = nextBytes[:cap(nextBytes)]
	fmt.Printf("The next bytes of the buffer:%v\n",nextBytes)
	fmt.Println()

	value := byte('X')
	fmt.Printf("The next bytes of the buffer:%v ...\n",value)
	nextBytes[len(nextBytes)-2]=value
	fmt.Printf("The unread bytes of the buffer:%v\n", buffer1.Bytes())
	fmt.Println()

	contents = "89101112"
	fmt.Printf("Writing contents %q ...\n",contents)
	buffer1.WriteString(contents)
	fmt.Printf("The capacity of buffer:%d\n", buffer1.Cap())
	fmt.Println()

	nextBytes = nextBytes[:cap(nextBytes)]
	fmt.Printf("The next bytes of the buffer:%v\n", nextBytes)
}