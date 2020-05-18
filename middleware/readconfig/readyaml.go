package readconfig

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// ConfigFileYaml 读取yaml格式的配置文件
type ConfigFileYaml struct {
	Enabled bool   `yaml:"enabled"` // yaml：yaml格式 Enabled：属性
	Path    string `yaml:"path"`
}

// ReadConfYaml 读取yaml格式的配置文件
func (conf *ConfigFileYaml) ReadConfYaml(path string) *ConfigFileYaml {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("YamlFile.Get err #%v", err)
	}
	err = yaml.Unmarshal(yamlFile,conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v",err)
	}
	return conf
}
