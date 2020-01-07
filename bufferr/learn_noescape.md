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

# 逃逸分析
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
