# go 数据类型转化
- go语言不存在隐式类型转换，所有的类型都必须是显式的声明
- valueOfTypeB = typeB(valueOfTypeA) (类型B的值 = 类型B(类型A的值))
    - 比如 a:=5.0;b:=int(a)
- 只有相同底层类型的变量之间可以相互转换(如将int16类型转换成int32类型)，不同底层类型的变量相互转换时会引发编译错误(如将bool类型转换为int类型)
- string与int类型之间的转换
    - Itoa()函数用于将int类型数据转换为对应的字符串类型func Itoa(i int) string
    - Atoi()函数用于将字符串类型的整数转换为int类型 func Atoi(s string)(i int,err error)
- parse解析:parse系列函数用于将字符串转换为指定类型的值，其中包括ParseBool(),ParseFloat(),ParseInt(),ParseUint()
    - ParseBool()函数用于将字符串转换为bool类型的值，它只能接受1、0、true、false、T、F、t、f、False、True、TRUE、FALSE 其它的值均返回错误：func ParseBool(str string) (value bool, err error)
        changBool, err := strconv.ParseBool("1")
        if err == nil {
            fmt.Printf("str1: %v\n", err)
        }
        //字符串转化为bool型
        fmt.Printf("type:%T value:%#v\n", changBool, changBool)　
    - ParseInt()函数用于返回字符串表示的整数值(可以包含正负号)ParseUint()函数功能类似于与ParseInt()函数，但ParseUint()不接受正负号，用于无符号整型
        //base 指定进制，取值范围是 2 到 36。如果 base 为 0，则会从字符串前置判断，“0x”是 16 进制，“0”是 8 进制，否则是 10 进制。
        //bitSize 指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64。
        //返回的 err 是 *NumErr 类型的，如果语法有误，err.Error = ErrSyntax，如果结果超出类型范围 err.Error = ErrRange
        //转化为整数ParseInt()
        changeInt, err := strconv.ParseInt("-11", 10, 0)
        fmt.Printf("%#v\n", changeInt)
        changeUint, err := strconv.ParseUint("10", 10, 0)
        if err == nil {
            printf("err:%v\n", err)
        }
        fmt.Printf("value:%v\n", changeUint)　　
    - ParseFloat()函数用于将一个表示浮点数字符串转换为flaot类型 
        func ParseFloat(s string, bitSize int) (f float64, err error)
    - format 格式：Format系列函数实现了将给定类型数据格式化为字符串类型的功能，其中包括了FormatBoll(),FormatInt()、FormatUint()、FormatFloat()
    - append 附加:Append()系列函数用于将制定类型转换成字符串后追加到一个切片中，其中包含AppendBool()、AppendFloat()、AppendInt()、AppendUint()。
        // 将转换为10进制的string，追加到slice中
        b10 = strconv.AppendInt(b10, -42, 10)
        fmt.Println(string(b10))
        b16 := []byte("int (base 16):")
        b16 = strconv.AppendInt(b16, -42, 16)
        fmt.Println(string(b16))