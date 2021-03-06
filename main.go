package main

import (
	"fmt"
	"io"
	"math/big"
	"sync"

	//"learn_go/highConcurrency/resource"
	"log"
	// "math/rand"
	// "sync"
	"sync/atomic"
	// "time"
)

const (
	// 模拟最大的goroutine
	maxGoroutine = 6
	// 资源池的大小
	poolRes = 2
)

func main1() {
	// 等待任务完成
	// var wg sync.WaitGroup
	// wg.Add(maxGoroutine)

	// // p, err := resource.New(createConnection, poolRes)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// // 模拟好几个goroutine同时使用资源池查询数据
	// for query := 0; query < maxGoroutine; query++ {
	// 	go func(q int) {
	// 		// dbQuery(q, p)
	// 		wg.Done()
	// 	}(query)
	// }
	// wg.Wait()
	// log.Println("开始关闭资源池")
	// p.Close()
}

// func dbQuery(query int, pool *resource.Pool) {
// 	conn, err := pool.Acquire()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer pool.Release(conn)

// 	// 模拟查询
// 	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
// 	log.Printf("第%d个查询，使用的是ID为%d的数据库连接", query, conn.(*dbConnection).ID)
// }

type dbConnection struct {
	ID int32 // 连接标志
}

func (db *dbConnection) Close() error {
	log.Println("关闭连接", db.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	// 并发安全，给数据库连接生成唯一标志
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}, nil
}

var (
	width        = 1 << widthBits
	maxHeight    = maxIndexBits/widthBits - 1
	maxIndexBits = 63
	widthBits    = 3
	MaxIndex     = uint64(1<<maxIndexBits) - 1
)

func myadd(a, b int) bool {
	return true
}

func main() {

	log.Println(maxIndexBits, maxHeight, width, widthBits)
	const SectorsMax = 1 << 40
	fmt.Println(SectorsMax / 1024 / 1024 / 1024)
	fmt.Println(SectorsMax)
	var BaselineInitialValue = big.NewInt(2_888_888_880_000_000_000)
	fmt.Println(BaselineInitialValue)
	Eibs := big.NewInt(1024*1024*1024*1024*1024*1024)
	fmt.Println(Eibs)
	ip := big.NewInt(0)
	ip.Div(BaselineInitialValue,Eibs)
	fmt.Println("----------------------")
	fmt.Printf("%v\n",ip)
	fmt.Println("----------------------")
	log.Println(MaxIndex)

	val := "fool"
	log.Println([]byte(val))
	// var data [10]byte = {"bmlzaGFuZw=="}
	// bytess := bytes.NewBuffer()
	// var b = []byte{'b','m','l','z','a','G','F','u','Z','w','=','='}
	var b = []byte{0xf4}
	log.Println(b)
	a, c := 0, 1
	log.Println(myadd(a, c))

	//var syncMap sync.Map
	//m1 := make(map[string]sync.Map, 1000)
	//syncMap.Store("w1", 1)
	//m1["ap"] = syncMap
	//
	//tempSyncMap := m1["p1"]
	//tempSyncMap.Store("w2", 5)
	//fmt.Println(tempSyncMap)

	var readyCh chan struct{}
	if readyCh == nil {
		fmt.Println("readyCh")
	}

	var lc sync.Mutex
	fmt.Printf("%v\n", lc)
	lc.Lock()
	fmt.Printf("%v\n",lc)
	lc.Unlock()
	fmt.Printf("%v\n",lc)

	defer func(){
		if r:=recover();r!=nil{
			fmt.Printf("%v\n",lc)
		}
	}()
}
