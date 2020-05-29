package main

import (
	"io"
	"learn_go/highConcurrency/resource"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	// 模拟最大的goroutine
	maxGoroutine = 6
	// 资源池的大小
	poolRes = 2
)

func main() {
	// 等待任务完成
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)

	p, err := resource.New(createConnection, poolRes)
	if err != nil {
		log.Println(err)
		return
	}

	// 模拟好几个goroutine同时使用资源池查询数据
	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
	log.Println("开始关闭资源池")
	p.Close()
}

func dbQuery(query int, pool *resource.Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Release(conn)

	// 模拟查询
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("第%d个查询，使用的是ID为%d的数据库连接", query, conn.(*dbConnection).ID)
}

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
