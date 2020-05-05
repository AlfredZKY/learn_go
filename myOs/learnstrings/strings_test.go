package learnstrings

import (
	"reflect"
	"testing"
	"strings"
	
)

func TestStringBuilder(t*testing.T){
	var builderTemp  strings.Builder
	res ,err := builderTemp.Write([]byte("hello"))
	if err == nil{
		t.Log(reflect.TypeOf(builderTemp))
		t.Log(builderTemp.String())
		t.Logf("writed %d bytes:\n",res)
	}
	
}

