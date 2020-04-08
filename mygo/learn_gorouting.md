# 什么是goroutine
Goroutine是建立在线程之上的轻量级的抽象，它允许我们以非常低的代价在用一个地址空间中并行地执行多个函数或者方法，相比于线程，它的创建和销毁的代价要小很多，并且它的调度是独立于线程的
和线程一样，golang的主函数其实也是跑在一个goroutine中，它并不会等待其它goroutine结束，如果主goroutine结束了，所有其它goroutine都将结束。

# goroutine和线程的区别
goroutine并不会比线程运行的更快，它只是增加了更多的并发性，当一个goroutine被阻塞时(比如IO等待)golang的schedule会调度其它可以执行的goroutine运行，所以相比线程，有以下优点:
- 内存消耗更少
  - goroutine所需要的内存通常只有2kb,而线程则需要1MB
- 创建与销毁的开销更小
  - 线程的创建需要向操作系统申请资源，并且销毁时将资源归还，因此它的创建和销毁的开销比较大，相比之下，goroutine的创建和销毁是由go语言在运行时自己管理的，因此开销更低
- 切换开销更小
  - 这是goroutine与线程的主要区别，也是golang能够实现高并发的主要原因，线程的调度方式是抢占式的，如果一个线程的执行时间超过了分配给它的时间片，就会被其他可执行的线程抢占,在线程切换的过程中需要保存/恢复所有的寄存器的信息，比如16个通用寄存器PC(Program Counter) SP(Stack Pointer),段寄存器等等。

# goroutine的调度
goroutine的调度方式是协同的，在协同式调度中，没有时间片的概念，为了并执行goroutine，调度器会在以下几个时间点对其进行切换:
- channel接受或者发送会造成阻塞的消息
- 当一个新的goroutine被创建时
- 可以造成阻塞的系统调用，如文件和网络操作
- 垃圾回收

# golang调度器中的三个感念
- Processor(P)
- OSThread(M)
- Goroutine(G)
在一个Go程序中，可用的线程数是通过GOMAXPROCS来设置的，默认值是可用的CPU核数。我们可以通过用runtime包动态改变这个值。OsThread调度在processor上，goroutines调度在OsThread上    

# 如何让主goroutine等待其他goroutine
- 通过睡眠函数，阻塞一下主goroutine(测试时可以用，一般不用，无法预估出时间)
- 通过通道来告诉主goroutine子goroutine结束了，缓冲通道容量来阻塞等待
- sync.WaitGroup类型(未学)

# 怎样让启用的goroutine按照既定顺序执行
