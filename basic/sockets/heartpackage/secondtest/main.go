package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"compress/gzip"
	pb "learn_go/sockets/heartpackage/secondtest/protocol"

	"google.golang.org/protobuf/proto"
)

func gzipCompress(content *[]byte) []byte {
	var compressData bytes.Buffer
	gzipWriter := gzip.NewWriter(&compressData)
	defer gzipWriter.Close()
	gzipWriter.Write(*content)
	gzipWriter.Flush()
	return compressData.Bytes()
}

func gzipUnCompress(content *[]byte) []byte {
	var uncompressData bytes.Buffer
	uncompressData.Write(*content)
	r, _ := gzip.NewReader(&uncompressData)
	defer r.Close()
	undatas, _ := ioutil.ReadAll(r)
	return undatas
}

func Init() *pb.MainTable {
	sub1 := &pb.MainInfo{
		Weight: 90,
		Status: false,
		Ip:     "192.168.10.20",
		Port:   1234,
	}
	sub2 := &pb.MainInfo{
		Weight: 80,
		Status: false,
		Ip:     "192.168.10.21",
		Port:   2345,
	}

	maintable := &pb.MainTable{}
	maintable.Maintable = append(maintable.Maintable, sub1)
	maintable.Maintable = append(maintable.Maintable, sub2)
	return maintable
}

func main() {
	maintable := Init()
	// 对结果数据进行编码
	data, err := proto.Marshal(maintable)
	if err == nil {
		fmt.Printf("proto encode data of len is %v\t%v\n", len(data), data)
	}
	// 对序列化的数据进行压缩
	gzipdata := gzipCompress(&data)
	fmt.Printf("gzip data of len is %v\t%v\n", len(gzipdata), gzipdata)

	// 对序列化的数据进行解压缩
	ungzipdata := gzipUnCompress(&gzipdata)
	fmt.Printf("proto encode data of len is %v\t%v\n", len(ungzipdata), ungzipdata)

	maintableres1 := &pb.MainTable{}
	proto.Unmarshal(ungzipdata, maintableres1)

	for _, v := range maintableres1.Maintable {
		str := strconv.Itoa(int(v.Port))
		fmt.Println(v.Weight, v.Ip, v.Port, v.Ip+":"+str,len(v.Ip+":"+str))
	}
	// 把数据写入到文件
	ioutil.WriteFile("./encode_Data", data, os.ModePerm)

	// 读取文件中的数据
	maintableres := &pb.MainTable{}
	data, _ = ioutil.ReadFile("./encode_Data")
	// 解码出数据
	proto.Unmarshal(data, maintableres)

	for _, v := range maintableres.Maintable {
		fmt.Println(v.Weight, v.Ip, v.Port)
	}
}
