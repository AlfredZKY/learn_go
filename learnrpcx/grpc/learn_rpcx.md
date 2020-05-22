# go rpc服务

## rpc框架原理
    RPC框架的目标就是让远程服务调用更加简单，透明，RPC框架负责屏蔽底层的传输方式(TCP或者UDP)、序列化方式(XML/Json/二进制)和通信细节。服务调用者可以像调用本地接口一样调用远程的服务提供者，而不用需要关心底层通信细节和调用过程。

## RPC框架
- 支持多语言的RPC框架
- 只支持特定语言的RPC框架 例如新浪微博的Motan
- 支持服务治理等服务化特性的分布式服务框架，其底层内核仍然是RPC框架，例如阿里的Dubbo

## protocol buff 
    高性能的序列化协议，通过序列化后，方便调用

# rpcx rpc框架
- rpcx也支持protoc的插件，方便生成go代码，下面是两种插件
    - protoc-gen-rpcx:基于官方的grpc插件，生成标准的protobuf GO类型和rpcx代码
    - protoc-gen-gogorpcx: 基于 gogo/protobuf的插件，可以生成性能更加优异、更多辅助代码的Go代码和rpcx代码
