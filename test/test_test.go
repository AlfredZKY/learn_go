package test

import (
	"strconv"
	"fmt"
	"testing"
	"github.com/multiformats/go-varint"
)

func TestHello(t *testing.T) {
	var name string
	greeting, err := hello(name)
	if err == nil {
		t.Errorf("The error is nil,but it should not be.(name=%q)", name)
	}
	if greeting != "" {
		t.Errorf("Nonempty greeting,but it should not be.(name=%q)", name)
	}
	name = "Robert"
	greeting, err = hello(name)
	if err != nil {
		t.Errorf("The error is not nil,but it should not be.(name=%q)", name)
	}
	if greeting == "" {
		t.Errorf("Empty greeting,but it should not be.(name=%q)", name)
	}

	expected := fmt.Sprintf("Hello,%s!", name)
	if greeting != expected {
		t.Errorf("The actual greeting %q is not the expected.(name=%q)", greeting, expected)
	}
}

func TestIntroduce(t *testing.T) {
	intro := introduce()
	expected := "Welcome to my Golang column."
	if intro != expected {
		t.Errorf("The actual introduce %q is not the expected.", intro)
	}
	t.Logf("The expected introduce is %q.\n", expected)
}

func TestFail(t *testing.T) {
	//t.Fail()
	t.Log("testing Fail")
	//t.FailNow()
	//t.Error("print info fail")
	t.Log("testing FailNow")
}


// Protocol represents which protocol an address uses.
type Protocol = byte

type Address struct{ str string }

const (
	// ID represents the address ID protocol.
	ID Protocol = iota
	// SECP256K1 represents the address SECP256K1 protocol.
	SECP256K1
	// Actor represents the address Actor protocol.
	Actor
	// BLS represents the address BLS protocol.
	BLS

	Unknown = Protocol(255)
)

func TestStrconv(*testing.T){
	s := "101"
	v,err := strconv.ParseUint(s,10,64)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("strconv.ParseUint(s,10,64)",v)
	raw := "101"
	id, err := strconv.ParseUint(raw, 10, 64)
	payload := varint.ToUvarint(id)
	fmt.Println("payload",payload)
	fmt.Println("varint.ToUvarint(id)",varint.ToUvarint(id))
	_, n, err := varint.FromUvarint(payload)
	if err != nil {
		return 
	}
	if n != len(payload) {
		return
	}

	fmt.Println("n",n)
	explen := 1 + len(payload)
	buf := make([]byte, explen)
	fmt.Println("explan ",explen)
	var protocol Protocol
	protocol = ID

	buf[0] = protocol
	copy(buf[1:], payload)
	fmt.Printf("string(buf) %s\n",string(buf))
	fmt.Printf("%x\n",&(Address{string(buf)}))
	fmt.Println("buf ",buf)
	
}