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

## //go:linkname
//go:linkname localname importpath.name 该指令只是编译器使用importpath.name作为源代码中声明为localname的变量或函数的目标文件符号名称，但是由于这个伪指令，可以破坏类型系统和包模块，因此只有当前引用了safe包的时候才可以用
简单来说，就是importpath.name 是localname的符号别名，编译器实际上会调用localname，但前提是使用了safe包才可以使用
例如 
- go 源码中的runtime/time.go
- go 源码中的runtime/timestub.go
- 其中 go:linkname time_now time.now 的意思是time_now是time.now的别名
    //go:linkname time_now time.now 
    func time_now() (sec int64, nsec int32, mono int64) {
        sec, nsec = walltime()
        return sec, nsec, nanotime()
    }

## //go:nowritebarrierrec 
该指令表示编译器遇到写屏障就会产生一个错误，并且允许递归，也就是这个函数调用的其他函数如果有写屏障也会报错，简单来说
就是针对写屏障的处理，防止其死循环
    //go:nowritebarrierrec
    func gcFlushBgCredit(scanWork int64) {
        ...
    }

## //go:yeswritebarrierrec
该指令与go:nowritebarrierrec相对，在标注go:nowritebarrierrec指令的函数上，遇到写屏障会产生错误，而当编译器遇到
go:yeswritebarrierrec指令时将会停止。


## //go:online
该指令表示该函数禁止进行内联，编译器在进行内联优化时，不一定说不好，所以要根据具体的实际情况进行判断


## //go:notinheap
该指令常用于类型声明，它表示这个类型不允许从GC堆上进行申请内存，在运行时常用来做较低层次的内部结构，避免调度器和内存分配中的写屏障。能够提高性能
