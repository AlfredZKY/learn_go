package slice

// 数组类型的值（以下简称数组）的长度是固定的，而切片类型的值（以下简称切片）是可变长的。
import (
	"strconv"
	"reflect"
	"fmt"
	"testing"
)


// func GetArrayLen
func TestGetArrayLen(t* testing.T){
	s1 := make([]int,5)
	fmt.Printf("The length of s1:%d\n",len(s1))
	fmt.Printf("The capacity of s1:%d\n",cap(s1))
	fmt.Printf("The value of s1:%d\n",s1)

	s2:=make([]int,5,8)
	fmt.Printf("The length of s2:%d\n",len(s2))
	fmt.Printf("The capacity of s2:%d\n",cap(s2))
	fmt.Printf("The value of s2:%d\n",s2)

	s3 :=[]int{1,2,3,4,5}
	fmt.Printf("The length of s3:%d\n",len(s3))
	fmt.Printf("The capacity of s3:%d\n",cap(s3))
	fmt.Printf("The value of s3:%d\n",s3)

	s4 := s3[2:4]
	fmt.Printf("The length of s4:%d\n",len(s4))
	fmt.Printf("The capacity of s4:%d\n",cap(s4))
	fmt.Printf("The value of s4:%d\n",s4[0:3])
	fmt.Printf("The most value of s4:%d\n",s4[0:cap(s4)])
}


// func Testsliceparam
func Testsliceparam(complexArray [3][]string)[3][]string{
	complexArray[1][1] = "hello"
	// return complexArray1
	return complexArray
}

func TestInitArray(t *testing.T){
	a := [...]int{1,2,3,4}

	for _,v:=range a{
		t.Log(v)
	}
}
func TestArraySection(t *testing.T){
	a := [...]int{1,2,3,4}
	arrSec := a[:]
	t.Log(arrSec)
}

func TestArrayType(t *testing.T){
	arrayA := [...]int{1,2,3} 
	arrayB := [...]int{1,2,3,4} 

	fmt.Println(reflect.TypeOf(arrayA) == reflect.TypeOf(arrayB))
	fmt.Println(reflect.TypeOf(arrayA))
	fmt.Println(reflect.TypeOf(arrayB))
}

func TestArrayStringAdd(t*testing.T){
	str1 := []string{"2","3","9","5"}
	for _,value := range str1{
		val,_:=strconv.Atoi(value)
		if val < 9 {
		}

		t.Log(value)
	}
}