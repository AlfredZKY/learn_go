package typejudgment

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
)

const pathSize = 16 << 20
const MetaFile = "sectorstore.json" // 元文件
type ID string

// [path]/sectorstore.json
type LocalStorageMeta struct {
	ID     ID     //id编号 "2da1072c-1452-4aa3-a8e4-f99c29077510" 不是扇区
	Weight uint64 // 0 = readonly

	CanSeal  bool
	CanStore bool
}
type FsStat struct {
	Capacity  uint64
	Available uint64 // Available to use for sector storage
	Used      uint64
}

// .lotusstorage/storage.json
type StorageConfig struct {
	StoragePaths []LocalPath
}

type LocalPath struct {
	Path string
}

type LocalStorage interface {
	GetStorage() (StorageConfig, error)
	SetStorage(func(*StorageConfig)) error

	Stat(path string) (FsStat, error)
}

type TestingLocalStorage struct {
	root string
	c    StorageConfig
}

func (t *TestingLocalStorage) GetStorage() (StorageConfig, error) {
	return t.c, nil
}

func (t *TestingLocalStorage) SetStorage(f func(*StorageConfig)) error {
	f(&t.c)
	return nil
}

func (t *TestingLocalStorage) Stat(path string) (FsStat, error) {
	return FsStat{
		Capacity:  pathSize,
		Available: pathSize,
		Used:      0,
	}, nil
}

func (t *TestingLocalStorage) init(subpath string) error {
	path := filepath.Join(t.root, subpath)
	if err := os.Mkdir(path, 0755); err != nil {
		return err
	}

	metaFile := filepath.Join(path, MetaFile)

	meta := &LocalStorageMeta{
		ID:       ID(uuid.New().String()),
		Weight:   1,
		CanSeal:  true,
		CanStore: true,
	}

	mb, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(metaFile, mb, 0644); err != nil {
		return err
	}

	return nil
}

// TestingLocalStorage结构体实现了LocalStorage 的所有接口 编译时会进行类型检查，如果没有实现会直接报错，注释掉上面的Stat函数报错如下
// cannot use &TestingLocalStorage literal (type *TestingLocalStorage) as type LocalStorage in assignment:
// *TestingLocalStorage does not implement LocalStorage (missing Stat method)
var _ LocalStorage = &TestingLocalStorage{}

func TestTpyeStructInterface(t *testing.T) {
	t.Log("test result info")
	t.Log()
}
