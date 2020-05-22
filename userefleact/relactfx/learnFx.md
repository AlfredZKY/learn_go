# golang反射框架Fx
Fx是一个golang版本的依赖注入框架，它使得golang通过可重用，可组合的模块化来构建golang应用程序变得非常容易。

Fx是通过使用依赖注入的方式替换了全局通过手动方式来连接不同函数调用的复杂度，也不同于其他的依赖注入方式，Fx能够像普通golang函数使用，而不需要通过使用struct标签或内嵌特定类型。

# 反射reflact
- 反射是指在程序运行期对程序本身进行访问和修改的能力，程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到执行部分，在运行程序时，程序无法获取自身的信息。

- 反射中的类型和种类
    - go中的类型(Type)是指原生系统数据类型，如int,string,bool等，以及使用type关键字定义的类型，这种类型的名称就是其类型本身的名称，如使用type A struct{}定义的结构体时，A就是struct{}的类型
    种类是指对象归属的品种。

# 使用反射值对象包装任意值
go中使用reflact.ValueOf()函数获得值得反射对象(reflact.Value)
    rValue := reflect.ValueOf(rawValue)
    reflect.ValueOf返回reflect.Value类型
reflect.ValueOf返回reflect.Value类型，包含有rawValue的值信息。reflect.Value与原值间可以通过值包装和值获取互相转化，reflect.Value是一些反射操作的重要类型，如反射调用函数

# 结构体成员的方法列表
方法 | 说明 | 
-----|-----|
 Field(i int) StructField  | 根据索引，返回索引对应的结构体字段的信息。当值不是结构体或索引超界时发生宕机 |
| NumField() int | 返回结构体成员字段数量。当类型不是结构体或索引超界时发生宕机|
| FieldByName(name string) (StructField, bool) | 根据给定字符串返回字符串对应的结构体字段的信息。没有找到时 bool 返回 false，当类型不是结构体或索引超界时发生宕机 |
| FieldByIndex(index []int) StructField | 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。没有找到时返回零值。当类型不是结构体或索引超界时 发生宕机 |
| FieldByNameFunc( match func(string) bool) (StructField,bool) | 根据匹配函数匹配需要的字段。当值不是结构体或索引超界时发生宕机 |

# 反射值获取原始值的方法
方法名 | 说明 |
---- | ----|
| Interface() interface{} | 将值以 interface{} 类型返回，可以通过类型断言转换为指定类型 |
| Int() int64 | 将值以 int 类型返回，所有有符号整型均可以此方式返回 |
| Uint() uint64 | 将值以 uint 类型返回，所有无符号整型均可以此方式返回 |
| Float() float64 | 将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回 |
| Bool() bool | 将值以 bool 类型返回 |
| Bytes() []bytes | 将值以字节数组 []bytes 类型返回 |
| String() string | 将值以字符串类型返回 |



