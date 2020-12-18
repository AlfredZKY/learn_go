package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	//"testing"
	"time"
)

const (
	// 模拟的最大goroutine
	maxGoroutine = 5
)

type dbConnection struct {
	ID int32 // 连接的标志
}

func (db *dbConnection) Close() error {
	log.Println("关闭连接", db.ID)
	return nil
}

var idCouter int32

// 生成数据库的连接方法，以供资源池使用
func createConnection() interface{} {
	// 并发安全，给数据库连接生成唯一标志
	id := atomic.AddInt32(&idCouter, 1)
	return &dbConnection{ID: id}
}

func dbQuery(query int, pool *sync.Pool) {
	conn := pool.Get().(*dbConnection)
	defer pool.Put(conn)

	// 模拟查询
	log.Println(time.Duration(rand.Intn(10)) * time.Millisecond)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	log.Printf("第%d个查询，使用的是ID为%d的数据库连接", query, conn.ID)
}

func main() {
	// 等待任务完成
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)

	p := &sync.Pool{
		New: createConnection,
	}

	// 模拟好几个goroutine同时使用资源池查询数据
	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
}