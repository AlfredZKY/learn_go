"".add STEXT nosplit size=20 args=0x10 locals=0x0
	0x0000 00000 (direct_topfunc_call.go:4)	TEXT	"".add(SB), NOSPLIT|ABIInternal, $0-16 
	0x0000 00000 (direct_topfunc_call.go:4)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_topfunc_call.go:4)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_topfunc_call.go:4)	MOVL	"".b+12(SP), AX
	0x0004 00004 (direct_topfunc_call.go:4)	MOVL	"".a+8(SP), CX
	0x0008 00008 (direct_topfunc_call.go:4)	ADDL	CX, AX
	0x000a 00010 (direct_topfunc_call.go:4)	MOVL	AX, "".~r2+16(SP)
	0x000e 00014 (direct_topfunc_call.go:4)	MOVB	$1, "".~r3+20(SP)
	0x0013 00019 (direct_topfunc_call.go:4)	RET
	0x0000 8b 44 24 0c 8b 4c 24 08 01 c8 89 44 24 10 c6 44  .D$..L$....D$..D
	0x0010 24 14 01 c3                                      $...
"".main STEXT size=66 args=0x0 locals=0x18
	0x0000 00000 (direct_topfunc_call.go:6)	TEXT	"".main(SB), ABIInternal, $24-0
	0x0000 00000 (direct_topfunc_call.go:6)	MOVQ	(TLS), CX
	0x0009 00009 (direct_topfunc_call.go:6)	CMPQ	SP, 16(CX)
	0x000d 00013 (direct_topfunc_call.go:6)	PCDATA	$0, $-2
	0x000d 00013 (direct_topfunc_call.go:6)	JLS	58
	0x000f 00015 (direct_topfunc_call.go:6)	PCDATA	$0, $-1
	0x000f 00015 (direct_topfunc_call.go:6)	SUBQ	$24, SP
	0x0013 00019 (direct_topfunc_call.go:6)	MOVQ	BP, 16(SP)
	0x0018 00024 (direct_topfunc_call.go:6)	LEAQ	16(SP), BP
	0x001d 00029 (direct_topfunc_call.go:6)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (direct_topfunc_call.go:6)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (direct_topfunc_call.go:6)	MOVQ	$137438953482, AX
	0x0027 00039 (direct_topfunc_call.go:6)	MOVQ	AX, (SP)
	0x002b 00043 (direct_topfunc_call.go:6)	PCDATA	$1, $0
	0x002b 00043 (direct_topfunc_call.go:6)	CALL	"".add(SB)
	0x0030 00048 (direct_topfunc_call.go:6)	MOVQ	16(SP), BP
	0x0035 00053 (direct_topfunc_call.go:6)	ADDQ	$24, SP
	0x0039 00057 (direct_topfunc_call.go:6)	RET
	0x003a 00058 (direct_topfunc_call.go:6)	NOP
	0x003a 00058 (direct_topfunc_call.go:6)	PCDATA	$1, $-1
	0x003a 00058 (direct_topfunc_call.go:6)	PCDATA	$0, $-2
	0x003a 00058 (direct_topfunc_call.go:6)	CALL	runtime.morestack_noctxt(SB)
	0x003f 00063 (direct_topfunc_call.go:6)	PCDATA	$0, $-1
	0x003f 00063 (direct_topfunc_call.go:6)	NOP
	0x0040 00064 (direct_topfunc_call.go:6)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 2b 48  eH..%....H;a.v+H
	0x0010 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 48 b8 0a  ...H.l$.H.l$.H..
	0x0020 00 00 00 20 00 00 00 48 89 04 24 e8 00 00 00 00  ... ...H..$.....
	0x0030 48 8b 6c 24 10 48 83 c4 18 c3 e8 00 00 00 00 90  H.l$.H..........
	0x0040 eb be                                            ..
	rel 5+4 t=17 TLS+0
	rel 44+4 t=8 "".add+0
	rel 59+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........



