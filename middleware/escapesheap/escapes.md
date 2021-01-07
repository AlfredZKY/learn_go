# 什么是堆/栈
在这里并不打算详细介绍堆栈，仅简单介绍本文所需的基础知识。如下： 
堆（Heap）：一般来讲是人为手动进行管理，手动申请、分配、释放。一般所涉及的内存大小并不定，一般会存放较大的 对象。另外其分配相对慢，涉及到的指令动作也相对多 
栈（Stack）：由编译器进行管理，自动申请、分配、释放。一般不会太大，我们常见的函数参数（不同平台允许存放的数 量不同），局部变量等等都会存放在栈上 


# 什么是逃逸分析
在编译程序优化理论中，逃逸分析是一种确定指针动态范围的方法，简单来说就是分析在程序的那些地方可以访问到该指针 
通缩地讲，逃逸分析就是确定一个变量要放在堆上还是栈上，规则如下：
- 是否有在其他地方(非局部)被引用。只要有可能被引用，那它一定分配到堆上。否则分配到栈上  
- 即使没有被外部引用，但对象过大，无法存放在栈区上，依然有可能分配到堆上
因此，逃逸分析是编译器用于决定变量分配到堆上还是栈上的一种行为

# 在什么阶段确立逃逸
在编译阶段确立逃逸，注意并不是在运行时

# 为什么需要逃逸
这个问题我们可以反过来想，如果变量都分配到堆上了会出现什么事情？例如：
- 垃圾回收（GC）的压力不断增大
- 申请、分配、回收内存的系统开销增大（相对于栈）
- 动态分配产生一定量的内存碎片

其实总的来说，就是频繁申请、分配堆内存是有一定 “代价” 的。会影响应用程序运行的效率，间接影响到整体系统。因此 “按需分配” 最大限度的灵活利用资源，才是正确的治理之道。

# 怎么确定是否逃逸
- 通过编译器命令，就可以看到详细的逃逸分析过程。而指令集 `-gcflags` 用于将标识参数传递给Go编译器:
    - -m 会打印出逃逸分析的优化策略，实际上最多可以4个-m ,但是信息量较大，一般用1个就可以了
    - -l 会禁用函数内联，在这里禁掉inline能更好的观察逃逸情况，减少干扰

    `go build -gcflags '-m -l ' main.go`
- 通过反编译命令查看 
    `go tool compile -S main.go`
    注:通过go tool compile -help 查看所有允许传递给编译器的标识参数

# 例
- 指针
```
// User sds
type User struct {
	ID     int64
	Name   string
	Avatar string
}

// GetUserInfo dsd
func GetUserInfo() *User {
	return &User{ID: 13746731, Name: "EDDYCJY", Avatar: "https://avatars0.githubusercontent.com/u/13746731"}
}

func main() {
	_ = GetUserInfo()
}

go build -gcflags '-m -l' main.go 
# command-line-arguments
./main.go:11:9: &User literal escapes to heap

go tool compile -S main.go       
"".GetUserInfo STEXT size=117 args=0x8 locals=0x18
        0x0000 00000 (main.go:11)       TEXT    "".GetUserInfo(SB), ABIInternal, $24-8
        0x0000 00000 (main.go:11)       MOVQ    (TLS), CX
        0x0009 00009 (main.go:11)       CMPQ    SP, 16(CX)
        0x000d 00013 (main.go:11)       PCDATA  $0, $-2
        0x000d 00013 (main.go:11)       JLS     110
        0x000f 00015 (main.go:11)       PCDATA  $0, $-1
        0x000f 00015 (main.go:11)       SUBQ    $24, SP
        0x0013 00019 (main.go:11)       MOVQ    BP, 16(SP)
        0x0018 00024 (main.go:11)       LEAQ    16(SP), BP
        0x001d 00029 (main.go:11)       FUNCDATA        $0, gclocals·2a5305abe05176240e61b8620e19a815(SB)
        0x001d 00029 (main.go:11)       FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (main.go:12)       LEAQ    type."".User(SB), AX
        0x0024 00036 (main.go:12)       MOVQ    AX, (SP)
        0x0028 00040 (main.go:12)       PCDATA  $1, $0
        0x0028 00040 (main.go:12)       CALL    runtime.newobject(SB) // 在这里申请了对象，并存储在堆上
        0x002d 00045 (main.go:12)       MOVQ    8(SP), AX
        0x0032 00050 (main.go:12)       MOVQ    $13746731, (AX)
        0x0039 00057 (main.go:12)       MOVQ    $7, 16(AX)
        0x0041 00065 (main.go:12)       LEAQ    go.string."EDDYCJY"(SB), CX
        0x0048 00072 (main.go:12)       MOVQ    CX, 8(AX)
        0x004c 00076 (main.go:12)       MOVQ    $49, 32(AX)
        0x0054 00084 (main.go:12)       LEAQ    go.string."https://avatars0.githubusercontent.com/u/13746731"(SB), CX
```    

- 未确定类型
```

// User sds
type User struct {
	ID     int64
	Name   string
	Avatar string
}

// GetUserInfo dsd
func GetUserInfo() *User {
	return &User{ID: 13746731, Name: "EDDYCJY", Avatar: "https://avatars0.githubusercontent.com/u/13746731"}
}

func main() {
	// _ = GetUserInfo()

	str := new(string)
	*str = "EDDYCJY"
	fmt.Println(str)
}

go build -gcflags '-m -l' main.go
# command-line-arguments
./main.go:16:9: &User literal escapes to heap
./main.go:22:12: new(string) escapes to heap
./main.go:24:13: ... argument does not escape
```

- 参数泄露
```
// User sds
type User struct {
	ID     int64
	Name   string
	Avatar string
}

// GetUserInfo dsd
func GetUserInfo() *User {
	return &User{ID: 13746731, Name: "EDDYCJY", Avatar: "https://avatars0.githubusercontent.com/u/13746731"}
}

func GetUserInfos(u *User) *User {
	return u
}


func main() {
	// _ = GetUserInfo()
	_ = GetUserInfos(&User{ID: 13746731, Name: "EDDYCJY", Avatar: "https://avatars0.githubusercontent.com/u/13746731"})
	// str := new(string)
	// *str = "EDDYCJY"
	// fmt.Println(str)
}

go build -gcflags '-m -l' main.go
# command-line-arguments
./main.go:16:9: &User literal escapes to heap
./main.go:19:19: leaking param: u to result ~r1 level=0
./main.go:26:19: &User literal does not escape
```