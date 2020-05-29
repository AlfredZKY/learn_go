package common

import "sync"

// SynchronizedMap 安全的map
type SynchronizedMap struct {
	rw   *sync.RWMutex
	data map[interface{}]interface{}
}

// NewSynchronizedMap 申请一个实例对象
func NewSynchronizedMap() *SynchronizedMap {
	return &SynchronizedMap{
		rw:   new(sync.RWMutex),
		data: make(map[interface{}]interface{}),
	}
}

// Put 存储操作
func (sm *SynchronizedMap) Put(k, v interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	sm.data[k] = v
}

// Delete 删除操作
func (sm *SynchronizedMap) Delete(k interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	delete(sm.data, k)
}

// Get 获取操作
func (sm *SynchronizedMap) Get(k interface{}) interface{} {
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	return sm.data[k]
}

// Each 遍历操作 传入一个回掉函数，对数据进行处理
func (sm *SynchronizedMap) Each(cb func(interface{}, interface{})) {
	sm.rw.RLock()
	defer sm.rw.RUnlock()

	for k, v := range sm.data {
		cb(k, v)
	}
}
