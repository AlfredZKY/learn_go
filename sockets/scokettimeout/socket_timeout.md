# Socket TimeOut 
- 网络通信中，为了防止长时间无响应的情况，经常会用到网络连接超时、读写超时的设置。
# 超时设置
## 连接超时
- func DialTimeout(network, address string, timeout time.Duration) (Conn, error) 
    - 第三个参数timeout可以用来设置连接超时设置，如果超过了timeout指定的时间，连接没有完成，会返回超时错误
## 读写超时
- SetDeadline(t time.Time) error
- SetReadDeadline(t time.Time) error
- SetWriteDeadline(t time.Time) error

通过读取源码可知，这里的参数t是一个未来的时间点，多以每次读取或者写入前，都要调用SetXXX重新设置超时时间
如果只设置一次，就会总出现超时问题。

