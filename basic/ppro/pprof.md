# GO 性能分析
## 性能分析的包
- runtime/pprof：采集程序(非Server)的运行数据进行分析
- net/http/pprof：采集HTTP Server的运行时数据进行分析，这个其实在上面的功能中包了一层提供http接口
- pprof可用于可视化和性能分析的工具，pprof以profile.proto读取分析样本的集合，并生成报告以可视化并帮助分析数据(支持文本和图形报告)
- 这个文件是一个ProtocolBuffer v3的描述文件，它描述了一组callstack和symbolization信息，作用是表示统计分析的一组采样的调用栈，是很常见的stacktrace的配置格式

## 使用 
- Report generation：报告生成，直接生成一个文件，解析这个文件得到结果
- Interactive terminal use：交互式终端使用，实时反馈，监控，需要开发人员输入指令,根据输入的指令返回想要的信息。
- Web interface：Web 界面，实时反馈,监控，对开发人员友好。很方便，直观的获取和统计需要的数据。

## 能做做什么
- CPU Profiling: CPU分析，按照一定的频率采集所监听的应用程序的CPU使用情况，可确定应用程序在主动消耗 CPU 周期时花费时间的位置。
- Memory Profiling:内存分析，在应用程序堆栈分配时记录跟踪，用于监视当前和历史内存使用情况，检查内存泄漏情况。
- Block Profiling：阻塞分析，记录goroutine阻塞等待同步的位置
- Mutex Profiling:互斥锁分析，报告互斥锁的竞争情况

## 命令介绍
- allocs：所有过去内存分配的采样
- block：导致同步原语阻塞的堆栈跟踪
- cmdline：当前程序的命令行调用
- goroutine：所有当前goroutine的堆栈跟踪
- heap：活动对象的内存分配的采样。在获取堆样本之前,可以指定gc GET参数来运行gc。
- metux：争用互斥锁持有者的堆栈跟踪
- profile：CPU配置文件。您可以在seconds GET参数中指定持续时间。获取配置文件后，使用go tool pprof命令调查配置文件
- threadcreate：导致创建新操作系统线程的堆栈跟踪
- trace：当前程序的执行轨迹。您可以在seconds GET参数中指定持续时间。获取跟踪文件后，使用go tool trace命令调查跟踪

## 交互式终端使用
- 60秒后进入pprof交互式命令中
    - go tool pprof http://localhost:9909/debug/pprof/profile?seconds=60
    - 标题解释
        - flat：给定函数上的运行耗时
        - flat% ：给定函数上的CPU运行耗时占比
        - sum% ：给定函数累积使用CPU总比例
        - cum ：当前函数加上它之前的调用运行总耗时
        - cum% ：当前函数加上他之前的调用CPU运行耗时占比