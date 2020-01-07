package v1

import (
	"errors"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

// Data 代表数据的类型
type Data []byte

// DataFile 代表数据文件的接口类型
type DataFile interface {
	Read() (rsn int64, d Data, err error) // Read 读取一个数据块
	Write(d Data) (wsn int64, err error)  // Write 写入一个数据块
	RSN() int64                           // RSN 会获取最后读取的数据块的序列号
	WSN() int64                           // WSN 会获取最后写入的数据块的序列号
	DataLen() uint32                      // 会获取数据块的长度
	Close() error                         // 关闭数据文件
}

// myDataFile 代表数据文件的实现类型
type myDataFile struct {
	f       *os.File     // 文件
	fmutex  sync.RWMutex // 被用于文件的读写锁
	woffset int64        // 写操作需要用到的偏移量
	roffset int64        // 读操作需要用到的偏移量
	wmutex  sync.Mutex   // 写操作需要用到的互斥锁
	rmutex  sync.Mutex   // 读操作需要用到的互斥锁
	dataLen uint32       // 数据块的长度
}

// NewDataFile 会新建一个数据文件的实例
func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	if dataLen == 0 {
		return nil, errors.New("Invalid data length")
	}

	df := &myDataFile{f: f, dataLen: dataLen}
	return df, nil
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新读偏移量
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	// 读取一个块
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	for {
		df.fmutex.RLock()
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.fmutex.RUnlock()
				continue
			}
			df.fmutex.RUnlock()
			return
		}
		d = bytes
		df.fmutex.RUnlock()
		return
	}
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	// 读取并更新写偏移量
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.woffset)
	df.wmutex.Unlock()

	// 写入一个数据块
	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}

	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	return
}

func (df *myDataFile) RSN() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) WSN() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func (df *myDataFile) Close() error {
	if df.f == nil {
		return nil
	}
	return df.f.Close()
}

func removeFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		return nil
	}
	file.Close()
	return os.Remove(path)
}

func testNew(path string, dataLen uint32, t *testing.T) {
	t.Logf("New a data file (path:%s, dataLen:%d)...\n", path, dataLen)
	dataFile, err := NewDataFile(path, dataLen)

	if err != nil {
		t.Logf("couldn't new a datafile :%s", err)
		t.FailNow()
	}

	if dataFile == nil {
		t.Log("Unnornal data file!")
		t.FailNow()
	}

	defer dataFile.Close()

	if dataFile.DataLen() != dataLen {
		t.Fatalf("Incorrent data length")
	}
}

func testRW(path string, dataLen uint32, max int, t *testing.T) {
	t.Logf("New a data file (path: %s,dataLen: %d)...\n", path, dataLen)
	dataFile, err := NewDataFile(path, dataLen)
	if err != nil {
		t.Logf("Couldn't new a data file: %s", err)
		t.FailNow()
	}
	defer dataFile.Close()

	var wg sync.WaitGroup
	wg.Add(5)

	// 写入
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			var prevWSN int64 = -1
			for j := 0; j < max; j++ {
				data := Data{
					byte(rand.Int31n(256)),
					byte(rand.Int31n(256)),
					byte(rand.Int31n(256)),
				}

				wsn,err := dataFile.Write(data)
				if err != nil{
					t.Fatalf("Unexpected writing error: %s\n",err)
				}

				if prevWSN >= 0 && wsn <= prevWSN {
					t.Fatalf("Incorect WSN %d! (lt %d)\n",wsn,prevWSN)
				}
				prevWSN = wsn
			}
		}()
	}

	// 读取
	for i:=0;i<2;i++{
		go func(){
			defer wg.Done()
			var prevRSN int64 = -1
			for i:=0;i<max;i++{
				rsn,date,err := dataFile.Read()
				if err != nil{
					t.Fatalf("Unexpected writing error: %s\n",err)
				}

				if date == nil{
					t.Fatalf("Unnormal data!")
				}

				if prevRSN >= 0 && rsn <= prevRSN{
					t.Fatalf("Incorect RSN %d! (lt %d)\n",rsn,prevRSN)
				}
				prevRSN = rsn
			}
		}()
	}
	wg.Wait()
}

func TestIDataFile(t *testing.T) {
	t.Run("v1/all", func(t *testing.T) {
		dataLen := uint32(3)
		path1 := filepath.Join(os.TempDir(), "data_file_test_new.txt")
		//t.Log(path1)
		defer func() {
			if err := removeFile(path1); err != nil {
				t.Errorf("Open file error:%s\n", err)
			}
		}()
		t.Run("New", func(t *testing.T) {
			testNew(path1, dataLen, t)
		})

		path2 := filepath.Join(os.TempDir(), "data_file_test.txt")
		defer func() {
			if err := removeFile(path2); err != nil {
				t.Fatalf("Open file error:%s\n", err)
			}
		}()
		max := 100
		t.Run("WriteAndRead", func(t *testing.T) {
			testRW(path2, dataLen, max, t)
		})

	})
}
