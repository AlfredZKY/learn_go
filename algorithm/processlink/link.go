package processlink


import "learn_go/dataStruct/linkstruct"

// CreateLink 创建一个新链表
func CreateLink(){
	head := linkstruct.New(0)
	head.Append(1)
	head.Append(2)
	head.Append(4)
	head.Append(5)
	head.Append(6)
	head.Traverse()
}




