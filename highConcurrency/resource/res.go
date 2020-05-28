package resource

import (
	"errors"
	"io"
	"log"
	"sync"
)

// ErrPoolClosed 和其他错误进行有效的区分
var ErrPoolClosed = errors.New("资源池已经关闭")

// Pool 一个安全的资源池，被管理的资源都必须实现io.Closer接口
type Pool struct {
	m       sync.Mutex                // 互斥锁，保护共享资源的安全
	res     chan io.Closer            // 一个有缓冲通道，类型io.Closer  初始化确定缓冲大小
	factory func() (io.Closer, error) // 函数类型，生成新资源的，类型由使用者决定
	closed  bool                      // 判断资源池是否被关闭，如果关闭的话，在访问就会报错
}

// New 实例化资源池对象
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size的值太小了")
	}
	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil
}

// Acquire 从资源池中申请一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquire:共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:新生成资源")
		return p.factory()
	}
}

// Close 关闭资源，释放资源 Close和Release是互斥的都会对closed标志修改
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	// 关闭通道，禁止写入
	close(p.res)

	// 关闭通道的资源
	for r := range p.res {
		r.Close()
	}
}

// Release 释放资源
func (p *Pool) Release(r io.Closer) {
	// 保证该操作和Close方法的操作是安全的
	p.m.Lock()
	defer p.m.Unlock()

	// 资源池都关闭了，就剩这一个没有释放的资源，释放即可
	if p.closed {
		r.Close()
		return
	}

	select {
	case p.res <- r:
		log.Println("资源释放到池子里了")
	default:
		log.Println("资源池满了,释放这个资源吧")
		r.Close()
	}
}
