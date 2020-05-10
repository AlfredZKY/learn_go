package opp

import (
	"bytes"
	"encoding/json"
	"math/rand"
)

// Box 代表数据盒子
type Box struct {
	Str   string
	Code  rune
	Bytes []byte
}

// MemProFile 内存信息
func MemProFile() error {
	max := 50000
	var buf bytes.Buffer
	for j := 0; j < max; j++ {
		seed := rand.Intn(95) + 32
		one := CreateBox(seed)
		b, err := GenJSON(one)
		if err != nil {
			return err
		}
		buf.Write(b)
		buf.WriteByte('\t')
	}
	_ = buf.String()
	return nil
}

// CreateBox 创建数据
func CreateBox(seed int) Box{
	if seed <= 0 {
		seed = 1
	}
	var array []byte
	size := seed * 8
	for i := 0; i < size; i++ {
		array = append(array, byte(seed))
	}
	return Box{
		Str:   string(seed),
		Code:  rune(seed),
		Bytes: array,
	}
}

// GenJSON 序列化为JSON文件
func GenJSON(one Box) ([]byte, error) {
	return json.Marshal(one)
}
