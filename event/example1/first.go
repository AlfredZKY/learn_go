package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Wait Struct
type Wait interface {
	Register(id uint64) <-chan interface{}
	Trigger(id uint64, x interface{})
	IsRegistered(id uint64) bool
}

type list struct {
	l sync.RWMutex
	m map[uint64]chan interface{}
}

// Newlist Return Wait struct
func Newlist() Wait {
	return &list{m: make(map[uint64]chan interface{})}
}

// func Register
func (w *list) Register(id uint64) <-chan interface{} {
	w.l.Lock()
	defer w.l.Unlock()

	ch := w.m[id]
	if ch != nil {
		log.Fatal("dup id error")
		return nil
	}

	ch = make(chan interface{}, 1)
	w.m[id] = ch
	return ch
}

// func Trigger
func (w *list) Trigger(id uint64, x interface{}) {
	w.l.Lock()
	ch := w.m[id]
	delete(w.m, id)
	w.l.Unlock()

	if ch != nil {
		ch <- x
		close(ch)
	}
}

// func IsRegistered
func (w *list) IsRegistered(id uint64) bool {
	w.l.RLock()
	defer w.l.Unlock()
	_, ok := w.m[id]
	return ok
}

var timeOutDuration = time.Minute * 1

func main() {
	fmt.Println("hello world")
	list := Newlist()
	rid := uint64(time.Now().UnixNano())
	go func() {
		ch := list.Register(rid)
		fmt.Println("start register:", rid)
		if ch == nil {
			return
		}
		select {
		case x :=<- ch:
			fmt.Printf("trigger over id:%d,x:%v\n", rid, x)
		case <-time.After(timeOutDuration):
			log.Println("timeout error:", rid)
		}
	}()
	time.Sleep(time.Second)

	rid2 := uint64(time.Now().UnixNano())
	go func() {
		ch := list.Register(rid2)
		fmt.Println("start register:", rid2)
		if ch == nil {
			return
		}
		select {
		case x :=<- ch:
			fmt.Printf("trigger over id:%d,x:%v\n", rid2, x)
		case <-time.After(timeOutDuration):
			log.Println("timeout error:", rid2)
		}
	}()
	
	go func (){
		time.Sleep(time.Second*5)
		list.Trigger(rid,"hello")
		time.Sleep(time.Second*3)
		list.Trigger(rid2,"world")
	}()

	select{

	}
}
