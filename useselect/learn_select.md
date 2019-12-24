# go 和出中select的区别
C 语言中的 select 关键字可以同时监听多个文件描述符的可读或者可写的状态，在文件描述符发生状态改变之前，select 会一直阻塞当前的线程，Go 语言中的 select 关键字与 C 语言中的有些类似，只是它能够让一个 Goroutine 同时等待多个 Channel 达到准备状态。
select是一种与switch非常相似的控制结构，与switch不同的是，select中虽然也有多个case，但是这些case中表达式都必须与Channel的操作有关，也就是Channel的读写操作
