package myinterface

import (
	"fmt"
	"reflect"
)

// Pet define three func
type Pet interface{
	SetName(name string)
	Name()string
	Category()string
}

// Pet1 define two func
type Pet1 interface{
	Name()string
	Category()string
}

// Dog define struct
type Dog struct{
	name string // 名字
}

// SetName assignment variable
func (dog *Dog) SetName(name string){
	dog.name = name
}

//Name return a string
func (dog Dog)Name()string{
	return dog.name
}

//Category return a string
func (dog Dog)Category()string{
	return "dog"
}

// VarOfInter test inert value
func VarOfInter(){
	dog := Dog{name:"little pig"}
	_,ok := interface{}(dog).(Pet)
	fmt.Printf("dog implement interface Pet :%v\n",ok)

	// Dog类型附带的所有值方法和指针方法，又由于这3个方法恰恰分别是Pet接口中某个方法的实现
	// 所以*Dog类型就成为了Pet接口的实现类型
	_,ok = interface{}(&dog).(Pet)
	fmt.Printf("dog implement interface Pet :%v\n",ok)

	// 动态变量
	var pet Pet = &dog
	fmt.Printf("This pet is a %s,the name is %q.\n",pet.Category(),pet.Name()) 

	var pet1 Pet1 = dog
	fmt.Printf("This pet1 is a %s,the name is %q.\n",pet1.Category(),pet1.Name())
	
	// 这里调用SetName改变Name变量的值
	dog.SetName("monster")

	// 注意下，
	// 1.此时对于pet1变量来说值未变(通用规则:如果我们使用一个变量给另外一个变量赋值，那么真正赋给后者的并不是前者斥候的那个值，而是该值的一个副本)
	// 2.对于dog变量来说SetName是它的指针方法，值一定会改变
	fmt.Printf("This pet1 is a %s,the name is %q.\n",pet1.Category(),pet1.Name())
	fmt.Printf("This dog is a %s,the name is %q.\n",dog.Category(),dog.Name())

	dog1 := Dog{name:"little pig"}
	dog2 := dog1
	dog1.name = "monster"
	fmt.Printf("This dog1 is a %s,the name is %q.\n",dog1.Category(),dog1.Name())
	fmt.Printf("This dog2 is a %s,the name is %q.\n",dog2.Category(),dog2.Name())
}

// VarOfInterNil test interface value nil
func VarOfInterNil(){
	// 定义一个指针变量未初始化
	var dog *Dog
	fmt.Println("The first dog is nil.",reflect.TypeOf(dog))
	dog1 := dog 
	fmt.Println("The second dog is nil",reflect.TypeOf(dog1))

	// 在go语言中，我们把由字面量nil表示的值叫做无类型的nil,这是真正的nil,因为它的类型也是nil的
	// dog1赋值给pet时，go会把它的类型和值放在一起考虑，也就是说go会识别出赋予pet的值是一个*Dog类型
	// 的nil,然后，go就会用一个iface的实例包装它，包装后的产物肯定就不是nil了。

	var pet2 Pet1
	if pet2 == nil{
		fmt.Println("pet2 is nil",reflect.TypeOf(pet2))
	}else{
		fmt.Println("pet2 is not nil",reflect.TypeOf(pet2))
	}

	// 接口类型的变量赋值操作，pet是dog1的副本，动态类型是*Dog
	var pet Pet1 = dog1
	if pet == nil{
		fmt.Println("The pet is nil",pet)
	}else{
		// 虽然不是nil，但是pet的动态值是nil
		fmt.Println("The pet is not nil",pet)
		// 打印出pet的动态类型
		fmt.Printf("%T", pet)
	}
}
