package common

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// ErrTimeOut 超时错误
var ErrTimeOut = errors.New("执行者执行超时")

// ErrInterrupt 终端错误
var ErrInterrupt = errors.New("执行者被中断")

// 一个执行者，可以执行任何任务，但是这些任务是限制完成的
// 该执行者可以通过发送终止信号终止它

// Runner 可以执行任务的结构体
type Runner struct {
	tasks     []func(int)      // 要执行的任务
	complete  chan error       // 用于通知任务全部完成
	timeout   <-chan time.Time // 这些任务在多久内完成
	interrupt chan os.Signal   // 可以强制终止的信号
}

// New 构建一个Runner 实例
func New(tm time.Duration) *Runner {
	return &Runner{
		complete:  make(chan error),
		timeout:   time.After(tm),
		interrupt: make(chan os.Signal, 1),
	}
}

// Add 添加一个任务到runner实例中
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrupt
		}
		// 执行任务
		task(id)
	}
	return nil
}

func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

// Start 开始执行所有任务，并且监视通道事件
func (r *Runner) Start() error {
	// 希望接受那些系统信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}
