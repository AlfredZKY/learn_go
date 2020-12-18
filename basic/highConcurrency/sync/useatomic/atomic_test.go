package useatomic

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	// 原子操作的第一个参数，是被操作的值 因为原子操作函数需要是被操作值得指针，而不是这个值本身，被传入函数的参数值都会被复制
	// 原子操作加法函数做原子减法操作 有符号类型
	num := int32(18)
	t.Logf("the num is %d\n", num)
	atomic.AddInt32(&num, int32(3))
	t.Logf("the num is %d\n", num)

	uintNum := uint32(18)
	t.Logf("the uintNum is %d\n", uintNum)
	delta := int32(-3)
	atomic.AddUint32(&uintNum, uint32(delta))
	t.Logf("the uintNum is %d\n", uintNum)

	uintNum = uint32(18)
	t.Logf("the uintNum is %d\n", uintNum)
	atomic.AddUint32(&uintNum, ^uint32(-(-3)-1))
	t.Logf("the uintNum is %d\n", uintNum)

	// -3的补码
	t.Logf("The two's complement of %d: %b\n", delta, uint32(int32(delta)))
	t.Logf("The equivalent:%b\n", ^uint32(-(-3)-1))
}

func TestForAndCAS1(t *testing.T) {
	// 原子交换自旋锁
	sign := make(chan struct{}, 2)
	num := int32(0)
	t.Logf("The number:%d\n", num)

	// 定时增加num的值
	go func() {
		defer func() {
			sign <- struct{}{}
		}()

		for {
			time.Sleep(time.Millisecond * 500)
			newNum := atomic.AddInt32(&num, 2)
			t.Logf("The number:%d\n", newNum)
			if newNum == 10 {
				break
			}
		}
	}()

	// 定时检查num的值，如果等于10就重置
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		// 不断比较检查&num地址指向的值与10比较，如果相等就和0进行交换，并返回true 否则为false
		for {
			if atomic.CompareAndSwapInt32(&num, 10, 0) {
				t.Log("The number has gone to zero.")
				break
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	
	<-sign
	<-sign
	t.Logf("The number has been swaped to :%d\n", num)
}

func TestForAndCAS2(t *testing.T) {
	// 原子交换的互斥锁,
	sign := make(chan struct{}, 2)
	num := int32(0)
	t.Logf("The number:%d\n", num)

	max := int32(20)

	// 定时增加 num的值
	go func(id int, max int32) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 0; ; i++ {
			curNum := atomic.LoadInt32(&num)
			if curNum >= max {
				break
			}
			newNum := curNum + 2
			time.Sleep(time.Millisecond * 200)
			if atomic.CompareAndSwapInt32(&num, curNum, newNum) {
				t.Logf("The number: %d [%d-%d]\n", newNum, id, i)
			} else {
				t.Logf("The CAS operation failed. [%d-%d]\n", id, i)
			}
		}
	}(1, max)

	// 定时增加num的值
	go func(id int, max int32) {
		defer func() {
			sign <- struct{}{}
		}()

		for j := 0; ; j++ {
			curNum := atomic.LoadInt32(&num)
			if curNum >= max {
				break
			}
			newNum := curNum + 2
			time.Sleep(time.Millisecond * 200)
			if atomic.CompareAndSwapInt32(&num, curNum, newNum) {
				t.Logf("The number:%d [%d-%d]\n", newNum, id, j)
			} else {
				t.Logf("The CAS operation failed. [%d-%d]\n",  id, j)
			}
		}
	}(2, max)

	<-sign
	<-sign
}
