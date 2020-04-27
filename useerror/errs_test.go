package useerror

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"testing"
)


func TestEcho(t*testing.T){
	Echo("")
}

func TestErrorType(t *testing.T) {
	var err error
	_, err = exec.LookPath(os.DevNull)
	fmt.Printf("error: %s\n", err)
	if execErr,ok:=err.(*exec.Error);ok{
		execErr.Name = os.TempDir()
		execErr.Err = os.ErrNotExist
	}
	fmt.Printf("error: %s\n", err)

	err = os.ErrPermission
	if os.IsPermission(err){
		fmt.Printf("error(permission):%s\n",err)
	}else{
		fmt.Printf("error(other):%s\n",err)
	}

	os.ErrPermission = os.ErrExist
	// 上面这行代码修改了os包中已定义的错误值
	// 这样做会导致下面判断的结果不正确，并且，这会影响当前go程序中所有的此类判断
	// 所有，一定要避免这样做
	if os.IsPermission(err){
		fmt.Printf("error(permission):%s\n",err)
	}else{
		fmt.Printf("error(other):%s\n",err)
	}

	const(
		ERR0 = Errno(0)
		ERR1 = Errno(1)
		ERR2 = Errno(2)
	)
	var myErr error = Errno(2)
	switch myErr{
	case ERR0:
		fmt.Println("ERR0")
	case ERR1:
		fmt.Println("ERR1")
	case ERR2:
		fmt.Println("ERR2")
	}
}

func TestUserErrorFirst(t *testing.T) {
	for _, req := range []string{"", "hello"} {
		//t.Logf("request:%s\n",req)
		resp, err := Echo(req)
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
