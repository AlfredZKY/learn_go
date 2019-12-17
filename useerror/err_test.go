package useerror

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty request")
		return
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}

func TestUserErrorFirst(t *testing.T) {
	for _, req := range []string{"", "hello"} {
		//t.Logf("request:%s\n",req)
		resp, err := echo(req)
		if err != nil {
			t.Logf("error: %s\n", err)
		} else {
			t.Logf("response: %s\n", resp)
		}
	}

	err1 := fmt.Errorf("invalid contents:%s", "#$%")
	err2 := errors.New(fmt.Sprintf("invalid contents:%s", "#$%"))
	if err1.Error() == err2.Error() {
		t.Log("The error message in err1 and err2 are the same")
		t.Log(err1.Error())
	}
}

// underlyingError 会返回已知的操作系统相关错误的潜在错误值
func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}

func TestOsError(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Logf("unexpected error:%s\n", err)
	}

	// 人为制造 *osPathError类型的错误
	r.Close()
	_, err = w.Write([]byte("hi"))
	uError := underlyingError(err)
	t.Logf("underlying error:%s (type %T)\n", uError, uError)
}

func TestOsPathError(t *testing.T) {
	path := []string{
		runtime.GOROOT(),     //当前环境下的Go语言根目录
		"/it/must/not/exist", // 不存在的目录
		os.DevNull,
	}

	printError := func(i int, err error) {
		if err == nil {
			t.Log("nil error")
		}
		err = underlyingError(err)
		if os.IsExist(err) {
			t.Logf("error(exist)[%d]:%s\n", i, err)
		} else if os.IsNotExist(err) {
			t.Logf("error(not exist)[%d]:%s\n", i, err)
		} else if os.IsPermission(err) {
			t.Logf("error(permission)[%d]:%s\n", i, err)
		} else {
			t.Logf("error(other)[%d]:%s\n", i, err)
		}
	}

	{
		index := 0
		err := os.Mkdir(path[index], 0700)
		printError(index, err)
	}

	{
		index := 1
		f, err := os.Open(path[2])
		printError(index, err)
		if f != nil {
			f.Close()
		}
	}

	{
		index := 2
		_, err := exec.LookPath(path[index])
		printError(index, err)
	}

}
