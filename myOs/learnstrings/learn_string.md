# 与string值相比，string.Buileder类型的值的区别
- 已存在的内容不会改边
- 减少了内容分配和内容拷贝的次数
- 可将内容重置，可重用值

# 不是并发安全的，需要手动处理