package useerror

import (
	"errors"
	"os/exec"
	"os"
	"fmt"
	"strconv"

)

// Errno 代表某种错误的类型
type Errno int

func (e Errno) Error() string {
	return "errno " + strconv.Itoa(int(e))
}

// underlyingError 会返回已知的操作系统相关错误的潜在错误值
func underlyingError(err error)error{
	switch err:=err.(type){
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return errors.New("unkown error")
}

// Echo determine parameter is None
func Echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty request")
		return
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}
