# 边界检查消除
Go是一个内存安全的语言。在数组和切片的索引和子切片操作中，Go运行时将检查操作中使用的下标是否越界。 如果下标越界，一个恐慌将产生，以防止这样的操作破坏内存安全。这样的检查称为边界检查。 边界检查使得我们的代码能够安全地运行；但是另一方面，也使得我们的代码运行效率略微降低。

从Go官方工具链1.7开始，官方标准编译器使用了一个新的基于SSA（single-assignment form，静态单赋值形式）的后端。 SSA使得Go编译器可以有效利用诸如BCE（bounds check elimination，边界检查消除）和CSE（common subexpression elimination，公共子表达式消除）等优化。 BCE可以避免很多不必要的边界检查，CSE可以避免很多重复表达式的计算，从而使得编译器编译出的程序执行效率更高。有时候这些优化的效果非常明显。

本文将展示一些例子来解释边界检查消除在官方标准编译器1.7+中的表现。

对于Go官方工具链1.7+，我们可以运行go build -gcflags="-d=ssa/check_bce/debug=1"来列出哪些代码行仍然需要边界检查。



