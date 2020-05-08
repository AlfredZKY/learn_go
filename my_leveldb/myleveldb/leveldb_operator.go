package myleveldb

import (
	"bytes"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"io/ioutil"
	"log"
	"learn_go/my_leveldb/conf"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type KeyValueStore interface {
	WriteData(key string,value string)error
	UpdateData(key string, value string) error
	SearchData(key string) error
	DeleteData(key string) error
	BatchWriteData(vars keyvaluedata)error
}
type KeyValue struct {
	DbPath string
}
type keyvaluedata []map[string]string


// 执行删除所有的测试文件和目录
func removeallfile(pathname string,flag string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("%s\n", pathname + "\\"+fi.Name())
			if flag == "windows"{
				removeallfile(pathname + "\\" + fi.Name(),flag)
				err:=os.Remove(pathname + "\\" + fi.Name())
				if err!=nil{
					fmt.Println(err)
				}
				fmt.Printf("removeall %s success\n",pathname + "\\" + fi.Name())
			} else if flag == "linux"{
				removeallfile(pathname + "/" + fi.Name(),flag)
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
				fmt.Printf("removeall %s success\n",pathname + "\\" + fi.Name())
			}else if flag == "linux"{
				filePath := pathname + "/" + fi.Name()
				err := os.Remove(filePath)
				if err != nil{
					fmt.Printf("%s is delete success",filePath)
				}
				fmt.Printf("removeall %s success\n",pathname + "/" + fi.Name())
			}
		}
	}
	return err
}

// 删除所有的测试文件和目录
func RecursiveDeleteDic(path string){
	res := isWinLin()
	// res := "linux"
	if res == "windows"{
		removeallfile(path,res)
		err:= os.Remove(path)
		if err!= nil{
			fmt.Println(err)
		}else{
			fmt.Printf("removeall %s success\n",path)
		}
	}else if res == "linux"{
		removeallfile(path,res)
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

/*
字节转换成int
*/
func Byte2int(val []byte) int {
	var result int
	result, _ = strconv.Atoi(string(val))
	return result
}

/*
int转换成byte
*/
func Int2byte(val int) []byte {
	result := []byte(strconv.Itoa(val))
	return result
}

/*
返回当前目录的路径
*/
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

/*
判断目录是否存在：存在，返回true，否则返回false
*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/*
判断文件是否存在：存在，返回true，否则返回false
*/
func IsFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println(info)
		return false
	}
	fmt.Println("exists", info.Name(), info.Size(), info.ModTime())
	return true
}

/*
读取配置文件
*/
func readConfig() bool {
	res := conf.CCFrom().Testnet
	if res{
		return res
	}
	return false
}

//判断代码的运行平台
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

// 设置数据的存储目录 #TODO
func getDbStorePath(curPath string,dicName string,res bool)string{
	levelDbDirName := ""
	plat := isWinLin()
	if res  {
		if plat == "windows"{
			levelDbDirName = curPath + "\\"  + dicName
		}else if plat == "linux"{
			levelDbDirName = curPath + "/"  + dicName
		}
	} else {
		if plat == "windows"{
			levelDbDirName = curPath + "\\" + "mainnet" + "\\" + dicName
		}else if plat == "linux"{
			levelDbDirName = curPath + "/" + "mainnet" + "/" + dicName
		}
	}

	exist, err := PathExists(levelDbDirName)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
	}

	if exist {
		fmt.Printf("has dir![%v]\n", levelDbDirName)
	} else {
		fmt.Printf("no dir![%v]\n", levelDbDirName)
		// 创建文件夹
		err := os.Mkdir(levelDbDirName, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
	return levelDbDirName
}

//Init
/*
初始化
*/
func Init(args string) string {
	levelDbDirName := ""
	//path := GetCurrentDirectory()
	curPath,tempErr := os.Getwd()
	if tempErr!= nil{
		fmt.Println(tempErr)
	}
	res := readConfig()
	switch args {
	case "walletDb":
		levelDbDirName = getDbStorePath(curPath,args,res)
	case "accountDb":
		levelDbDirName = getDbStorePath(curPath,args,res)
	}
	return levelDbDirName
}

/*
从leveldb读出数据
*/
func (m *KeyValue)SearchData(key string)error {

	db, err := leveldb.OpenFile(m.DbPath, nil)
	if err != nil {
		fmt.Println(err)
	}
	data, err := db.Get([]byte(key), nil)
	defer db.Close()
	if err != nil {
		fmt.Printf("this %s is not exist and err is %s\n", key,err)
		return err
	}
	res := fmt.Sprintf("search account is %s and data is %s",key,data)
	fmt.Println(string(res))
	return nil
}

/*
向leveldb删除数据
*/
func (m *KeyValue)DeleteData(key string)error {
	db, err := leveldb.OpenFile(m.DbPath, nil)
	if err != nil {
		fmt.Println(err)
	}

	res := db.Delete([]byte(key), nil)
	if res != nil {
		return res
	}
	result := fmt.Sprintf("delete  account %s is successful",key,)
	fmt.Println(string(result))
	defer db.Close()
	return nil

}

/*
向leveldb更新数据
*/
func (m *KeyValue)UpdateData(key string, value string) error {
	db, err := leveldb.OpenFile(m.DbPath, nil)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	res1 := db.Delete([]byte(key),nil)
	if res1 != nil {
		return res1
	}
	res2 := db.Put([]byte(key), []byte(value), nil)
	if res2 != nil {
		return res2
	}
	fmt.Printf("account %s update success\n", key)
	return nil
}

/*
向leveldb写入数据
*/
func (m *KeyValue)WriteData(key string,value string)error{
	db, err := leveldb.OpenFile(m.DbPath, nil)
	if err != nil {
		fmt.Println(err)
	}
	k1 := []byte(key)
	v1 := []byte(value)
	err1 := db.Put(k1,v1,nil)
	if err1 != nil{
		return err1
	}
	res := fmt.Sprintf("account is %s and data is %s write success",key,value)
	fmt.Println(string(res))
	defer db.Close()
	return nil
}

/*
向leveldb批量写入数据
*/
func (m *KeyValue)BatchWriteData(vars keyvaluedata)error{
	db, err := leveldb.OpenFile(m.DbPath, nil)
	for a:=0 ;a < len(vars); a++{
		for key,value := range vars[a]{
			if err != nil {
				fmt.Println(err)
			}
			k1 := []byte(key)
			v1 := []byte(value)
			err1 := db.Put(k1,v1,nil)
			if err1 != nil{
				errMessage := fmt.Sprintf("account %s is write err and err is %s",key,err1)
				sendMsgDing(errMessage)
				return err1
			}
		}
	}
	defer db.Close()
	return nil
}

func sendMsgDing(txt string) {

	sendMessage := fmt.Sprintf(`{"msgtype": "text", "text": {"content": "%s" }}`, txt)
	var jsonStr = []byte(sendMessage)
	client := &http.Client{}

	post := "POST"
	url := "https://oapi.dingtalk.com/robot/send?access_token=a1ffd4047fc895698e49f97d0ddc22d50b439c6c1cafe3ef6475b9ceab5da982"
	req, err1 := http.NewRequest(post,url , bytes.NewBuffer(jsonStr))
	if err1 != nil {
		fmt.Println("Can't add http request.", err1)
		os.Exit(1)
	}
	req.Header.Add("content-type", "application/json")

	resp, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println("Can't get http response.", err2)
		os.Exit(1)
	}
	defer resp.Body.Close()
}