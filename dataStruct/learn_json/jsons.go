package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	//"github.com/filecoin-project/go-state-types/abi"
	"io/ioutil"
	"log"
)

//定义一个结构体
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

type Student struct {
	Name string
	Age uint8
	Address string
}

func serializeJson() {
	monster := Monster{
		Name:     "minger",
		Age:      23,
		Birthday: "1997-11-13",
		Sal:      2000.0,
		Skill:    "Linux C/C++ Go",
	}

	data, err := json.Marshal(&monster)

	if err != nil {
		fmt.Printf("序列号错误 err = %v\n", err)
	}
	fmt.Printf("monster 序列化后= %v\n", data)

	var newData Monster

	err = json.Unmarshal(data, &newData)
	if err != nil {
		fmt.Println("enter")
		fmt.Println(newData)
	}
	fmt.Println(monster)
}



func serializeGob(){
	//序列化
	ss := make([]Student,10)
	s1:=Student{"张三",18,"江苏省"}
	ss = append(ss, s1)
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)//创建编码器
	err1 := encoder.Encode(&ss)//编码
	if err1!=nil{
		log.Panic(err1)
	}
	fmt.Printf("序列化后：%x\n",buffer.Bytes())
	if err:=ioutil.WriteFile("./piecesinfo.json",buffer.Bytes(),0644);err != nil {
		fmt.Println(errors.New("errors"))
	}
	//反序列化
	newss := make([]Student,10)
	byteEn:=buffer.Bytes()
	decoder := gob.NewDecoder(bytes.NewReader(byteEn)) //创建解密器
	//var s2 Student
	err2 := decoder.Decode(&newss)//解密
	if err2!=nil{
		log.Panic(err2)
	}

	fmt.Println("反序列化后：",newss)


}



//结构体序列化
func main() {
	serializeGob()
}
