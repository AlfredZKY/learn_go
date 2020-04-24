package channels

import (
	"fmt"
)


// BlockChanel easy use channel block
func BlockChanel(){
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

// UseChanel easy use channel
func UseChanel(){
	ch1 := make(chan int,3)
	ch1 <- 2  // 该通道接受一个元素值
	ch1 <- 1
	ch1 <- 3

	elem1:= <-ch1 // <-通道接受表达式 通过通道流出一个值赋值给一个变量
	fmt.Printf("The first element received from channel ch1:%v\n",elem1)
}

// UseChannelPanic Test channel panic
func UseChannelPanic(){
	ch1 := make(chan int,2)
	
	// 发送方
	go func(){
		for i:=0;i<10;i++{
			fmt.Printf("Sender:sending element:%v\n",i)
			ch1 <- i
		}
		fmt.Println("Sender:close the channel...")
		close(ch1)
	}()

	// 接收方
	for {
		element,ok := <-ch1
		if !ok{
			fmt.Printf("Receiver:closed channel")
			break
		}
		fmt.Printf("Receiver:received an element:%v\n",element)
	}
	fmt.Println("End.")
}

// UseChannelArray Test channel test Array
func UseChannelArray(){
	ch := make(chan []int,1)
	s1 := []int{1,2,3}
	ch <- s1
	s2 := <- ch

	s2[0] = 100
	fmt.Println(s1,s2)

	ch2 := make(chan [3]int,1)
	s3 := [3]int{1,2,3}
	ch2 <- s3
	s4 := <- ch2
	s3[0] = 100
	fmt.Println(s3,s4)
}


func sum(s []int,c chan int){
	sum := 0
	for _,v := range s {
		sum += v
	}
	c <- sum		//send sum to c
}

// GetSumArray get sum from array
func GetSumArray(){
	s := []int{1,4,5,-2,-7,9,12}
	c := make(chan int)
	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:],c)
	x,y := <-c,<-c
	fmt.Println(x,y,x+y)
 }
