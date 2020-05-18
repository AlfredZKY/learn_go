package readconfig

import (
	"fmt"
	"encoding/json"
	"os"
)

// ConfigFileJSON 构造存储json文件的结构体
type ConfigFileJSON struct {
	Enabled bool
	Path    string
}

// ReadConfJSON 读取json配置文件
func ReadConfJSON(path string) ConfigFileJSON{
	file,_ := os.Open(path)

	// 关闭文件
	defer file.Close()

	// 从NewDecodeer中创建一个file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据
	decoder := json.NewDecoder(file)
	conf := ConfigFileJSON{}

	// Decoder 从输出流中读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:",err)
	}
	// fmt.Println("path:"+conf.Path)
	return conf
}
