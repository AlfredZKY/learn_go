# TLS 线程本地存储
- 线程本地存储又叫线程局部存储，其英文为Thread Local Storage，简称TLS，看似一个很高大上的东西，其实就是线程私有的全局变量而已。


# gcc 编译
- __thread int g = 0;  // 1，这里增加了__thread关键字，把g定义成私有的全局变量，每个线程都有一个g变量
- gcc -g main.c -o thread -lpthread 进行编译直接运行
- 运行结果如下：
    ```
        ./thread        
        start, g[0x7f904e405aa0] : 0
        main, g[0x7f904e405a90] : 100
    ```

# gdb调试
- 打断点(主要是全局变量g)
    ```
    (gdb) b main.c:19
    Breakpoint 1 at 0x82c: file main.c, line 19.
    ```
- 打完断点后，运行
    ```
    (gdb) r
    Starting program: /root/zhou/thread
    [Thread debugging using libthread_db enabled]
    Using host libthread_db library "/lib/x86_64-linux-gnu/libthread_db.so.1".

    Breakpoint 1, main (argc=1, argv=0x7fffffffe5c8) at main.c:19
    19              g = 100;  // 2，主线程给私有全局变量赋值为100
    ```

- 反编译main函数
    ```
    (gdb) disass
    Dump of assembler code for function main:
    0x000055555555480e <+0>:     push   %rbp
    0x000055555555480f <+1>:     mov    %rsp,%rbp
    0x0000555555554812 <+4>:     sub    $0x20,%rsp
    0x0000555555554816 <+8>:     mov    %edi,-0x14(%rbp)
    0x0000555555554819 <+11>:    mov    %rsi,-0x20(%rbp)
    0x000055555555481d <+15>:    mov    %fs:0x28,%rax
    0x0000555555554826 <+24>:    mov    %rax,-0x8(%rbp)
    0x000055555555482a <+28>:    xor    %eax,%eax
    => 0x000055555555482c <+30>:    movl   $0x64,%fs:0xfffffffffffffffc
    0x0000555555554838 <+42>:    lea    -0x10(%rbp),%rax
    0x000055555555483c <+46>:    mov    $0x0,%ecx
    0x0000555555554841 <+51>:    lea    -0x8e(%rip),%rdx        # 0x5555555547ba <start>
    0x0000555555554848 <+58>:    mov    $0x0,%esi
    0x000055555555484d <+63>:    mov    %rax,%rdi
    0x0000555555554850 <+66>:    callq  0x555555554660 <pthread_create@plt>
    0x0000555555554855 <+71>:    mov    -0x10(%rbp),%rax
    0x0000555555554859 <+75>:    mov    $0x0,%esi
    0x000055555555485e <+80>:    mov    %rax,%rdi
    0x0000555555554861 <+83>:    callq  0x555555554690 <pthread_join@plt>
    0x0000555555554866 <+88>:    mov    %fs:0xfffffffffffffffc,%eax
    0x000055555555486e <+96>:    mov    %fs:0x0,%rdx
    0x0000555555554877 <+105>:   lea    -0x4(%rdx),%rcx
    0x000055555555487e <+112>:   mov    %eax,%edx
    0x0000555555554880 <+114>:   mov    %rcx,%rsi
    0x0000555555554883 <+117>:   lea    0xbd(%rip),%rdi        # 0x555555554947
    0x000055555555488a <+124>:   mov    $0x0,%eax
    0x000055555555488f <+129>:   callq  0x555555554680 <printf@plt>
    0x0000555555554894 <+134>:   mov    $0x0,%eax
    0x0000555555554899 <+139>:   mov    -0x8(%rbp),%rcx
    0x000055555555489d <+143>:   xor    %fs:0x28,%rcx
    0x00005555555548a6 <+152>:   je     0x5555555548ad <main+159>
    0x00005555555548a8 <+154>:   callq  0x555555554670 <__stack_chk_fail@plt>
    0x00005555555548ad <+159>:   leaveq
    0x00005555555548ae <+160>:   retq
    End of assembler dump.
    ```

- 从反汇编代码中可以看到下面这条指令
    - `=> 0x000055555555482c <+30>:    movl   $0x64,%fs:0xfffffffffffffffc` 
    - 运行到当前指令断点处，%fs代表段寄存器
    - 这句汇编指令的意思如下:
        - 是把常量100(0x64)复制到地址为%fs:0xfffffffffffffffc的内存中，
            可以看出全局变量g的地址为%fs:0xfffffffffffffffc，
            fs是段寄存器，0xfffffffffffffffc是有符号数-4，
            所以全局变量g的地址为：0x000055555555482c(fs段基址 - 4)
        - fs是段寄存器是段寄存器