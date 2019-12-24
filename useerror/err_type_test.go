package useerror

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

// 代表某种错误的类型
type Errno int

func (e Errno) Error() string {
	return "errno " + strconv.Itoa(int(e))
}

// func underlyingError(err error)error{
// 	switch err:=err.(type){
// 	case *os.PathError:
// 		return err.Err
// 	case *os.LinkError:
// 		return err.Err
// 	case *os.SyscallError:
// 		return err.Err
// 	case *exec.Error:
// 		return err.Err
// 	}
// }

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
