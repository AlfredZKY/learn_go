# go生成汇编代码
- go tool compile -N -l -S name.go >> main.S
- go build -gcflags -S main.go

# 生成反汇编
- go tool objdump name.o

# 注意
- 注意:机器不同用的指令不同的
- amd64 6g
- i386 8g
