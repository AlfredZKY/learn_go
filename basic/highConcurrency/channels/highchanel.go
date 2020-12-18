package channels

import (
	"fmt"
	"time"
	"math/rand"
)


// <- 紧挨类型字面量右边的<- 表示这个通道是单向的，只能发不能收
var sendChan = make(chan<- int,1)

// <- 紧挨类型字面量左边的<- 表示这个通道是单向的，只能收不能发 
var receiveChan = make(<-chan int,1)

// 单向通道的用途主要就是约束其他代码的行为

// SendInt 例如函数 参数可以传入一个双通道，go自动会把双通道转化为函数所需的单通道
func SendInt(ch chan<- int){
	ch <- rand.Intn(1000)
}

// GetIntChan return <-chan int 
func GetIntChan()<-chan int{
	num:=5
	ch :=make(chan int,num)
	for i:=0;i<num;i++{
		ch <- i
	}
	close(ch)
	return ch
}

// UseChanSelect use select
func UseChanSelect(){
	intChannels :=[3]chan int{
		make(chan int,1),
		make(chan int,1),
		make(chan int,1),
	}

	// 随机选择一个通道，并向它发送元素值
	index := rand.Intn(3)
	fmt.Printf("The index:%d\n",index)
	intChannels[index] <- index

	// 哪一个通道中有值，对应的分支就会被执行
	select{
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The seconde candidate case is selected.")
	case <-intChannels[2]:
		fmt.Println("The third candidate case is selected.")
	default:
		fmt.Println("No candidate case is selected!")
	}
}

// DetermineChanClose select 语句中判断是否通关已关闭
func DetermineChanClose(){
	intChan := make(chan int,1)
	time.AfterFunc(time.Second, func(){
		close(intChan)
	})

	select{
	case _,ok:=<-intChan:
		if !ok{
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}
