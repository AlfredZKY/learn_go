"".main STEXT size=87 args=0x0 locals=0x38
	0x0000 00000 (main.go:3)	TEXT	"".main(SB), ABIInternal, $56-0  // 56表示该函数栈帧的大小(变量和参数都要算在内) 0表示返回值大小
	0x0000 00000 (main.go:3)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:3)	PCDATA	$0, $-2
	0x000d 00013 (main.go:3)	JLS	80
	0x000f 00015 (main.go:3)	PCDATA	$0, $-1
	0x000f 00015 (main.go:3)	SUBQ	$56, SP    // 分配56bytes的栈空间    SP是一个指针寄存器，并且指向栈顶，栈又是从高地址向低地址分配，那么对它做一次减法，就是从高地址向低地址移动了指针 ，同理 加法就是把SP恢复到了初始状态
	0x0013 00019 (main.go:3)	MOVQ	BP, 48(SP) // 将BP的值记录到该位置 把BP的内容放入到48(SP)中
	0x0018 00024 (main.go:3)	LEAQ	48(SP), BP // 将栈低指针记录到BP 把48(SP)的地址放到BP寄存器中
	0x001d 00029 (main.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:4)	MOVQ	$3, "".a+40(SP) // 表示本地变量a
	0x0026 00038 (main.go:5)	MOVQ	$2, "".b+32(SP) // 表示本地变量b
	0x002f 00047 (main.go:6)	MOVQ	"".a+40(SP), AX // 将a变量写入寄存器
	0x0034 00052 (main.go:6)	MOVQ	AX, (SP) 		// 通过寄存器传入参数a
	0x0038 00056 (main.go:6)	MOVQ	$2, 8(SP)		// 通过寄存器传入参数b
	0x0041 00065 (main.go:6)	PCDATA	$1, $0     		
	0x0041 00065 (main.go:6)	CALL	"".returnTwo(SB) // 调用函数returnTwo
	0x0046 00070 (main.go:7)	MOVQ	48(SP), BP 		 // 将旧的BP值取回，赋值给BP
	0x004b 00075 (main.go:7)	ADDQ	$56, SP			 // 回收函数的栈空间 （释放掉了，并未清空）
	0x004f 00079 (main.go:7)	RET
	0x0050 00080 (main.go:7)	NOP
	0x0050 00080 (main.go:3)	PCDATA	$1, $-1
	0x0050 00080 (main.go:3)	PCDATA	$0, $-2
	0x0050 00080 (main.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0055 00085 (main.go:3)	PCDATA	$0, $-1
	0x0055 00085 (main.go:3)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 41 48  eH..%....H;a.vAH
	0x0010 83 ec 38 48 89 6c 24 30 48 8d 6c 24 30 48 c7 44  ..8H.l$0H.l$0H.D
	0x0020 24 28 03 00 00 00 48 c7 44 24 20 02 00 00 00 48  $(....H.D$ ....H
	0x0030 8b 44 24 28 48 89 04 24 48 c7 44 24 08 02 00 00  .D$(H..$H.D$....
	0x0040 00 e8 00 00 00 00 48 8b 6c 24 30 48 83 c4 38 c3  ......H.l$0H..8.
	0x0050 e8 00 00 00 00 eb a9                             .......
	rel 5+4 t=17 TLS+0
	rel 66+4 t=8 "".returnTwo+0
	rel 81+4 t=8 runtime.morestack_noctxt+0
"".returnTwo STEXT nosplit size=79 args=0x20 locals=0x10
	0x0000 00000 (main.go:9)	TEXT	"".returnTwo(SB), NOSPLIT|ABIInternal, $16-32 同main函数 
	0x0000 00000 (main.go:9)	SUBQ	$16, SP 	// 分配栈空间
	0x0004 00004 (main.go:9)	MOVQ	BP, 8(SP) 	// 记录BP的值记录到该位置
	0x0009 00009 (main.go:9)	LEAQ	8(SP), BP   // 将栈低指针记录到底BP
	0x000e 00014 (main.go:9)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (main.go:9)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (main.go:9)	MOVQ	$0, "".c+40(SP) // c的返回值初始为0
	0x0017 00023 (main.go:9)	MOVQ	$0, "".d+48(SP) // d的返回值初始为0
	0x0020 00032 (main.go:10)	MOVQ	$1, "".tmp(SP)  // 移动1到tmp中
	0x0028 00040 (main.go:11)	MOVQ	"".a+24(SP), AX // a b 通过寄存器 传参
	0x002d 00045 (main.go:11)	ADDQ	"".b+32(SP), AX
	0x0032 00050 (main.go:11)	MOVQ	AX, "".c+40(SP) // 计算后存储 40(SP)
	0x0037 00055 (main.go:12)	MOVQ	"".b+32(SP), AX
	0x003c 00060 (main.go:12)	SUBQ	"".tmp(SP), AX
	0x0040 00064 (main.go:12)	MOVQ	AX, "".d+48(SP) // 计算后存储 48(SP)
	0x0045 00069 (main.go:13)	MOVQ	8(SP), BP // 还原操作 BP 复位
	0x004a 00074 (main.go:13)	ADDQ	$16, SP  // 回收栈空间
	0x004e 00078 (main.go:13)	RET
	0x0000 48 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 48 c7  H...H.l$.H.l$.H.
	0x0010 44 24 28 00 00 00 00 48 c7 44 24 30 00 00 00 00  D$(....H.D$0....
	0x0020 48 c7 04 24 01 00 00 00 48 8b 44 24 18 48 03 44  H..$....H.D$.H.D
	0x0030 24 20 48 89 44 24 28 48 8b 44 24 20 48 2b 04 24  $ H.D$(H.D$ H+.$
	0x0040 48 89 44 24 30 48 8b 6c 24 08 48 83 c4 10 c3     H.D$0H.l$.H....
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
