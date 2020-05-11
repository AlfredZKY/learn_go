# 总结
## WaitGroup
- sync包中的WaitGroup类型和Once类型都是非常易用的同步工具，它们都是开箱即用和并发安全的
- 利用WaitGroup值，可以很方便地实现一对多的goroutine协作流程，即：一个分发子任务的goroutine，和多个执行子任务的goroutine共同完成一个任务
- 此外使用WaitGroup值时，注意，千万不要让其中计数器的值小于0,否则会引发panic
另外在使用上，最好采用"先统一Add,再并发Done,最后Wait"这种标准方式，来使用WaitGroup值，尤其不要在调用Wait方法的同时，并发地通过Add方法去增加其计数器的值，(跨越了多个声明周期)也有可能引发panic
- 不要把增加计数和值的操作和调用其Wait方法的代码放在不同的goroutine中执行

## Once
Once值的使用更加简单，它只有一个Do方法，同一个Once值的Do方法，永远只会执行第一次被调用时传入的函数参数，不论这个函数会以怎样的方式结束。此外，只要传入的某个Do方法的参数函数没有执行结束，任何之后调用该方法的goroutine都会被阻塞，(Once的互斥锁的哪行代码上)只有在这个函数执行结束后，哪些goroutine才会被逐一被唤醒