# go编译指令
类似C++中的 #pragma pack(2)，Golang中也有一些编译指令。它们的实现方式是一些特殊的注释
编译指令不是语言的一部分。它们可能是编译器实现的，编程规范中也没有对它们的描述
语法：
//go:directive
编译指令的语法是一行特殊的注释，关键字//和go之间没有空格。
//go:noescape
`func NewBook()(*Book){`
`    b := Book{Mice:12,Men:9}`
`    return &b`
`}`
这段代码在C/C++中这样做，返回的是不可用的地址，显然是要出问题的。在go中是可以的。因为逃逸分析，b将会被分配在堆上

## //go:noescape 禁止逃逸分析
逃逸分析可以识别生命周期超出变量声明函数的生命周期，并将变量从栈的分配上移动到堆中
`func BuildLibrary() {`
`        b := Book{Mice: 99: Men: 3}`
`        AddToCollection(&b)`
`}`
此时b逃逸到堆中了？这取决于AddToCollection()对b做了什么？
`func AddToCollection(b *Book) {`
`        b.Classification = "fiction"`
`}` 此时分析发现AddToCollection并没有将*Book继续传递下去，所以此时b会被分配在栈上。
`var AvailableForLoan [] *Book`
`func AddToCollection（b * Book）{`
`        AvailableForLoan = append（AvailableForLoan，b）`
`}`此时分析发现AddToCollection这样做的话，把b append到了一个生命周期更长的slice中，所以b必须被分配在堆上，以保证生命周期大于AddToCollection和BuildLibrary的，逃逸分析必须知道AddToCollection对b做了什么，调用了什么func等，以了解分配在栈上还是堆上。

## //go:nosplit 跳过栈溢出检测
- 正是因为一个 Goroutine 的起始栈大小是有限制的，且比较小的，才可以做到支持并发很多 Goroutine，并高效调度。
stack.go 源码中可以看到，_StackMin 是 2048 字节，也就是 2k，它不是一成不变的，当不够用时，它会动态地增长。
那么，必然有一个检测的机制，来保证可以及时地知道栈不够用了，然后再去增长。
回到话题，nosplit 就是将这个跳过这个机制。

- 优劣
显然地，不执行栈溢出检查，可以提高性能，但同时也有可能发生 stack overflow 而导致编译失败

## //go:norace 跳过竟态检测
我们知道在多线程中，程序，难免会出现数据竞争，正常情况下，当编译器检测到有数据竞争，就会给出提示
- 优势：
    减少编译时间
- 缺点：
    竞争会导致程序的不确定性


