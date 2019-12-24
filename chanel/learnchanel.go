package chanel

import (
	"fmt"
	"time"
	"math/rand"
)

// func UseChanel
func UseChanel(){
	ch1 := make(chan int,3)
	ch1 <- 2  // 该通道接受一个元素值
	ch1 <- 1
	ch1 <- 3

	elem1:= <-ch1 // <-通道接受表达式 通过通道流出一个值赋值给一个变量
	fmt.Printf("The first element received from channel ch1:%v\n",elem1)
}

// func BlockChanel
func SyncChanel(){
	ch1 := make(chan int,1)
	//依次敲入通道变量的名称（比如ch1）、接送操作符<-以及想要发送的元素值（比如2），并且这三者之间最好用空格进行分割。
	ch1 <- 1 // 接受通道表达式
	//ch1 <- 2

	ch2 := make(chan int,1)
	ch2 <- 2

	// 示例3。
	var ch3 chan int
	//ch3 <- 1 // 通道的值为nil，因此这里会造成永久的阻塞！
	//<-ch3 // 通道的值为nil，因此这里会造成永久的阻塞！
	_ = ch3

}

// func CloseChanel 双向通道
func CloseChanel(){
	ch1 := make(chan int,2)

	// 发送方
	go func (){
		for i:=0;i<10;i++{
			fmt.Printf("Sender:sending element %v...\n", i)
			ch1 <- i
			time.Sleep(1)
		}
		fmt.Println("Sender:close the chanel...")
		close(ch1)
	}()

	// 接收方
	for {
		elem,ok:=<-ch1
		if !ok{
			fmt.Println("Receiver:close channel")
			break
		}
		fmt.Printf("Receiver:received an element:%v\n", elem)
	}
	fmt.Println("END.")
}


// <- 紧挨类型字面量右边的<- 表示这个通道是单向的，只能发不能收
var sendChan = make(chan<- int,1)

// <- 紧挨类型字面量左边的<- 表示这个通道是单向的，只能收不能发 
var receiveChan = make(<-chan int,1)

// 单向通道的用途主要就是约束其他代码的行为
// func SendInt 例如函数 参数可以传入一个双通道，go自动会把双通道转化为函数所需的单通道
func SendInt(ch chan<- int){
	ch <- rand.Intn(1000)
}

// func GetIntChan
func GetIntChan()<-chan int{
	num:=5
	ch :=make(chan int,num)
	for i:=0;i<num;i++{
		ch <- i
	}
	close(ch)
	return ch
}

// func UseChanSelect
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

// func DetermineChanClose select 语句中判断是否通关已关闭
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