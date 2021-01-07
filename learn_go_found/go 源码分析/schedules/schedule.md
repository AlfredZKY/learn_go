| 情形 | 说明 |  
| ---- | --- |
| 使用关键字 go | go 创建一个新的 goroutine，Go scheduler 会考虑调度 |  
| GC | 由于进行 GC 的 goroutine 也需要在 M 上运行，因此肯定会发生调度。当然，Go scheduler 还会做很多其他的调度，例如调度不涉及堆访问的 goroutine 来运行。GC 不管栈上的内存，只会回收堆上的内存 |  
| 系统调用 | 当 goroutine 进行系统调用时，会阻塞 M，所以它会被调度走，同时一个新的 goroutine 会被调度上来 |  
| 内存同步访问 | atomic，mutex，channel 操作等会使 goroutine 阻塞，因此会被调度走。等条件满足后（例如其他 goroutine 解锁了）还会被调度上来继续运行 | 


# Guard page
为了防止发生这种情况，OS 在 Stack 和 Heap 之间设置了一段不可被 overwrite 的区域：Guard Page


# G-M 模型
M：代表线程，goroutine都是由线程来执行的；  
Global G Queue：全局goroutine队列，其中G就代表goroutine，所有M都从这个队列中取出goroutine来执行。  
这种模型比较简单，但是问题也很明显：  
- 多个M访问一个公共的全局G队列，每次都需要加互斥锁保护，造成激烈的锁竞争和阻塞；  
- 局部性很差，即如果M1上的G1创建了G2，需要将G2交给M2执行，但G1和G2是相关的，最好放在同一个M上执行。  
- M中有mcache(内存分配状态)，消耗大量内存和较差的局部性。  
- 系统调用syscall会阻塞线程，浪费不能合理的利用CPU。   


# G M P 模型
- 每个P有个局部队列，局部队列保存待执行的goroutine  
- 每个P和一个M绑定，M是真正的执行P中goroutine的实体  
- 正常情况下，M从绑定的P中的局部队列获取G来执行  
- 当M绑定的P的的局部队列已经满了之后就会把goroutine放到全局队列  
- M是复用的，不需要反复销毁和创建，拥有work stealing和hand off策略保证线程的高效利用。  
- 当M绑定的P的局部队列为空时，M会从其他P的局部队列中偷取G来执行，即work stealing；当其他P偷取不到G时，M会从全局队列获取到本地队列来执行G。  
- 当G因系统调用(syscall)阻塞时会阻塞M，此时P会和M解绑即hand off，并寻找新的idle的M，若没有idle的M就会新建一个M。  
- 当G因channel或者network I/O阻塞时，不会阻塞M，M会寻找其他runnable的G；当阻塞的G恢复后会重新进入runnable进入P队列等待执行  
- mcache(内存分配状态)位于P，所以G可以跨M调度，不再存在跨M调度局部性差的问题  
- G是抢占调度。不像操作系统按时间片调度线程那样，Go调度器没有时间片概念，G因阻塞和被抢占而暂停，并且G只能在函数调用时有可能被抢占，极端情况下如果G一直做死循环就会霸占一个P和M，Go调度器也无能为力。  


# 查看本地调度队列
```

SCHED 0ms: gomaxprocs=8 idleprocs=5 threads=5 spinningthreads=1 idlethreads=0 runqueue=0 [0 0 0 0 0 0 0 0]
---> a
---> b
---> c
c
b
a
SCHED 1008ms: gomaxprocs=8 idleprocs=0 threads=11 spinningthreads=0 idlethreads=2 runqueue=229755 [207 0 39 0 153 140 113 119]
SCHED 2016ms: gomaxprocs=8 idleprocs=0 threads=11 spinningthreads=0 idlethreads=2 runqueue=542378 [249 197 179 112 221 193 168 226]
SCHED 3026ms: gomaxprocs=8 idleprocs=0 threads=11 spinningthreads=0 idlethreads=2 runqueue=542725 [207 153 139 68 178 150 141 183]
SCHED 4029ms: gomaxprocs=8 idleprocs=0 threads=11 spinningthreads=0 idlethreads=2 runqueue=543043 [167 113 99 29 138 110 101 143]
SCHED 5041ms: gomaxprocs=8 idleprocs=0 threads=11 spinningthreads=0 idlethreads=2 runqueue=362218 [126 72 58 93 97 69 60 101]
```
- SCHED 1000ms  
    - 自程序运行开始经历的时间  
- gomaxprocs=4  
    - 当前程序使用的逻辑processor，即P，小于等于CPU的核数。  
- idleprocs=4  
    - 空闲的线程数  
- threads=8  
    - 当前程序的总线程数M，包括在执行G的和空闲的  
- spinningthreads=0  
    - 处于自旋状态的线程，即M在绑定的P的局部队列和全局队列都没有G，M没有销毁而是在四处寻觅有没有可以steal的G，这样可以减少线程的大量创建。  
- idlethreads=3  
    - 处于idle空闲状态的线程  
- runqueue=0  
    - 全局队列中G的数目
- [0 0 0 6]
    - 本地队列中的每个P的局部队列中G的数目，我的电脑是四核所有有四个P。

#     
- 在用户态就可以完成切换操作，不涉及到OS