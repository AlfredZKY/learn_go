package readconfig

// go get gopkg.in/gcfg.v1
import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

// ConfigFileIni ini格式配置文件序列化对象
type ConfigFileIni struct {
	Section
}

// Section ini格式配置文件中的节名
type Section struct {
	Enabled bool
	Path    string
}

// ReadConfIni 读取格式为ini的配置文件
func ReadConfIni(path string) ConfigFileIni {
	config := ConfigFileIni{}
	err := gcfg.ReadFileInto(&config, path)
	if err != nil {
		fmt.Printf("Failed to parse config file: %s", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
	return config
}
