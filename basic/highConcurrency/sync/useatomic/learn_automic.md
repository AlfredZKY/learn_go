# 为什么要用原子操作
虽然通过对互斥锁的合理使用，我们可以使一个goroutine在执行临界区中的代码时，不被其他的goroutine打扰，不过，虽然不会被打扰，但是它仍然可能被中断(interruption)

由于go语言运行时系统中调度器，会恰当地安排其中所有的goroutine的运行，不过，在同一时刻，只能有少说的goroutine真正地处于运行状态，并且这个数量是固定的,为了公平起见，调度器会频繁地换上或者换下这个goroutine，换上的意思是让一个goroutine由非运行状态转为运行状态，并促使其中的代码在某个CPU核心上执行。换下的意思正好相反，使一个goroutine中代码中断执行，并让它由运行状态转化为非运行状态
所以互斥锁虽然可以保证临界区中代码的串行，却不能保证这个代码执行的原子性(atomicity)

# 原子操作的原理
真正能够保证原子性执行的只有原子操作(atomic operation).原子操作在进行的过程中时不允许中断的，在底层，这会有CPU提供芯片级别的支持，所以绝对有效，即使在拥有多CPU核心，或者多CPU的计算机系统中，原子操作的保证也是不可撼动的。

# 原子操作的优缺点
- 这使得原子擦欧洲可以完全消除静态条件，并能够绝对的保证安全性，并且它的执行速度要比其他的同步工具快得多，通常高出好几个数量级
- 原子操作不能中断，所以它需要足够简单，并且要求快速

# sync/atomic包中提供了几种原子操作？可操作性的数据类型又有哪些？
- sync/atomic包中的函数可以做的原子操作有:加法(add)、比较并交换(compare and swap,简称CAS)、加载(load)、存储(store)、交换(swap)
- 这些函数针对的数据类型并不多，但是，对这些类型中的每一个，sync/atomic包都会有一套函数给与支持。这些数据类型有:int32、int64、uint32、uint64、uintptr以及unsafe包中的Pointer，不过针对unsafe.Pointer类型，该包并未提供进行原子加法操作的函数
- 此外，sync/atomic包还提供一个名为Value的类型，它可以被用来存储任意类型的值，指针

# 原子操作的参数类型为什么只能是指针
因为原子操作函数需要的是被操作值的指针，而不是这个值的本身；被传入函数的参数值都会被复制，像这种基本类型的值一旦被传入函数，就已经与函数外的那个值毫无关系了。所以传入值本身没有任何意义。
只有原子操作函数拿到了被操作值的指针们就可以定位到存储该值的内存地址，只有这样，它们才能够通过底层的指令，准确地操作这个内存地址上的数据


# 自旋锁
自旋锁是计算机科学用于多线程同步的一种锁，线程反复检查锁变量是否可用。由于线程在这一过程中保持执行，因此是一种忙等待，一旦获取了自旋锁，线程就会一直保持该锁，直至显式释放自旋锁
自旋锁避免了进程上下文的调度开销，因此对于线程只会阻塞很短时间的场合是有效的

# 原子值的使用建议
- 不要把内部使用的原子值暴露给外界。比如声明一个全局的原子变量
- 如果必须要访问原子值，可以声明一个公开函数，让外界间接的使用
- 如果通过某个函数可以向内部的原子值存储值的话，就应该在这个函数中先判断被存储值的类型合法性，若不合法，则直接返回对应的错误值，从而避免panic
- 如果可能的话，我们可以把原子值封装到一个数据类型中，如果一个结构体。