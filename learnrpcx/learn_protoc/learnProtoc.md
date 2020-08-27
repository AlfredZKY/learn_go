# proto buffer
- 优点
    - 相比较于json来说，数据量小
    - 性能更高，更加规范
    - 编解码速度快，数据体积小
    - 使用统一的规范，不用担心大小写解析失败等问题

- 缺点
    - 改动协议自断，需要重新生成文件
    - 数据完全没有可读性

# 安装protoc 
- [github网站](https://github.com/protocolbuffers/protobuf/releases)

## 安装直接从github上下载即可，注意平台即可

## 安装protobuf库文件
    go get github.com/golang/protobuf/proto

## 安装插件
    go get github.com/golang/protobuf/protoc-gen-go

## 生成go文件
- ‵protoc --go_out=. *.proto‵
    - .:生成的.pb.go后缀所在的路径为当前路径
    - *.*proto:自己写的protoc文件

# 安装gogoprotobuf
## 安装插件
- protoc-gen-gogo: 和protoc-gen-go生成的文件差不多，性能也几乎一样(稍微快一点点)
    //gogo
    go get github.com/gogo/protobuf/protoc-gen-gogo

- protoc-gen-gofast:生成的文件更复杂，性能也更高(快5-7倍)
    //gofast
    go get github.com/gogo/protobuf/protoc-gen-gofast

## 安装gogoprotobuf库文件
    go get github.com/gogo/protobuf/proto
    go get github.com/gogo/protobuf/gogoproto  //这个不装也没关系

## 生成go文件
    //gogo
    protoc --gogo_out=. *.proto
    
    //gofast
    protoc --gofast_out=. *.proto

# 问题
## 1.下载好的protoc文件路径问题
-  如果选择的是linux平台，请把下载好的elf文件放入环境变量中，切记所带的include文件夹中的内容也要放入到环境变量中，这样就可以使用protoc命令了
-  sudo cp -r ~/Downloads/protoc/include/* .  

## 2.--go_out: protoc-gen-go: Plugin failed with status code 1.
- go get -u github.com/golang/protobuf/protoc-gen-go
- cp protoc-gen-go  /usr/local/bin/

# proto文件常识
## 第一行一定要指定proto的版本号(非空非注释的第一行)不指定默认为2版本
## message相当于go的struct 相当于一个字面量
## 指定字段编号
- 在message定义中每个字段都有一个唯一的编号，这些编号被用来在二进制消息体中识别定义的这些字段，一旦message类型被用到后就不能在修改这些编号。
- <font face="黑体" color=red>注意:在二进制消息中1-15占用一个字节，16-2047占用两个字节</font>
- 19000-199999是给protofol buffers实现保留的字段标号，定义message时不能使用
## 定义字段的规则
- singular:一个遵循singular规则的字段，也就是说字段编码后消息体中只能出现一次
- repeated:遵循repeated规则的字段在消息体中可以有任意多个该字段制，也就是slice
## 指定包名
    package helloword;
    // 在proto文件中.会被替换成_
    // 替代上面的go包名，生成的个go文件的包名为hw
    option go_package = "hw";


# protobuff 与rpcx搭配生成rpc代码
- rpcx是一个流行的Go语言实现的服务治理的框架，只要你简单会写Go的函数，你就能实现带服务治理的RPC服务。
- proto格式的文件编写与上述介绍一致，编译时注意插件的选择，--gofast_out不支持timestamp.
- rpcx目前也支持了从proto生成rpc服务端和客户端的代码，也就是rpcio组织的两个插件
    - protoc-gen-rpcx:基于官方grpc插件，生成标准的protocbug GO类型和rpcx代码
    - protoc-gen-gogorpcx:基于gogo/protobuf的插件，可以生成性能更加优异，更过辅助代码的go代码和rpcx代码
## protoc-gen-rpcx
- 依照官方的说法(issue#1111), github.com/golang/protobuf/protoc-gen-go/grpc代码已经死了，不再维护
- golang/protobuf v1.4 以及后续版本基于protobuf 新版本(APIv2）还在开发，不稳定
- 现在基于1.3.5上实现
   - 因为protoc-gen-go并没有设计成易于使用的lib方式，所以protoc-gen-rpcx采用将代码复制到golang/protobuf v1.3.5中进行编译。
    export GO111MODULE=off
    export GOPATH="$(go env GOPATH)"
    go get github.com/golang/protobuf/{proto,protoc-gen-go}
    export GIT_TAG="v1.3.5" 
    git -C $GOPATH/src/github.com/golang/protobuf checkout $GIT_TAG
    go get github.com/rpcxio/protoc-gen-rpcx
    cd $GOPATH/src/github.com/golang/protobuf/protoc-gen-go &&  cp -r $GOPATH/src/github.com/rpcxio/protoc-gen-rpcx/{link_rpcx.go, rpcx} .
    go install github.com/golang/protobuf/protoc-gen-go 重新编译出protoc-gen-go执行程序
    protoc -I.:${GOPATH}/src  --go_out=plugins=rpcx:. helloworld.proto 生成rpcx格式的代码

## protoc-gen-gogorpcx
-更快的序列化和反序列化方法
- 更规范的Go数据结构
- 兼容 go protobuf
- 非常多的辅助方法
- 可以产生测试代码和benchmark代码
- 其它序列化格式

protoc --gofast_out=plugins=rpcx:. helloworld.proto
protoc --gogofast_out=plugins=rpcx:. helloworld.proto
protoc --gogofaster_out=plugins=rpcx:. helloworld.proto
protoc --gogoslick_out=plugins=rpcx:. helloworld.proto