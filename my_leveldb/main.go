package main

import (
	"fmt"
	"io/ioutil"
	"learn_go/my_leveldb/layerswallet"
	"learn_go/my_leveldb/myleveldb"
	"os"
	"runtime"
	"strings"
	"time"
)

func isWinLin() string{
	sysType := runtime.GOOS

	if sysType == "linux" {
		res := "linux"
		return res
	}

	if sysType == "windows" {
		res := "windows"
		return res
	}
	res := "Unkown"
	return res
}

// RemoveAllFile 读取所有文件
func RemoveAllFile(pathname string,flag string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("%s\n", pathname + "\\"+fi.Name())
			if flag == "windows"{
				RemoveAllFile(pathname + "\\" + fi.Name(),flag)
				err:=os.Remove(pathname + "\\" + fi.Name())
				if err!=nil{
					fmt.Println(err)
				}
				fmt.Printf("removeall %s success\n",pathname + "\\" + fi.Name())
			} else if flag == "linux"{
				RemoveAllFile(pathname + "/" + fi.Name(),flag)
				err:=os.Remove(pathname + "/" + fi.Name())
				if err!=nil{
					fmt.Println(err)
				}
				fmt.Printf("removeall %s success\n",pathname + "/" + fi.Name())
			}
		} else {
			if flag == "windows"{
				filePath := pathname + "\\" + fi.Name()
				err := os.Remove(filePath)
				if err != nil{
					fmt.Printf("%s is delete success",filePath)
				}
			}else if flag == "linux"{
				filePath := pathname + "/" + fi.Name()
				err := os.Remove(filePath)
				if err != nil{
					fmt.Printf("%s is delete success",filePath)
				}
			}
		}
	}
	return err
}

// RecursiveDeleteDic 递归删除所有文件
func RecursiveDeleteDic(path string){
	res := isWinLin()
	// res := "linux"
	if res == "windows"{
		RemoveAllFile(path,res)
		err:= os.Remove(path)
		if err!= nil{
			fmt.Println(err)
		}else{
			fmt.Printf("removeall %s success\n",path)
		}
	}else if res == "linux"{
		RemoveAllFile(path,res)
		err:= os.Remove(path)
		if err!= nil{
			fmt.Println(err)
		}else{
			fmt.Printf("removeall %s success\n",path)
		}
	}else{
		fmt.Println(res)
	}
}

func get_args(dbpath string,ky myleveldb.KeyValue){
	fmt.Println("请输入要查找的key,输入exit(Exit)退出:")
	for true{
		str := ""
		fmt.Scanf("%s", &str)
		newKey := strings.Trim(str,"\n")
		if str == ""{
			continue
		}
		//str = strings.Replace(str, "\n", "", -1)
		//str = strings.Replace(str, "\r", "", -1)
		if str == "exit" || str == "Exit"{
			break
		}
		ky.SearchData(newKey)
		//fmt.Printf("输入了：%s\n", str)
	}
}

func main() {
	if 2>1 {
		key := "hello"
		value := "world"
		val := "newer"
		layerswallet.WriteWalletDB(key, value)
		layerswallet.SearchWalletDB(key)
		layerswallet.UpdateWalletDb(key, val)
		layerswallet.DeleteWalletDb(key)
	}
	myDir := "F:/gopath/src/my_leveldb/src/test/arc/te"
	err := os.MkdirAll(myDir,os.ModePerm)
	if err != nil{
		fmt.Println(err)
	}

	time.Sleep(time.Second * 3)

	// 删除制定路径下的所有目录和问价
	deDir := "F:/gopath/src/my_leveldb/src"
	err = os.RemoveAll(deDir)
	if err != nil{
		fmt.Println(err)
	}
}
