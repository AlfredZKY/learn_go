package main

import (
	"fmt"
	"io/ioutil"
	pb "learn_go/myprobuf/firsttest/test"
	"os"

	// "github.com/golang/protobuf/proto"
	"encoding/json"

	"google.golang.org/protobuf/proto"
)

func write() {
	p1 := &pb.Person{
		Id:     1,
		Name:   "小张",
		Phones: []*pb.Person_Phone{{Type: pb.Person_HOME, Number: "111111"}, {Type: pb.Person_WORK, Number: "222222"}},
	}

	p2 := &pb.Person{
		Id:     2,
		Name:   "小王",
		Phones: []*pb.Person_Phone{{Type: pb.Person_HOME, Number: "333333"}, {Type: pb.Person_WORK, Number: "444444"}},
	}

	// 创建地址薄
	book := &pb.ContackBook{}
	book.People = append(book.People, p1)
	book.People = append(book.People, p2)

	// 编码数据
	data, _ := proto.Marshal(book)
	fmt.Printf("proto encode data of len is %v\t%v\n", len(data), data)

	// 使用json对比
	jsData, er := json.Marshal(book)
	if er == nil {
		fmt.Printf("js encode data of len is %v\t%v\n", len(jsData), jsData)
	}

	// 把数据写入文件
	ioutil.WriteFile("./test.txt", data, os.ModePerm)
}

func read() {
	// 读取数据
	data, _ := ioutil.ReadFile("./test.txt")
	book := &pb.ContackBook{}

	// 解码数据
	proto.Unmarshal(data, book)
	for _, v := range book.People {
		fmt.Println(v.Id, v.Name)
		for _, vv := range v.Phones {
			fmt.Println(vv.Type, vv.Number)
		}
	}
}

func main() {
	fmt.Println("hello world")
	write()
	read()
}
