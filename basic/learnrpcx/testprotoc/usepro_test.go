package testprotoc

import (
	"testing"
	pb "learn_go/learnrpcx/genprobuf/tutorial"
	proto "github.com/golang/protobuf/proto"
	
)

func TestPerson(t*testing.T){
	p := pb.Person{
		Id:1234,
		Name:"zky",
		Email:"1763030@163.com",
		Phones:[]*pb.Person_PhoneNumber{
			{Number:"110-119",Type:pb.Person_HOME},
		},
	}
	t.Log(p)
	res ,err := proto.Marshal(&p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
	var receive pb.Person
	err = proto.Unmarshal(res,&receive)
	if err != nil {
		panic(err)
	}
	t.Log(receive)
}