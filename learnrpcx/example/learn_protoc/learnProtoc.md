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
- ‵protoc --go_out=. .*proto‵
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

## 2.--go_out: protoc-gen-go: Plugin failed with status code 1.
- go get -u github.com/golang/protobuf/protoc-gen-go
- cp protoc-gen-go  /usr/local/bin/