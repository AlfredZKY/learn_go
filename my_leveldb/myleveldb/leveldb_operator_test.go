package myleveldb

import (
	"fmt"
	"testing"
)



// TestKeyValue_BatchWriteData
// 批量写入操作
func TestKeyValue_BatchWriteData(t *testing.T) {
	ky := KeyValue{Init("accountDb")}
	var i = 0
	li := make([] map[string]string,0)
	dics := make(map[string]string)
	for true {
		if i > 10 {
			break
		}
		st := "st" + string(Int2byte(i)[:])
		dics[st] = st
		i++
	}
	data := append(li,dics)
	err := ky.BatchWriteData(data)
	if err!= nil{
		fmt.Println("call BatchWriteData success")
	}
}

// TestKeyValue_WriteData
// 单一key value键值对写入操作
func TestKeyValue_WriteData(t *testing.T){
	ky := KeyValue{Init("accountDb")}
	key := "hello"
	value := "world"
	err := ky.WriteData(key,value)
	if err!=nil{}
	fmt.Println("call WriteData1 success")
}

//TestKeyValue_SearchData
// 根据key键搜索对应的值
func TestKeyValue_SearchData(t *testing.T) {
	ky := KeyValue{Init("accountDb")}
	key := "hello"
	err := ky.SearchData(key)
	if err != nil{
		fmt.Println("call SearchData success ")
	}
}

// TestKeyValue_UpdateData
// 根据key键更新对应的值
func TestKeyValue_UpdateData(t *testing.T) {
	ky := KeyValue{Init("accountDb")}
	key := "hello"
	value := "haha"
	err := ky.UpdateData(key,value)
	if err != nil{
		fmt.Println("call SearchData success ")
	}
}

//TestKeyValue_DeleteData
// 根据key键删除对应的值
func TestKeyValue_DeleteData(t *testing.T) {
	ky := KeyValue{Init("accountDb")}
	key := "hello"
	err := ky.DeleteData(key)
	if err != nil{
		fmt.Println("call DeleteData success ")
	}
	RecursiveDeleteDic(ky.DbPath)
}
