# context的场景
处理多个goroutine之间同步问题
优点:
- context类型之所以受到了标准库中众多代码包的积极支持，主要因为它是一种非常通用的同步工具，它的值不但可以任意地扩散，而且还可以被用来传递额外的信息和信号
- context类型可以提供一类代表上下文的值，此类值是并发安全的，也就是说它可以被传播给多个goroutine

# context值的类型
- 根Context值
- 可撤销的Context值
  - 只可手动撤销的context值
  - 可定时撤销的context值
- 含数据的Context值
  - 含数据的context值可以携带数据，每个值都可以存储一对键值对，我们在调用它的Value方法的时候，它会沿着上下文树的根节点的方法逐个值的进行查找，如果发现相等的键，它就会立即返回对应的值，否则将在最后返回nil
  - 含数据的context值不能被撤销，而可撤销的context值又无法携带数据，但是由于他们共同组成了一个有机的整体，功能上要比sync.WaitGroup强大得多
  
所有的context值共同构成了一颗上下文树。这棵树的作用域是全局的，而根context值就是这颗树的根，它是全局唯一的，并且不提供任何额外的功能。

由于context类型实际上是一个接口类型，而context包中实现该接口的所有私有类型，都是基于某个数据类型的指针类型，所以如此传播并不会影响该类型值的功能和安全。

context类型的值是可以繁衍的，这意味着我们可以通过一个context值产生出任意个子值，这些子值可以携带其父值的属性和数据，也可以响应我们通过其父值传达的信号
正因此，所有的context值共同构成了一颗代表了上下文全貌的树形结构，这棵树的树根(或者上下文根节点)是一个已经在context包中预定义好的context值，它是全局唯一的。通过调用context.Backgroud函数，我们就可以获取到它。

这个上下文根节点仅仅是一个最基本的支点，它不提供任何额外的功能，也就是说，它即不可以被撤销，也不能携带任何数据。
这些函数的第一个参数类型都是context.Context,而名称都是parent.顾名思义，这个位置上的参数对应的都是它们将会产生的context值的父值

# 通过context包，实现一对多的goroutine协作流程


# Context使用原则和技巧
- 不要把context放在结构体中，要以参数的方式传递，parent Context一般为Background
- 应该要把Context作为第一个参数传递给入口请求和出口请求链路上的每一个函数，挡在第一位，变量名建议都统一，如ctx
- 给一个函数传递Context的时候，不要传递nil，否则在trace追踪的时候，就会断了连接
- Context的Value相关方法应该传递必须的数据，不要什么数据都用这个传递
- Context是线程安全的，可以放心的在多个goroutine中传递
- 可以把一个Context对象传递给任意个数的goroutine，对它执行"取消"操作时，所有的goroutine都会接受到取消信号
  
# context 接口
- Deadline方法需要返回当前Context被取消的时间，也就是完成工作的截止日期
- Done方法需要返回一个channel，这个channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done会返回用一个channel
- Err方法会返回当前context结束的原因，它只会在done返回的channel被关闭时才会返回非空值
  - 如果当前context被取消就返回Canceled错误
  - 如果当前context超时就会返回DeadlineExceeded错误
- Value方法会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value并传入相同的key会返回相同的结果，这个功能可以用来传递请求特定的数据