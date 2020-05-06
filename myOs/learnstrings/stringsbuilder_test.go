package learnstrings

import (
	"reflect"
	"testing"
	"strings"
	
)

func TestStringBuilderWirte(t*testing.T){
	var builderTemp  strings.Builder
	// res ,err := builderTemp.Write([]byte("hello"))
	res ,err := builderTemp.Write([]byte("世界"))
	// res ,err := builderTemp.WriteRune('h')
	// res, err := builderTemp.WriteByte('h')
	if err == nil{
		t.Log(reflect.TypeOf(builderTemp))
		t.Log(builderTemp.String())
		t.Logf("writed %d bytes:\n",res)
	}
}

func TestStringBuilderWirteString(t*testing.T){
	var builderTemp  strings.Builder
	// res,err := builderTemp.WriteString("hello world")
	res,err := builderTemp.WriteString("世界")
	if err == nil{
		t.Log(reflect.TypeOf(builderTemp))
		t.Log(builderTemp.String())
		t.Logf("writed %d bytes:\n",res)
	}
}

func TestStringBuilderWirteRune(t*testing.T){
	var builderTemp  strings.Builder
	res ,err := builderTemp.WriteRune('h')
	
	if err == nil{
		t.Log(reflect.TypeOf(builderTemp))
		t.Log(builderTemp.String())
		t.Logf("writed %d bytes:\n",res)
	}
}

func TestStringBuilderWirteByte(t*testing.T){
	var builderTemp  strings.Builder
	err := builderTemp.WriteByte('h')
	if err == nil{
		t.Log(reflect.TypeOf(builderTemp))
		t.Log(builderTemp.String())
		// t.Logf("writed %d bytes:\n",res)
	}
}

func TestStringBuilderReset(t*testing.T){
	var b strings.Builder
	res,err := b.Write([]byte("hello world"))
	if err == nil{
		t.Log(res)
		t.Log(b.String())
	}
	temp := b.String()
	b.Reset()
	
	// res,err = b.Write([]byte("世界"))
	res,err = b.Write([]byte(temp))
	if err == nil{
		t.Log(res)
		t.Log(b.String())
	}
}
